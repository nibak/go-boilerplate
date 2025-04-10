package main

import (
	"fmt"
	"go-boilerplate/internal/config"
	"log"

	"github.com/ddosify/go-faker/faker"
)

type User struct {
	Username string `faker:"username"`
	Email    string `faker:"email"`
	Password string `faker:"password"`
}

func SeedUser() {
	f := faker.NewFaker()

	for i := 0; i < 10; i++ {
		user := User{
			Username: f.RandomUsername(),
			Email:    f.RandomEmail(),
			Password: f.RandomPassword(),
		}
		// Insert the user into the database
		// config.Db.Db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, user.Password)
		stmt, err := config.Db.Db.Prepare("INSERT INTO user (username, email, password_hash) VALUES (?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(user.Username, user.Email, user.Password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Inserted user: %s, %s, %s\n", user.Username, user.Email, user.Password)
	}

}
