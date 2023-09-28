package test

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "os"
    "testing"

    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/mysql"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    "github.com/gorilla/mux"
    "github.com/ronak-tatvasoft/go-blog-api/handler"
    "github.com/ronak-tatvasoft/go-blog-api/model"
)

func TestMain(m *testing.M) {
    upMigrations()
    code := m.Run()
    downMigrations()
    routerT()
    os.Exit(code)
}

func upMigrations() error {

    // init database handler
    dbH, err := model.DB()
    if err != nil {
        return err
    }

    // init database migration
    driver, err := mysql.WithInstance(dbH, &mysql.Config{})
    m, err := migrate.NewWithDatabaseInstance(
        "file://./../migrations",
        "mysql",
        driver,
    )
    if err != nil {
        fmt.Println("Failed to create migration instance:", err)
    } else {
        // Apply pending migrations
        if err := m.Up(); err != nil && err != migrate.ErrNoChange {
            fmt.Println("Error applying migrations up:", err)
            return err
        }
    }

    return err
}

func downMigrations() error {

    // init database handler
    dbH, err := model.DB()
    if err != nil {
        return err
    }

    // init database migration
    driver, err := mysql.WithInstance(dbH, &mysql.Config{})
    m, err := migrate.NewWithDatabaseInstance(
        "file://./../migrations",
        "mysql",
        driver,
    )
    if err != nil {
        fmt.Println("Failed to create migration instance:", err)
    } else {
        // Apply pending migrations
        if err := m.Down(); err != nil && err != migrate.ErrNoChange {
            fmt.Println("Error applying migrations down:", err)
            return err
        }
    }

    return err
}

func routerT() *mux.Router {
    r := mux.NewRouter().StrictSlash(true)
    r.Use(commonMiddleware)
    r.HandleFunc("/articles", handler.ArtilceList).Methods("GET")
    r.HandleFunc("/articles", handler.CreateArticle).Methods("POST")
    r.HandleFunc("/articles/{id}", handler.ArticleDetails).Methods("GET")
    r.HandleFunc("/articles/{id}/comments", handler.CommentList).Methods("GET")
    r.HandleFunc("/articles/{id}/comments", handler.CreateComment).Methods("POST")
    r.HandleFunc("/articles/{id}/comments/{commnetId}/reply", handler.ReplyComment).Methods("POST")
    return r
}

func commonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    routerT().ServeHTTP(rr, req)

    return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}
