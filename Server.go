package main

import (
	"github.com/hybridgroup/gobot/platforms/beaglebone"
	"github.com/hybridgroup/gobot"
	"github.com/eirikwang/Trappa/drivers"
	"fmt"
	"time"
)

func main() {
	gbot := gobot.NewGobot()
	beagleboneAdaptor := beaglebone.NewBeagleboneAdaptor("beaglebone")
	temp := drivers.NewMCP9828Driver(beagleboneAdaptor, drivers.IO_port)
	work := func() {
		gobot.Every(1*time.Second, func() {
				fmt.Println("Temp", temp.Temperature)
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
