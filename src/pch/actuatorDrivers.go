package pch

import (
	"periph.io/x/conn/v3/gpio"
	"time"
	//"fmt"
)

type actuatorDriver func()

const (
	timeOn 	= 2 * time.Second
	timeOff = 4 * time.Second
)

func pumpDriver(){
	moistureActPin.Out(gpio.High)
	time.Sleep(timeOn)
	moistureActPin.Out(gpio.Low)
	time.Sleep(4 * time.Second)
	//fmt.Printf("*** MOISTURE WAS REGULATED ***\n")
}


//readings := GetPCHClientInstance().GetReadings()
//snap := models.BuildSnapshot(readings)
//strategy := actions.newNeededActionRegulationStrategy("moisture", time.Duration(30 * time.Second))
//decision, critRange, actType := strategy.decide(&snap)
//if (decision && critRange == CRITICAL_LOW) {
