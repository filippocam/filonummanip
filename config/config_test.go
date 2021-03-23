package config

import (
	"fmt"
	"log"
	"testing"
)

func TestEnv(m *testing.T) {
	config, err := LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// Print
	fmt.Println("---------- ENV file ----------")
	fmt.Printf("%#v %s", config, "\n")
}