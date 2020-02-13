package rest

import "github.com/gin-gonic/gin"

/*type Server struct {
	port *string
}*/

func NewRouter() *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/api/v1/tm")

	//v1.POST("/tasks", rest.CreateTask)
	/*v1.GET("/tasks", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})*/

	v1.POST("/add", AddData)
	v1.GET("/get", AddData)
	return router
}

/*func NewHttpServer(p *string) *Server {
	return &Server{
		port:p,
	}
}*/
