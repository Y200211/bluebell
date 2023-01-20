package controller

import (
	"bluebell/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不轻易吧服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
