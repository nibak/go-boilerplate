package main

import (
	"context"
	"fmt"
	"go-boilerplate/internal/config"
)

func main() {
	// Initialize the database connection
	if err := config.InitDB(context.Background()); err != nil {
		fmt.Printf("Failed to initialize database: %v\n", err)
		return
	}
	defer config.Db.Db.Close()

	SeedUser()
	SeedPost()
}
