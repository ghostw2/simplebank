package api

import (
	db "menribardhi/micro-go-psql/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.listAccounts)
	router.GET("/accounts/:id", server.getAccountById)

	server.router = router
	return server
}

func (s *Server) Start(adress string) error {
	return s.router.Run(adress)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
