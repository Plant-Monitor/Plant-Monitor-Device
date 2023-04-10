package pch

import (
<<<<<<< HEAD
	//"bufio"
	"fmt"
	//"os"
	"testing"
	"time"
	//"pcs/actions"
	//"pcs/models"
	//"periph.io/x/conn/v3/gpio"
)

func TestPumpDriver(t *testing.T) {
=======
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestGetReadings(t *testing.T) {
>>>>>>> beddb5414abaccd00fce4616dbdb37c636e69636
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
<<<<<<< HEAD
		// Read from moisture sensor
			GetPCHClientInstance().PerformActuations()
			fmt.Printf("*** MOISTURE WAS REGULATED ***\n")
		
		
			time.Sleep(time.Second)
			// Check if there's input waiting on stdin
			if scanner.Scan() {
				// A key was pressed, so exit the loop
				done = true
			}
		
	}
		fmt.Println("Loop exited")
	})
}



func TestGetReadings(t *testing.T){
	t.Run("Read a specific chip until a key is pressed", func(t *testing.T) {
		setupPCH()
		scanner := bufio.NewScanner(os.Stdin)
		done := false

		fmt.Println("Press any key to stop output at any time")
		time.Sleep(3 * time.Second)

		for !done {
=======
>>>>>>> beddb5414abaccd00fce4616dbdb37c636e69636
			// Read from moisture sensor
			readings := GetPCHClientInstance().GetReadings()
			for metric, reading := range readings{
				fmt.Printf("%s: %.2f %s \n",metric, reading.Level, reading.Unit)
				//fmt.Printf("\n")
			}
			//fmt.Printf("Readings: %+v\n", readings)
			//fmt.Println("Reading on peripheral 0: ", moisture.Level, moisture.Unit)
			fmt.Printf("*******\n")
			time.Sleep(time.Second)
			// Check if there's input waiting on stdin
			if scanner.Scan() {
				// A key was pressed, so exit the loop
				done = true
			}
		}

		fmt.Println("Loop exited")
	})
<<<<<<< HEAD
	
=======
>>>>>>> beddb5414abaccd00fce4616dbdb37c636e69636
}

//func TestRead(t *testing.T) {
//setup := func() {
//err := godotenv.Load("../.env")
//if err != nil {
//return
//}
//}

<<<<<<< HEAD

//Hardware setup: A peripheral must be on CS pins specified in metricPeripheralNumberMapping.json
//
=======
///*
//Hardware setup: A peripheral must be on CS pins specified in metricPeripheralNumberMapping.json
//*/
>>>>>>> beddb5414abaccd00fce4616dbdb37c636e69636
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
