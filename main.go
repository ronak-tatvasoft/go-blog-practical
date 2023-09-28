package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/ronak-tatvasoft/go-blog-api/handler"
	"github.com/ronak-tatvasoft/go-blog-api/model"
)

func initDBWithMigrations() error {

	// init database handler
	dbH, err := model.DB()
	if err != nil {
		return err
	}

	// init database migration
	driver, err := mysql.WithInstance(dbH, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Println("Failed to create migration instance:", err)
	} else {
		// Apply pending migrations
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			fmt.Println("Error applying migrations:", err)
			return err
		}
		fmt.Println("Migrations applied successfully!")
	}

	return err
}

func main() {
	//db init
	err := initDBWithMigrations()

	if err != nil {
		panic(err)
	}

	// route using mux package
	r := mux.NewRouter().StrictSlash(true)
	r.Use(commonMiddleware)
	r.HandleFunc("/articles", handler.ArtilceList).Methods("GET")
	r.HandleFunc("/articles", handler.CreateArticle).Methods("POST")
	r.HandleFunc("/articles/{id}", handler.ArticleDetails).Methods("GET")
	r.HandleFunc("/articles/{id}/comments", handler.CommentList).Methods("GET")
	r.HandleFunc("/articles/{id}/comments", handler.CreateComment).Methods("POST")
	r.HandleFunc("/articles/{id}/comments/{commnetId}/reply", handler.ReplyComment).Methods("POST")

	fmt.Printf("Starting server.... \n")
	if err := http.ListenAndServe(os.Getenv("HOST")+":"+os.Getenv("PORT"), r); err != nil {
		log.Fatal(err)
	}
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
