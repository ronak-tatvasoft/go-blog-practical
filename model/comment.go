package model

import (
	"database/sql"
	"fmt"
	"time"
)

type Comment struct {
	ID              int64         `db:"id" json:"id"`
	Nickname        string        `db:"nickname" json:"nickname"`
	Content         string        `db:"content" json:"content"`
	CreatedAt       time.Time     `db:"created_at" json:"created_at"`
	ArticleID       int64         `db:"content" json:"article_id"`
	ParentCommentID sql.NullInt64 `db:"content" json:"parent_comment_id"`
}

// to insert row in comments table
func NewComment(articleId int, nickname string, content string) (int64, error) {

	result, err := db.Exec("INSERT INTO comments (article_id, nickname, content) VALUES (?, ?, ?)", articleId, nickname, content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// to get list of all comments of given article
func CommentsList(articleId int64) ([]Comment, error) {

	commentQuery := fmt.Sprintf("SELECT * FROM comments WHERE article_id=%d", articleId)

	rows, err := db.Query(commentQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}

	for rows.Next() {
		var cmt Comment

		err := rows.Scan(&cmt.ID, &cmt.Nickname, &cmt.Content, &cmt.CreatedAt, &cmt.ArticleID, &cmt.ParentCommentID)
		if err != nil {
			return nil, err
		}

		comments = append(comments, cmt)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

// to get details of given comment
func CommentDetails(id int) (Comment, error) {

	commentQuery := fmt.Sprintf("SELECT * FROM comments WHERE id=%d limit 1", id)

	row := db.QueryRow(commentQuery)

	comment := Comment{}

	err := row.Scan(&comment.ID, &comment.Nickname, &comment.Content, &comment.CreatedAt, &comment.ArticleID, &comment.ParentCommentID)

	return comment, err
}

// to insert row in comments with parent comment
func (c *Comment) NewChildComment(articleId int, nickname string, content string) (int64, error) {

	result, err := db.Exec("INSERT INTO comments (article_id, nickname, content, parent_comment_id) VALUES (?, ?, ?, ?)", articleId, nickname, content, c.ID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
