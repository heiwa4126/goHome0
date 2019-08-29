package main

import (
	"fmt"
	"github.com/heiwa4126/goHome0/config"
	"log"
)

func ex1() error {

	cnfFile, _, err := config.GetPath()
	if err != nil {
		return err
	}
	fmt.Printf("config file='%s'\n", cnfFile)
	return nil
}

func main() {
	fmt.Println("世界の皆さん、こんにちは!")
	err := ex1()
	if err != nil {
		log.Fatal(err)
	}
}
