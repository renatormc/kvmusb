package main

import (
	"testing"

	"github.com/renatormc/kvmusb/config"
	"github.com/renatormc/kvmusb/services"
)

func TestMain(t *testing.T) {

	config.InitWithExeDir("/home/renato/src/kvmusb/exedir")
	_, err := services.ListRunningVms()
	if err != nil {
		panic(err)
	}
}
