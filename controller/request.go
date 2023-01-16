package controller

import (
	"bluebell/middlewares"
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(middlewares.CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
