package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	//1. 生成postID
	p.ID = snowflake.GenID()
	//2. 保存到数据库
	return mysql.CreatePost(p)
}

func GetPostByID(pid int64) (data *models.Post, err error) {
	return mysql.GetPostByID(pid)

}
