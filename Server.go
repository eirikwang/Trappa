package main

import (
	"github.com/hybridgroup/gobot/platforms/beaglebone"
	"github.com/hybridgroup/gobot"
	"github.com/eirikwang/Trappa/drivers"
	_"fmt"
	"time"

	//"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/api"
)

func main() {
	gbot := gobot.NewGobot()
	beagleboneAdaptor := beaglebone.NewBeagleboneAdaptor("beaglebone")

	//lux := drivers.NewTSL2561Driver(beagleboneAdaptor, "TSL2561 motion sensor")
    temperature:= drivers.NewMCP9808Driver(beagleboneAdaptor, "temp")
    motion := gpio.NewDirectPinDriver(beagleboneAdaptor, "motion", "P9_12")
	sonar := gpio.NewAnalogSensorDriver(beagleboneAdaptor, "sonar", "P9_39")


	work := func() {
		gobot.Every(1*time.Second, func() {
				//fmt.Println("Bevegelse: ", motion.DigitalRead());
				//fmt.Println("Temp: ", temperature.Temperature);
				//fmt.Println("Sonar: ", sonar.Read())
			})
	}

	robot := gobot.NewRobot("trappa",
		[]gobot.Connection{beagleboneAdaptor},
		[]gobot.Device{sonar, motion, temperature},
		work,
	)
	gbot.AddRobot(robot)
	server := api.NewAPI(gbot)
	server.Port = "4040"
	server.Start()

	gbot.Start()
}
