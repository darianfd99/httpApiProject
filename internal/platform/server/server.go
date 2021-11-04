package server

import (
	"fmt"
	"log"

	"github.com/darianfd99/httpApiProject/internal/platform/server/handler/courses"
	"github.com/darianfd99/httpApiProject/internal/platform/server/handler/health"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port uint) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("#{host}:#{port}"),
	}

	srv.registerRoutes()

	return srv
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler())
}

func (s *Server) Run() error {
	log.Println("Server running on:", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}
