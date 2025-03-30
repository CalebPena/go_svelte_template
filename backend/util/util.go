package util

import (
	"fmt"
	"log"
	"os"
)

var IsDebug = GetEnvironmentVariable("DEBUG") == "true"

func GetEnvironmentVariable(name string) string {
	variable := os.Getenv(name)

	if variable == "" {
		log.Fatal(fmt.Sprintf("Please set the '%s' environment variable", name))
	}

	return variable
}
