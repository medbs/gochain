package core

import (
	"github.com/gin-gonic/gin"
	)

type Data struct {
	Field string `json:"field"`
}

func (b *Chain) AddData(cxt *gin.Context) {
	b.WriteDataRest(cxt)
}


func (b *Chain) ReadData(cxt *gin.Context) {
	b.ReadDataRest(cxt)
}
