package core

import (
	"github.com/gin-gonic/gin"
	)

type Data struct {
	Field string `json:"field"`
}

func (b *Chain) AddData(cxt *gin.Context) {

	//og
	b.WriteDataRest(cxt)

	//variable req = Data{}
	//err := cxt.Bind(&req)
	//if err != nil {
	//	cxt.JSON(400, err.Error())
	//	return
	//}


	//cxt.JSON(200,b)
}

func (b *Chain) GetData(cxt *gin.Context) {
	b.WriteDataRest(cxt)
}
