package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.Community()
}
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.CommunityDetail(id)
}
