package core

import (
	"github.com/gin-gonic/gin"
)


func (b *Chain) Run(port string) error {
	router := gin.Default()
	v1 := router.Group("/api/v1/ledger")
	v1.POST("", b.AddData)
	v1.GET("", b.ReadData)

	err := router.Run(port)
	if err != nil {
		return err
	}
	return nil
}
