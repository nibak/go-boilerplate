package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-boilerplate/internal/api"
	"go-boilerplate/internal/config"
	"net/http"
)

func main() {
	// Initialize the database connection
	if err := config.InitDB(context.Background()); err != nil {
		fmt.Printf("Failed to initialize database: %v\n", err)
		return
	}
	defer config.Db.Db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Define the GET endpoint for posts
	mux.HandleFunc("/api/posts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		posts, _ := api.GetPosts(w, r)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(posts); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	})

	// Start the HTTP server
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}
