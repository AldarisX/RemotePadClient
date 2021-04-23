package main

import (
	"RemotePad/src/common"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var testMode = flag.String("test", "0", "测试模式")

func main() {
	// var xbox bool = false
	flag.Parse()

	common.Connect()

	time.Sleep(time.Duration(4) * time.Second)

	common.UseXbox360()
	// xbox = true

	time.Sleep(time.Duration(5) * time.Second)

	common.SendBtn(common.PadBtn_A, 1)

	common.SendBtn(common.PadBtn_Y, 1)

	common.SendBtn(common.PadBtn_LAX, 16383)

	common.SendBtn(common.PadBtn_L2, 127)

	common.SendBtn(common.PadBtn_DUP, 1)

	time.Sleep(time.Duration(3) * time.Second)

	// common.UseDS4()
	// common.UseXbox360()

	common.SendBtn(common.PadBtn_DS4TouchPad, 1)

	time.Sleep(time.Duration(1) * time.Second)

	common.SendBtn(common.PadBtn_DS4TouchPad, 0)

	common.SendBtn(common.PadBtn_DUP, 1)

	common.SendBtn(common.PadBtn_DDOWN, 1)

	if *testMode == "1" {
		// go func() {
		// 	for {
		// 		if rand.Intn(100) < 10 {
		// 			if xbox {
		// 				common.UseDS4()
		// 			} else {
		// 				common.UseXbox360()
		// 			}
		// 		}
		// 		time.Sleep(time.Duration(10) * time.Second)
		// 	}
		// }()

		go func() {
			for {
				time.Sleep(time.Duration(20) * time.Millisecond)
				common.SendBtn(common.PadBtn_DUP, 1)
				time.Sleep(time.Duration(20) * time.Millisecond)
				common.SendBtn(common.PadBtn_DDOWN, 0)
				time.Sleep(time.Duration(20) * time.Millisecond)
				common.SendBtn(common.PadBtn_DLEFT, 1)
				time.Sleep(time.Duration(20) * time.Millisecond)
				common.SendBtn(common.PadBtn_DRIGHT, 0)
				time.Sleep(time.Duration(20) * time.Millisecond)
				common.SendBtn(common.PadBtn_DUP, 0)
				time.Sleep(time.Duration(20) * time.Millisecond)
				common.SendBtn(common.PadBtn_DDOWN, 1)
				time.Sleep(time.Duration(20) * time.Millisecond)
				common.SendBtn(common.PadBtn_DLEFT, 0)
				time.Sleep(time.Duration(20) * time.Millisecond)
				common.SendBtn(common.PadBtn_DRIGHT, 1)
			}
		}()

		go func() {
			for {
				time.Sleep(time.Duration(2) * time.Millisecond)
				common.SendBtn(common.PadBtn_LAX, int32(rand.Intn(65535)-32768))
				time.Sleep(time.Duration(2) * time.Millisecond)
				time.Sleep(time.Duration(2) * time.Millisecond)
				common.SendBtn(common.PadBtn_LAY, int32(rand.Intn(65535)-32768))
				time.Sleep(time.Duration(2) * time.Millisecond)
				common.SendBtn(common.PadBtn_R2, int32(rand.Intn(255)))
				time.Sleep(time.Duration(2) * time.Millisecond)
			}
		}()
	}

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-ch)
}
