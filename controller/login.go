package controller

import (
	"ChsSeltApi/utils/common"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/sijms/go-ora/v2"
	"net/http"
)

// 用户API管理对象
var Login = new(loginCtl)

type loginCtl struct{}

func (c *loginCtl) Login(ctx *gin.Context) {
	id, _ := ctx.GetQuery("id")
	fmt.Println(id)
	//body := ctx.Request.Body
	var json common.JsonResult
	//json := common.JsonResult{}
	ctx.ShouldBindJSON(&json)
	fmt.Println(json.Msg)
	// 已登录
	ctx.JSONP(http.StatusOK, common.JsonResult{
		Code:  0,
		Msg:   "查询成功",
		Data:  "",
		Count: 0,
	})
}
