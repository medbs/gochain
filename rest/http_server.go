package rest

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/api/v1/ledger")

	v1.POST("", AddData)
	v1.GET( "", ReadData)
	return router
}

