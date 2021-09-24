package handler

import (
	"github.com/dvnhanh/thewolddata/internal/core/port"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HTTPServer interface {
	Begin(address string) error
}

func NewHTTPServer(svc port.TheworlddataService) HTTPServer {
	return &httpServer{
		svc: svc,
	}
}

type httpServer struct {
	svc port.TheworlddataService
}

// middleware
func authen() gin.HandlerFunc {
	return func(c *gin.Context) {
		// something...
	}
}

// Begin is used to start the Http
func (p *httpServer) Begin(address string) error {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"HEAD", "GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Access-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}), authen())

	p.setup(router.Group("thewolddata"))
	return router.Run(address)
}

// setup API
func (p *httpServer) setup(router *gin.RouterGroup) {
	router.GET("/ping", p.ping)
	router.POST("/register", p.register)
}
