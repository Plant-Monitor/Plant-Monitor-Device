package pch

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"testing"
	"time"
)

func TestGetReadings(t *testing.T) {
	//setup := func() {
		//err := godotenv.Load("../.env")
		//if err != nil {
			//return
		//}
	//}

	/*
		Hardware setup: A peripheral must be on CS0 for this function to work.
	*/
	t.Run("Read a specific chip until a key is pressed", func(t *testing.T) {
		setupPCH()
		scanner := bufio.NewScanner(os.Stdin)
		done := false

		fmt.Println("Press any key to stop output at any time")
		time.Sleep(3 * time.Second)

		for !done {
			// Read from moisture sensor
			readings := GetPCHClientInstance().GetReadings()
			moisture := readings["moisture"]
			fmt.Println("Reading on peripheral 0: ", moisture.Level, moisture.Unit)

			// Check if there's input waiting on stdin
			if scanner.Scan() {
				// A key was pressed, so exit the loop
				done = true
			}
		}

		fmt.Println("Loop exited")
	})
}

//func TestRead(t *testing.T) {
	//setup := func() {
		//err := godotenv.Load("../.env")
		//if err != nil {
			//return
		//}
	//}

	///*
		//Hardware setup: A peripheral must be on CS pins specified in metricPeripheralNumberMapping.json
	//*/
	//t.Run("Read all peripherals repeatedly until a key is pressed", func(t *testing.T) {
		//setup()
		//scanner := bufio.NewScanner(os.Stdin)
		//done := false

		//fmt.Println("Press any key to stop output at any time")
		//time.Sleep(3 * time.Second)

		//for !done {
			//// Read from all peripheral
			//fmt.Println("Readings Collection: ", GetGpioClientInstance().Read())

			//// Check if there's input waiting on stdin
			//if scanner.Scan() {
				//// A key was pressed, so exit the loop
				//done = true
			//}
		//}

		//fmt.Println("Loop exited")
	//})
//}

