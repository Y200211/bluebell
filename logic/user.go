package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

// SignUp 存放业务逻辑代码
func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户存不存在

	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 生成UID
	userID := snowflake.GenID()
	// 密码加密
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 保存进数据库

	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	// user.UserID
	return jwt.GenToken(user.UserID, user.Username)

}
func GetCommunityList() (data string, err error) {

}
