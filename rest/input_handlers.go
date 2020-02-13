package rest

import (
	"github.com/gin-gonic/gin"
)

type Data struct {
	Field  string `json:"field"`
}


func AddData(cxt *gin.Context){
	var req = Data{}
	err := cxt.Bind(&req)
	if err != nil {
		cxt.JSON(400, err.Error())
		return
	}
	cxt.JSON(200,req)
}


func ReadData(cxt *gin.Context){

}