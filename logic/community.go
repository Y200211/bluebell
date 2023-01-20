package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.Community()
}
