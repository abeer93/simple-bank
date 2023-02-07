package api

import (
	"github.com/gin-gonic/gin"
	"github.com/abeer93/simple-bank.git/db/sqlc"
)

type SreverType interface {
	Start(address string) error
}

type Server struct {
	store db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"Error": err.Error()}
}