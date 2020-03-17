package core

import "github.com/gin-gonic/gin"

func NewRouter(b *Chain) *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/api/v1/ledger")

	v1.POST("", b.AddData)
	v1.GET( "", b.GetData)
	return router
}

