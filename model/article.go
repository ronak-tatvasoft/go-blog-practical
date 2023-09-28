package model

import (
	"fmt"
	"time"
)

type Article struct {
	ID        int64     `db:"id" json:"id"`
	Nickname  string    `db:"nickname" json:"nickname"`
	Title     string    `db:"title" json:"title"`
	Content   string    `db:"content" json:"content"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// to get all articles list
func ArticlesList(offset int, limit int) ([]Article, error) {

	articleQuery := fmt.Sprintf("SELECT * FROM articles limit %d, %d", offset, limit)

	rows, err := db.Query(articleQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := []Article{}

	for rows.Next() {
		var art Article

		err := rows.Scan(&art.ID, &art.Nickname, &art.Title, &art.Content, &art.CreatedAt)
		if err != nil {
			return nil, err
		}

		articles = append(articles, art)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

// to get details of given article
func ArticleDetails(id int) (Article, error) {

	articleQuery := fmt.Sprintf("SELECT * FROM articles WHERE id=%d limit 1", id)

	row := db.QueryRow(articleQuery)

	article := Article{}

	err := row.Scan(&article.ID, &article.Nickname, &article.Title, &article.Content, &article.CreatedAt)

	return article, err
}

// to insert row in articles table
func NewArticle(nickname string, title string, content string) (int64, error) {

	result, err := db.Exec("INSERT INTO articles (nickname, title, content) VALUES (?, ?, ?)", nickname, title, content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// to check given article is exist or not
func IsArticleExist(id int) (bool, error) {
	var exists bool
	row := db.QueryRow(fmt.Sprintf("SELECT EXISTS(SELECT * FROM articles WHERE id=%d limit 1)", id))
	if err := row.Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

// get comments of given article
func (a *Article) Comments() ([]Comment, error) {
	comments, err := CommentsList(a.ID)
	return comments, err
}
