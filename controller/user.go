package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	//1.获取参数，参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUpHandler with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": removeTopStruct(errs.Translate(trans)),
		//})
		return
	}
	//手动对请求参数进行详细的业务规则校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
	//	zap.L().Error("SignUpHandler with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}
	fmt.Println(p)

	//2.业务处理

	if err := logic.SignUp(p); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return

		fmt.Println(err)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)

}

func LoginHandler(c *gin.Context) {
	//1 获取参数
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("LoginHandler with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": removeTopStruct(errs.Translate(trans)),
		//})
		return
	}
	//2 逻辑处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		fmt.Println(err)
		return
	}
	//3 返回参数
	ResponseSuccess(c, nil)
}
