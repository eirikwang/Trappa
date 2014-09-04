package main

import (
	"github.com/hybridgroup/gobot/platforms/beaglebone"
	"github.com/hybridgroup/gobot"
	"github.com/eirikwang/Trappa/drivers"
	"fmt"
	"time"
)

type localBB struct {
	beaglebone.BeagleboneAdaptor
}

func main() {
	gbot := gobot.NewGobot()
	beagleboneAdaptor := beaglebone.NewBeagleboneAdaptor("beaglebone")
	temp := drivers.NewMCP9808Driver(beagleboneAdaptor, "MCP9808 Temperature sensor")
	work := func() {
		gobot.Every(1*time.Second, func() {
				fmt.Println("Temp", temp.Temperature.FloatString(4))
			})
	}

	robot := gobot.NewRobot("servoBot",
		[]gobot.Connection{beagleboneAdaptor},
		[]gobot.Device{temp},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
