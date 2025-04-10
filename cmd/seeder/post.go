package main

import (
	"fmt"
	"go-boilerplate/internal/api"
	"go-boilerplate/internal/config"
	"log"

	"github.com/ddosify/go-faker/faker"
)

func SeedPost() {
	f := faker.NewFaker()

	for i := 0; i < 10; i++ {
		post := api.Post{
			Title:    f.RandomJobTitle(),
			Content:  f.RandomLoremParagraphs(),
			AuthorID: 1, // Assuming a default author ID for seeding
		}
		stmt, err := config.Db.Db.Prepare("INSERT INTO post (title, content, author_id) VALUES (?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(post.Title, post.Content, post.AuthorID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Inserted post: %s, %s\n", post.Title, post.Content)
	}
}
