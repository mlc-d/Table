package api

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	port = os.Getenv("port")
)

type Server struct {
	server *echo.Echo
	port   int
}

func GetServer() Server {
	// NOTE: there's no need to enforce a singleton pattern,
	// since this function will be called only once from the
	// entrypoint of the program
	if port == "" {
		port = "1313"
	}

	port, err := strconv.Atoi(port)
	if err != nil {
		log.Panicln("PORT env variable must be a number")
	}

	return Server{
		server: echo.New(),
		port:   port,
	}
}

// Start listens on the [Server.port]
func (s *Server) Start() error {
	return s.server.Start(fmt.Sprintf(":%d", s.port))
}
