package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/volatiletech/sqlboiler/v4/boil"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	// Connect to the database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var connStr string = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection is successful
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	boil.SetDB(db)
	DB = db
	fmt.Println("😁 Connected to database")
}
