package controller

import (
	"ChsSeltApi/utils/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户API管理对象
var Login = new(loginCtl)

type loginCtl struct{}

func (c *loginCtl) Login(ctx *gin.Context) {
	// 已登录
	ctx.JSONP(http.StatusOK, common.JsonResult{
		Code:  0,
		Msg:   "查询成功",
		Data:  "",
		Count: "",
	})
}
