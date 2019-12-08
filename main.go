package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/omkz/golang-postgre-hex/database/postgre"
	"github.com/omkz/golang-postgre-hex/post"
	"log"
)

func main() {
	var postRepo post.PostRepository
	postRepo = postgre.NewPostgresPostRepository(postgresConnection("postgresql://omz@localhost/blog_golang?sslmode=disable"))
	postService := post.NewPostService(postRepo)
	postHandler := post.NewPostHandler(postService)
	postHandler.Get()
}

func postgresConnection(database string) *sql.DB {
	fmt.Println("Connecting to PostgreSQL DB")
	db, err := sql.Open("postgres", database)
	if err != nil {
		log.Fatalf("%s", err)
		panic(err)
	}
	return db
}

