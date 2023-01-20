package mysql

import (
	"bluebell/models"
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

func Community() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
		fmt.Println(err)
	}
	return
}
