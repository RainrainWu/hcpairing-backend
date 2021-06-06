package hcpairing

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
	AddListener()
	Start()
}

type server struct {
	router *gin.Engine
}

func NewServer() Server {
	instance := server{
		router: gin.Default(),
	}
	return &instance
}

func (s *server) AddListener() {
	s.router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func (s *server) Start() {
	s.router.Run()
}
