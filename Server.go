package main

import (
	"github.com/hybridgroup/gobot/platforms/beaglebone"
	"github.com/hybridgroup/gobot"
	"github.com/eirikwang/Trappa/drivers"
	"fmt"
	"time"

	//"github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
	gbot := gobot.NewGobot()
	beagleboneAdaptor := beaglebone.NewBeagleboneAdaptor("beaglebone")

	lux := drivers.NewTSL2561Driver(beagleboneAdaptor, "TSL2561 motion sensor")
		//"temperature" : drivers.NewMCP9808Driver(beagleboneAdaptor, "MCP9808 Temperature sensor"),
		//"motion" : gpio.NewDirectPinDriver(beagleboneAdaptor, "motion", "P9_12"),

	work := func() {
		gobot.Every(1*time.Second, func() {
				fmt.Println("Lux: ", lux.Lux);
			})
	}

	robot := gobot.NewRobot("servoBot",
		[]gobot.Connection{beagleboneAdaptor},
		[]gobot.Device{lux},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
