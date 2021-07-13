package common

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"math"
	"net"
	"os"
	"sync"
	"time"

	log "github.com/YoungPioneers/blog4go"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/mervick/aes-everywhere/go/aes256"
)

var conn *net.TCPConn
var hbFlag bool = false
var sLock *sync.Mutex
var id string

var v = flag.String("v", "0", "调试信息")
var server = flag.String("s", "127.0.0.1", "服务器地址")
var port = flag.String("p", "3360", "服务器端口")
var pwd = flag.String("pwd", "123456", "驱动端密码")
var spwd = flag.String("spwd", "654321", "服务端密码")
var mode = flag.String("m", "xbox", "模拟模式")
var order = flag.String("order", "false", "字节序")
var testMode = flag.String("test", "0", "测试模式")

// optionally set user defined hook for logging
type ConsoleHook struct {
	something string
}

// when log-level exceed level, call the hook
// level is the level associate with that logging action.
// message is the formatted string already written.
func (self *ConsoleHook) Fire(level log.LevelType, tags map[string]string, args ...interface{}) {
	if *v == "1" {
		fmt.Println(args...)
	}
}

func Connect() {
	var err error
	err = log.NewWriterFromConfigAsFile("log_config.xml")
	if err != nil {
		fmt.Println(err)
	}
	uuid, _ := uuid.NewUUID()
	id = uuid.String()

	hook := new(ConsoleHook)
	log.SetHook(hook) // writersFromConfig can be replaced with writers
	log.SetHookLevel(log.DEBUG)
	log.SetHookAsync(true) // hook will be called in async mode

	if err != nil {
		fmt.Println(err)
	}

	raddr, err := net.ResolveTCPAddr("tcp", *server+":"+*port)
	if err != nil {
		log.Critical(err)
	}
	conn, err = net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Critical(err)
	}
	conn.SetNoDelay(true)
	sLock = new(sync.Mutex)
	startRecive()
	readUserInput()
	SendHello()
}

func GetConn() net.Conn {
	return conn
}

func readUserInput() {
	go func() {
		in := bufio.NewReader(os.Stdin)
		for {
			line, _, _ := in.ReadLine()
			str := string(line)
			switch str {
			case "ds4":
				UseDS4()
			case "xbox":
				UseXbox360()
			}
		}
	}()
}

func startRecive() {
	// go func() {
	// 	time.Sleep(time.Second * 3)
	// 	hbFlag = true
	// 	for {
	// 		hbData := &Data{
	// 			Cmd:     CmdType_T_Ping,
	// 			MsgType: MsgType_Server,
	// 		}
	// 		SendData(hbData)
	// 		time.Sleep(time.Second * 3)
	// 	}
	// }()
	go func() {
		bufferReader := bufio.NewReader(conn)
		var err error
		for {
			if err != nil {
				// if err == io.EOF {
				// 	fmt.Printf("client %s is close!\n", conn.RemoteAddr().String())
				// }
				log.Critical(err)
			}
			dataLen := []byte{0, 0, 0, 0}
			len, err := bufferReader.Read(dataLen)
			if err == io.EOF || len == 0 {
				continue
			}
			data, _ := bufferReader.Peek(int(bytesToInt(dataLen)))
			// log.Info(data)
			protoData := &Data{}
			err = proto.Unmarshal(data, protoData)
			log.Info("REC: ")
			log.Info(len)
			log.Info(protoData)
			if err != nil {
				log.Critical(err)
			}
			switch protoData.Cmd {
			case CmdType_T_Hello:
				if protoData.MsgType == MsgType_Driver {
					if protoData.Hello.Msg == "" {
						SendHello()
					}
					switch *mode {
					case "ds4":
						UseDS4()
					case "xbox":
						UseXbox360()
					}
				}
			case CmdType_T_Ping:
				if protoData.MsgType == MsgType_Server {
					SendData(protoData)
					break
				}
				protoData.MsgType = MsgType_Pad
				SendData(protoData)
			}
			bufferReader.Read(data)
		}
	}()
}

func UseDS4() {
	protoData := &Data{
		Cmd:     CmdType_T_Pad_Type,
		MsgType: MsgType_Pad,
		PadType: PadType_DS4,
	}
	SendData(protoData)
}

func UseXbox360() {
	protoData := &Data{
		Cmd:     CmdType_T_Pad_Type,
		MsgType: MsgType_Pad,
		PadType: PadType_Xbox360,
	}
	SendData(protoData)
}

func SendHello() {
	// var order bool
	// sysType := runtime.GOOS
	// cpuType := runtime.GOARCH

	// if sysType == "linux" {
	// 	switch cpuType {
	// 	case "arm64":
	// 		order = false
	// 	case "arm":
	// 		order = false
	// 	default:
	// 		order = false
	// 	}
	// }

	// if sysType == "windows" {
	// 	order = true
	// }

	var byteOrder bool = false
	switch *order {
	case "true":
		byteOrder = true
	case "false":
		byteOrder = false
	}

	hello := aes256.Encrypt("hello", *pwd)
	shello := aes256.Encrypt("hello", *spwd)
	protoData := &Data{
		Cmd:     CmdType_T_Hello,
		MsgType: MsgType_Pad,
		Hello: &Hello{
			Group:     "test",
			Msg:       hello,
			ServerMsg: shello,
			Order:     byteOrder,
		},
	}
	SendData(protoData)
}

func SendBtn(btn PadBtn, val int32) {
	protoData := &Data{
		Cmd:     CmdType_T_Pad_Data,
		MsgType: MsgType_Pad,
		PadData: &PadBtnData{
			BtnType: btn,
			BtnVal:  val,
		},
	}
	SendData(protoData)
}

func SendData(protoData *Data) {
	sLock.Lock()
	protoData.Id = id
	// log.Info(protoData)
	data, _ := proto.Marshal(protoData)
	dataLen := intToBytes(int32(len(data)))
	log.Debug("SEND:")
	log.Debug(int32(len(data)))
	// log.Debug(data)
	log.Debug(protoData)
	_, err := conn.Write(dataLen)
	if err != nil {
		log.Info(err)
		os.Exit(1)
		return
	}
	_, err = conn.Write(data)
	if err != nil {
		log.Info(err)
		time.Sleep(time.Second * 1)
		os.Exit(1)
	}
	sLock.Unlock()
}

func intToBytes(x int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

func bytesToInt(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return x
}

func encodeUInt32(n uint) []byte {
	byteList := make([]byte, 0)
	for n != 0 {
		tmp := n % 128
		next := n >> 7
		if next != 0 {
			tmp = tmp + 128
		}
		byteList = append(byteList, byte(tmp))
		n = next
	}
	return byteList
}

func decodeUInt32(data []byte, offset int) (uint32, int) {
	n := uint32(0)
	length := 0
	for i := offset; i < len(data); i++ {
		length++
		m := uint32(data[i])
		n = n + ((m & 0x7f) * uint32(math.Pow(2, float64(7*(i-offset)))))
		if m < 128 {
			break
		}
	}
	return n, length
}
