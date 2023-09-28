CREATE TABLE comments (
	id SERIAL PRIMARY KEY,
	nickname VARCHAR(255) NOT NULL, 
	content TEXT NOT NULL, 
	created_at TIMESTAMP DEFAULT NOW(),
	article_id INT(11) UNSIGNED NOT NULL, 
	parent_comment_id INT(11) UNSIGNED NULL DEFAULT NULL, 
	FOREIGN KEY (article_id) REFERENCES articles(id),
	FOREIGN KEY (parent_comment_id) REFERENCES comments(id)
);