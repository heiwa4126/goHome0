package main

import (
	"fmt"
	"github.com/heiwa4126/goHome0/config"
	"log"
)

// example code
func ex1() error {

	// get config path
	cnfFile, _, err := config.GetPath()
	if err != nil {
		return err
	}
	fmt.Printf("config file='%s'\n", cnfFile)

	// write config
	conf := config.Config{APIKey: "api-key(test)", Endpoint: "endpoint(dummy)"}
	if err = conf.Write(); err != nil {
		return err
	}

	// read config
	conf2, err := config.Read()
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", *conf2)

	// over
	return nil
}

func main() {
	fmt.Println("Hello world!")
	if err := ex1(); err != nil {
		log.Fatal(err)
	}
}
