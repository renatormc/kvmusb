package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/renatormc/kvmusb/config"
	"github.com/renatormc/kvmusb/server/routes"
)

type Server struct {
	server *gin.Engine
}

func NewServer() Server {
	server := Server{
		server: gin.Default(),
	}
	return server
}

func (s *Server) Run() {
	config := config.GetConfig()
	router := routes.ConfigRoutes(s.server)

	log.Printf("Server running at port: %s", config.Port)
	log.Fatal(router.Run(fmt.Sprintf(":%s", config.Port)))
}
