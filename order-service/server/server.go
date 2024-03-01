// Package server maintains server components
package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port  string
	Debug bool
}

type Server struct {
	*gin.Engine
}

func New() Server {
	server := gin.Default()
	return Server{server}
}

func Start(e *Server, cfg *Config) {

	s := &http.Server{
		Addr:    cfg.Port,
		Handler: e.Engine,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		if err := s.Close(); err != nil {
			log.Println("failed To ShutDown Server", err)
		}
		log.Println("Shut Down Server")
	}()

	if err := s.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server Closed After Interruption")
		} else {
			log.Println("Unexpected Server Shutdown. err:", err)
		}
	}
}
