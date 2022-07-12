package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/mayowa1305/simpleBank/db/sqlc"
)

//server serves HTTP requests for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

//newserver creates a new http server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	server.router = router
	return server
}

//start runs the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
