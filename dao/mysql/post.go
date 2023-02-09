package mysql

import (
	"bluebell/models"
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post 
    		   (post_id, title, content, author_id, community_id)	
			   values (?, ?, ?, ?, ?)
			   `
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return

}

func GetPostByID(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select 
			   title, content, post_id, author_id, community_id, create_time 
			   from post
			   where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query post failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorInvalidID
		return
	}
	return
}

func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select title, content, post_id, author_id, community_id, create_time 
			   from post
			   ORDER BY create_time
			   DESC 
			   limit ?,?
				`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

func GetPostListByIDs(ids []string) (PostList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post 
	where post_id in (?)
	order by FIND_IN_SET(post_id, ?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&PostList, query, args...)
	return
}
