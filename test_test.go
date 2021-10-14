package main

import (
	"fmt"
	"testing"

	"github.com/renatormc/kvmusb/config"
	"github.com/renatormc/kvmusb/services"
)

func TestMain(t *testing.T) {

	config.InitWithExeDir("/home/renato/src/kvmusb/exedir")
	usbs, err := services.ListUsbs()
	if err != nil {
		panic(err)
	}
	for _, usb := range usbs {
		fmt.Println(usb.ID)
	}
}
