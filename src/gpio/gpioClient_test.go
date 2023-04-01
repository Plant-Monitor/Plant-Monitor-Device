package gpio

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestReadDigitalValue(t *testing.T) {
	setup := func() {
		err := godotenv.Load("../.env")
		if err != nil {
			return
		}
	}

	/*
		Hardware setup: A peripheral must be on CS0 for this function to work.
	*/
	t.Run("Read a specific chip until a key is pressed", func(t *testing.T) {
		setup()
		scanner := bufio.NewScanner(os.Stdin)
		done := false

		for !done {
			// Read from peripheral 0
			fmt.Println("Reading on peripheral 0: ", GetGpioClientInstance().readDigitalValue(0))

			// Check if there's input waiting on stdin
			if scanner.Scan() {
				// A key was pressed, so exit the loop
				done = true
			}
		}

		fmt.Println("Loop exited")
	})
}

func TestRead(t *testing.T) {
	setup := func() {
		err := godotenv.Load("../.env")
		if err != nil {
			return
		}
	}

	/*
		Hardware setup: A peripheral must be on CS pins specified in metricPeripheralNumberMapping.json
	*/
	t.Run("Read all peripherals repeatedly until a key is pressed", func(t *testing.T) {
		setup()
		scanner := bufio.NewScanner(os.Stdin)
		done := false

		for !done {
			// Read from all peripheral
			fmt.Println("Reading on peripheral 0: ", GetGpioClientInstance().Read())

			// Check if there's input waiting on stdin
			if scanner.Scan() {
				// A key was pressed, so exit the loop
				done = true
			}
		}

		fmt.Println("Loop exited")
	})
}
