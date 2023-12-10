package api

import (
	db "github.com/2281501756/BackendClass/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Db     db.Store
	Router *gin.Engine
}

func NewServer(db db.Store) *Server {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(200, "ok")
	})
	return &Server{db, router}
}
