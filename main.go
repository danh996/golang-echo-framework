package main

import (
	"database/sql" // add this
	"fmt"
	"log"

	"github.com/danh996/go-fiber/book"
	"github.com/danh996/go-fiber/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq" // add this
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func initPostgresDB() {
	connStr := "postgresql://<admin>:<admin>@<localhost>/todos?sslmode=disable"
	// Connect to database
	_, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Postgres Database Success")

}

func main() {
	app := fiber.New()
	initDatabase()
	initPostgresDB()

	setupRoutes(app)
	app.Listen(3000)

	defer database.DBConn.Close()
}
