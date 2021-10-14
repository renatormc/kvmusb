package main

import (
	"github.com/renatormc/kvmusb/config"
	"github.com/renatormc/kvmusb/server"
)

func main() {
	s := server.NewServer()
	config.Init()
	s.Run()
}
