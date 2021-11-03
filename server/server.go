package server

import (
	"log"
	"web-api-go/server/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigureRoutes(s.server)
	log.Printf("Server is running on port %s", s.port)
	log.Fatal(router.Run(":" + s.port))
}
