package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	//1. 生成postID
	p.ID = snowflake.GenID()
	//2. 保存到数据库
	return mysql.CreatePost(p)
}

func GetPostByID(pid int64) (data *models.ApiPostDetail, err error) {

	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql GetPostByID(pid int64) failed", zap.Error(err))
		return
	}
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql GetUserByID(post.AuthorID) failed",
			zap.Int64("author_id", post.AuthorID),
			zap.Error(err))
		return
	}
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql GetCommunityDetailByID(post.CommunityID) failed",
			zap.Int64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return

}
