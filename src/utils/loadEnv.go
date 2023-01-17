package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	envFilename := ".env"

	if len(os.Args) > 1 {
		envFilename = fmt.Sprintf(".env.%s", os.Args[len(os.Args)-1])
	}

	err := godotenv.Load(fmt.Sprintf("../../%s", envFilename))
	if err != nil {
		log.Fatalf("%s file does not exist", envFilename)
	}
}