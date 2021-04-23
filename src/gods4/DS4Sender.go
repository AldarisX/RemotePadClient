package main

import (
	"RemotePad/src/common"
	"flag"
	"os"
	"os/signal"
	"syscall"

	log "github.com/YoungPioneers/blog4go"
	"github.com/kpeu3i/gods4"
)

func main() {
	flag.Parse()
	common.Connect()

	var err error

	// Find all controllers connected to your machine via USB or Bluetooth
	controllers := gods4.Find()
	if len(controllers) == 0 {
		panic("No connected DS4 controllers found")
	}

	// Select first controller from the list
	controller := controllers[0]

	// Connect to the controller
	err = controller.Connect()
	if err != nil {
		panic(err)
	}

	log.Infof("* Controller #1 | %-10s | name: %s, connection: %s\n", "Connect", controller, controller.ConnectionType())

	// Disconnect controller when a program is terminated
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		err := controller.Disconnect()
		if err != nil {
			panic(err)
		}
		log.Infof("* Controller #1 | %-10s | bye!\n", "Disconnect")
	}()

	// Register callback for "BatteryUpdate" event
	// controller.On(gods4.EventBatteryUpdate, func(data interface{}) error {
	// 	battery := data.(gods4.Battery)
	// 	log.Infof("* Controller #1 | %-10s | capacity: %v%%, charging: %v, cable: %v\n",
	// 		"Battery",
	// 		battery.Capacity,
	// 		battery.IsCharging,
	// 		battery.IsCableConnected,
	// 	)

	// 	return nil
	// })

	// Register callback for "CrossPress" event
	controller.On(gods4.EventCrossPress, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: press\n", "Cross")
		common.SendBtn(common.PadBtn_A, 1)
		return nil
	})

	// Register callback for "CrossRelease" event
	controller.On(gods4.EventCrossRelease, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: release\n", "Cross")
		common.SendBtn(common.PadBtn_A, 0)
		return nil
	})

	controller.On(gods4.EventCirclePress, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: press\n", "Circle")
		common.SendBtn(common.PadBtn_B, 1)
		return nil
	})

	controller.On(gods4.EventCircleRelease, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: release\n", "Circle")
		common.SendBtn(common.PadBtn_B, 0)
		return nil
	})

	controller.On(gods4.EventSquarePress, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: press\n", "Square")
		common.SendBtn(common.PadBtn_X, 1)
		return nil
	})

	controller.On(gods4.EventSquareRelease, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: release\n", "Square")
		common.SendBtn(common.PadBtn_X, 0)
		return nil
	})

	controller.On(gods4.EventTrianglePress, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: press\n", "Triangle")
		common.SendBtn(common.PadBtn_Y, 1)
		return nil
	})

	controller.On(gods4.EventTriangleRelease, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: release\n", "Triangle")
		common.SendBtn(common.PadBtn_Y, 0)
		return nil
	})

	controller.On(gods4.EventL1Press, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: press\n", "L1")
		common.SendBtn(common.PadBtn_LB, 1)
		return nil
	})

	controller.On(gods4.EventL1Release, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: release\n", "L1")
		common.SendBtn(common.PadBtn_LB, 0)
		return nil
	})

	controller.On(gods4.EventR1Press, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: press\n", "R1")
		common.SendBtn(common.PadBtn_RB, 1)
		return nil
	})

	controller.On(gods4.EventR1Release, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: release\n", "R1")
		common.SendBtn(common.PadBtn_RB, 0)
		return nil
	})

	controller.On(gods4.EventDPadUpPress, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: press\n", "DUP")
		common.SendBtn(common.PadBtn_DUP, 1)
		return nil
	})

	controller.On(gods4.EventDPadUpRelease, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: release\n", "DUP")
		common.SendBtn(common.PadBtn_DUP, 0)
		return nil
	})

	controller.On(gods4.EventDPadDownPress, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: press\n", "DDOWN")
		common.SendBtn(common.PadBtn_DDOWN, 1)
		return nil
	})

	controller.On(gods4.EventDPadDownRelease, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: release\n", "DDOWN")
		common.SendBtn(common.PadBtn_DDOWN, 0)
		return nil
	})

	controller.On(gods4.EventDPadLeftPress, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: press\n", "DLEFT")
		common.SendBtn(common.PadBtn_DLEFT, 1)
		return nil
	})

	controller.On(gods4.EventDPadLeftRelease, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: release\n", "DLEFT")
		common.SendBtn(common.PadBtn_DLEFT, 0)
		return nil
	})

	controller.On(gods4.EventDPadRightPress, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: press\n", "DRIGHT")
		common.SendBtn(common.PadBtn_DRIGHT, 1)
		return nil
	})

	controller.On(gods4.EventDPadRightRelease, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: release\n", "DRIGHT")
		common.SendBtn(common.PadBtn_DRIGHT, 0)
		return nil
	})

	controller.On(gods4.EventOptionsPress, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: press\n", "Options")
		common.SendBtn(common.PadBtn_Start, 1)
		return nil
	})

	controller.On(gods4.EventOptionsRelease, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: release\n", "Options")
		common.SendBtn(common.PadBtn_Start, 0)
		return nil
	})

	controller.On(gods4.EventSharePress, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: press\n", "Share")
		common.SendBtn(common.PadBtn_Select, 1)
		return nil
	})

	controller.On(gods4.EventShareRelease, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: release\n", "Share")
		common.SendBtn(common.PadBtn_Select, 0)
		return nil
	})

	controller.On(gods4.EventPSPress, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: press\n", "PS")
		common.SendBtn(common.PadBtn_Xbox, 1)
		return nil
	})

	controller.On(gods4.EventPSRelease, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: release\n", "PS")
		common.SendBtn(common.PadBtn_Xbox, 0)
		return nil
	})

	controller.On(gods4.EventL3Press, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: press\n", "L3")
		common.SendBtn(common.PadBtn_L3, 1)
		return nil
	})

	controller.On(gods4.EventL3Release, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: release\n", "L3")
		common.SendBtn(common.PadBtn_L3, 0)
		return nil
	})

	controller.On(gods4.EventR3Press, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: press\n", "R3")
		common.SendBtn(common.PadBtn_R3, 1)
		return nil
	})

	controller.On(gods4.EventR3Release, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: release\n", "R3")
		common.SendBtn(common.PadBtn_L3, 0)
		return nil
	})

	controller.On(gods4.EventL2Press, func(data interface{}) error {
		triger := data.(byte)
		log.Infof("* Controller #1 | %-10s | x: %v\n", "L2", triger)
		common.SendBtn(common.PadBtn_L2, int32(triger))
		return nil
	})

	controller.On(gods4.EventR2Press, func(data interface{}) error {
		triger := data.(byte)
		log.Infof("* Controller #1 | %-10s | x: %v\n", "R2", triger)
		common.SendBtn(common.PadBtn_R2, int32(triger))
		return nil
	})

	controller.On(gods4.EventLeftStickMove, func(data interface{}) error {
		stick := data.(gods4.Stick)
		log.Infof("* Controller #1 | %-10s | x: %v, y: %v\n", "LeftStick", stick.X, stick.Y)
		x := common.Map(float64(stick.X), 0, 255, -32768, 32767)
		y := common.Map(float64(stick.Y), 0, 255, 32767, -32768)

		common.SendBtn(common.PadBtn_LAX, int32(x))

		common.SendBtn(common.PadBtn_LAY, int32(y))

		return nil
	})

	// Register callback for "RightStickMove" event
	controller.On(gods4.EventRightStickMove, func(data interface{}) error {
		stick := data.(gods4.Stick)
		log.Infof("* Controller #1 | %-10s | x: %v, y: %v\n", "RightStick", stick.X, stick.Y)
		x := common.Map(float64(stick.X), 0, 255, -32768, 32767)
		y := common.Map(float64(stick.Y), 0, 255, 32767, -32768)

		common.SendBtn(common.PadBtn_RAX, int32(x))

		common.SendBtn(common.PadBtn_RAY, int32(y))

		return nil
	})

	controller.On(gods4.EventTouchpadPress, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: press\n", "Touchpad")
		common.SendBtn(common.PadBtn_DS4TouchPad, 1)
		return nil
	})

	controller.On(gods4.EventTouchpadRelease, func(data interface{}) error {
		log.Infof("* Controller #1 | %-10s | state: release\n", "Touchpad")
		common.SendBtn(common.PadBtn_DS4TouchPad, 0)
		return nil
	})

	controller.On(gods4.EventTouchpadSwipe, func(data interface{}) error {
		touchpad := data.(gods4.Touchpad)
		for _, swipe := range touchpad.Swipe {
			log.Infof("* Controller #1 | %-10s | x: %v, y: %v\n", "Touchpad", swipe.X, swipe.Y)
		}
		return nil
	})

	// // Enable left and right rumble motors
	// err = controller.Rumble(rumble.Both())
	// if err != nil {
	// 	panic(err)
	// }

	// // Enable LED (yellow) with flash
	// err = controller.Led(led.Yellow().Flash(50, 50))
	// if err != nil {
	// 	panic(err)
	// }

	// Start listening for controller events
	err = controller.Listen()
	if err != nil {
		panic(err)
	}
}
