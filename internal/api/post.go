package api

import (
	"go-boilerplate/internal/config"
	"net/http"
)

type Post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int    `json:"author_id"`
}

func GetPosts(w http.ResponseWriter, r *http.Request) ([]Post, error) {
	ctx := r.Context()

	rows, err := config.Db.Db.QueryContext(ctx, "SELECT id, title, content, author_id FROM posts")
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content); err != nil {
			http.Error(w, "Failed to scan post", http.StatusInternalServerError)
			return posts, err
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating over posts", http.StatusInternalServerError)
		return posts, err
	}
	return posts, nil

}
