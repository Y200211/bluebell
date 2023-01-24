package mysql

import "bluebell/models"

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post 
    		   (post_id, title, content, author_id, community_id)	
			   values (?, ?, ?, ?, ?)
			   `
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return

}
