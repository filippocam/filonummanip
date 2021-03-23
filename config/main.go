package config

import (
	"fmt"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	// read yaml file
	readYamlConfig()

	// read env file
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// Print
	fmt.Println("---------- ENV file ----------")
	fmt.Printf("%#v %s", config, "\n")

}