package main

import (
	"RemotePad/src/common"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/YoungPioneers/blog4go"
	. "github.com/splace/joysticks"
)

func main() {
	flag.Parse()
	common.Connect()

	device := Connect(1)

	log.Infof("HID#1:- Buttons:%d, Hats:%d\n", len(device.Buttons), len(device.HatAxes)/2)

	b1press := device.OnClose(1)
	b1open := device.OnOpen(1)
	b2press := device.OnClose(2)
	b2open := device.OnOpen(2)
	b3press := device.OnClose(3)
	b3open := device.OnOpen(3)
	b4press := device.OnClose(4)
	b4open := device.OnOpen(4)

	b5press := device.OnClose(5)
	b5open := device.OnOpen(5)
	b6press := device.OnClose(6)
	b6open := device.OnOpen(6)
	// b7press := device.OnClose(7)
	// b7open := device.OnOpen(7)
	// b8press := device.OnClose(8)
	// b8open := device.OnOpen(8)

	b9press := device.OnClose(9)
	b9open := device.OnOpen(9)
	b10press := device.OnClose(10)
	b10open := device.OnOpen(10)
	b11press := device.OnClose(11)
	b11open := device.OnOpen(11)
	b12press := device.OnClose(12)
	b12open := device.OnOpen(12)
	b13press := device.OnClose(13)
	b13open := device.OnOpen(13)

	pan1X := device.OnPanX(1)
	pan2X := device.OnPanX(2)
	pan3X := device.OnPanX(3)
	pan4X := device.OnPanX(4)

	pan1Y := device.OnPanY(1)
	pan2Y := device.OnPanY(2)
	pan3Y := device.OnPanY(3)
	pan4Y := device.OnPanY(4)

	if device == nil {
		panic("no HIDs")
	}

	go device.ParcelOutEvents()

	// handle event channels
	go func() {
		for {
			select {
			case <-b1press:
				log.Info("button #1 pressed")
				common.SendBtn(common.PadBtn_A, 1)
			case <-b1open:
				log.Info("button #1 open")
				common.SendBtn(common.PadBtn_A, 0)
			case <-b2press:
				log.Info("button #2 pressed")
				common.SendBtn(common.PadBtn_B, 1)
			case <-b2open:
				log.Info("button #2 open")
				common.SendBtn(common.PadBtn_B, 0)
			case <-b3press:
				log.Info("button #3 pressed")
				common.SendBtn(common.PadBtn_Y, 1)
			case <-b3open:
				log.Info("button #3 open")
				common.SendBtn(common.PadBtn_Y, 0)
			case <-b4press:
				log.Info("button #4 pressed")
				common.SendBtn(common.PadBtn_X, 1)
			case <-b4open:
				log.Info("button #4 open")
				common.SendBtn(common.PadBtn_X, 0)
			case <-b5press:
				log.Info("button #5 pressed")
				common.SendBtn(common.PadBtn_LB, 1)
			case <-b5open:
				log.Info("button #5 open")
				common.SendBtn(common.PadBtn_LB, 0)
			case <-b6press:
				log.Info("button #6 pressed")
				common.SendBtn(common.PadBtn_RB, 1)
			case <-b6open:
				log.Info("button #6 open")
				common.SendBtn(common.PadBtn_RB, 0)
			// case <-b7press:
			// 	log.Info("button #7 pressed")
			// 	input := []byte{0, 0, 0}
			// 	input[2] = 1
			// 	input[1] = 2
			// 	conn.Write([]byte(input))
			// case <-b7open:
			// 	log.Info("button #7 open")
			// 	input := []byte{0, 0, 0}
			// 	input[2] = 0
			// 	input[1] = 2
			// 	conn.Write([]byte(input))
			// case <-b8press:
			// 	log.Info("button #8 pressed")
			// 	input := []byte{0, 0, 0}
			// 	input[2] = 1
			// 	input[1] = 3
			// 	conn.Write([]byte(input))
			// case <-b8open:
			// 	log.Info("button #8 open")
			// 	input := []byte{0, 0, 0}
			// 	input[2] = 0
			// 	input[1] = 3
			// 	conn.Write([]byte(input))
			case <-b9press:
				log.Info("button #9 pressed")
				common.SendBtn(common.PadBtn_Select, 1)
			case <-b9open:
				log.Info("button #9 open")
				common.SendBtn(common.PadBtn_Select, 0)
			case <-b10press:
				log.Info("button #10 pressed")
				common.SendBtn(common.PadBtn_Start, 1)
			case <-b10open:
				log.Info("button #10 open")
				common.SendBtn(common.PadBtn_Start, 0)
			case <-b11press:
				log.Info("button #11 pressed")
				common.SendBtn(common.PadBtn_Xbox, 1)
			case <-b11open:
				log.Info("button #11 open")
				common.SendBtn(common.PadBtn_Xbox, 0)
			case <-b12press:
				log.Info("button #12 pressed")
				common.SendBtn(common.PadBtn_L3, 1)
			case <-b12open:
				log.Info("button #12 open")
				common.SendBtn(common.PadBtn_L3, 0)
			case <-b13press:
				log.Info("button #13 pressed")
				common.SendBtn(common.PadBtn_R3, 1)
			case <-b13open:
				log.Info("button #13 open")
				common.SendBtn(common.PadBtn_L3, 0)
			case p := <-pan1X:
				hpos := p.(AxisEvent)
				// val := float32(0)
				// if hpos.V < 0 {
				// 	val = 32767 * hpos.V
				// } else {
				// 	val = 32768 * hpos.V
				// }
				// val = -val
				val := common.Map(float64(hpos.V), -1, 1, 32767, -32768)
				log.Info("pan #1 x moved to:", hpos.V)
				common.SendBtn(common.PadBtn_LAY, int32(val))
			case p := <-pan2X:
				hpos := p.(AxisEvent)
				// val := float32(0)
				// val = 32767 * hpos.V
				val := common.Map(float64(hpos.V), -1, 1, -32768, 32767)
				log.Info("pan #2 x moved to:", hpos.V)
				common.SendBtn(common.PadBtn_RAX, int32(val))
			case p := <-pan3X:
				hpos := p.(AxisEvent)
				val := common.Map(float64(hpos.V), -1, 1, 0, 255)
				log.Info("pan #3 x moved to:", hpos.V)
				common.SendBtn(common.PadBtn_R2, int32(val))
			case p := <-pan4X:
				hpos := p.(AxisEvent)
				log.Info("pan #4 x moved to:", hpos.V)
				if hpos.V == -1 {
					common.SendBtn(common.PadBtn_DUP, 1)
				} else if hpos.V == 0 {
					common.SendBtn(common.PadBtn_DUP, 0)
					common.SendBtn(common.PadBtn_DDOWN, 0)
				} else if hpos.V == 1 {
					common.SendBtn(common.PadBtn_DDOWN, 1)
				}
			case p := <-pan1Y:
				hpos := p.(AxisEvent)
				// val := float32(0)
				// if hpos.V < 0 {
				// 	val = 32768 * hpos.V
				// } else {
				// 	val = 32767 * hpos.V
				// }
				val := common.Map(float64(hpos.V), -1, 1, -32768, 32767)
				log.Info("pan #1 y moved to:", hpos.V)
				common.SendBtn(common.PadBtn_LAX, int32(val))
			case p := <-pan2Y:
				hpos := p.(AxisEvent)
				log.Info("pan #2 y moved to:", hpos.V)
				val := common.Map(float64(hpos.V), -1, 1, 0, 255)
				common.SendBtn(common.PadBtn_L2, int32(val))
			case p := <-pan3Y:
				hpos := p.(AxisEvent)
				// val := float32(0)
				// if hpos.V < 0 {
				// 	val = 32767 * hpos.V
				// } else {
				// 	val = 32768 * hpos.V
				// }
				// val = -val
				val := common.Map(float64(hpos.V), -1, 1, 32767, -32768)
				log.Info("pan #3 y moved to:", hpos.V)
				common.SendBtn(common.PadBtn_RAY, int32(val))
			case p := <-pan4Y:
				hpos := p.(AxisEvent)
				log.Info("pan #4 y moved to:", hpos.V)
				if hpos.V == -1 {
					common.SendBtn(common.PadBtn_DLEFT, 1)
				} else if hpos.V == 0 {
					common.SendBtn(common.PadBtn_DLEFT, 0)
					common.SendBtn(common.PadBtn_DRIGHT, 0)
				} else if hpos.V == 1 {
					common.SendBtn(common.PadBtn_DRIGHT, 1)
				}
			}
		}
	}()

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-ch)
}
