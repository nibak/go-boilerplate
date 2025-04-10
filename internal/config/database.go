package config

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type DbConnection struct {
	Db *sql.DB
}

var Db *DbConnection

func InitDB(ctx context.Context) error {
	db := &DbConfig{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Username: os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Database: os.Getenv("MYSQL_DATABASE"),
	}
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true", db.Username, db.Password, db.Host, db.Port, db.Database)

	// Connect to the database using the DSN
	DbCon, err := sql.Open("mysql", DSN)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Check if the connection is successful
	if err = DbCon.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}
	fmt.Println("Database connection established successfully.")
	// Store the connection in the DbConnection struct
	Db = &DbConnection{Db: DbCon}
	// You can use dbConn.DB to perform database operations
	runDbMigration()
	return nil
}

func runDbMigration() error {
	driver, err := mysql.WithInstance(Db.Db, &mysql.Config{})
	if err != nil {
		fmt.Printf("Failed to create MySQL driver: %v\n", err)
		return err
	}
	// dir, err := os.Getwd()
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("Current working directory:", dir)
	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/migration",
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Printf("Failed to create migration instance: %v\n", err)
		return err
	}
	// Run the migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Printf("Failed to run migrations: %v\n", err)
		return err
	}
	fmt.Println("Database migrations applied successfully.")
	return nil
}
