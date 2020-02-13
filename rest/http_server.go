package rest

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/api/v1/tm")

	v1.POST("/add", AddData)
	v1.GET("/get", AddData)
	return router
}

