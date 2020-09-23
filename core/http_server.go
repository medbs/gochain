package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func (b *Chain) Run(port string) error {
	router := gin.Default()
	router.Use(CORS)
	v1 := router.Group("/api/v1/ledger")
	v1.POST("", b.AddData)
	v1.GET("", b.ReadData)

	err := router.Run(port)
	if err != nil {
		return err
	}
	return nil
}


// CORS Middleware
func CORS(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {

		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}
