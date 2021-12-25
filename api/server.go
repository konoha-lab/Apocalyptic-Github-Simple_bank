package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	handler Handler
	Router  *gin.Engine
}

func NewServer(handler Handler) *Server {
	server := &Server{handler: handler}
	router := gin.Default()

	// TODO: add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.Router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
