package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/jszymanowski/alive/models"
)

func initDB() {
	db, err := gorm.Open(sqlite.Open("db/development.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	migrateErr := db.AutoMigrate(&models.Feature{})
	if migrateErr != nil {
		panic("failed to migrate database: " + migrateErr.Error())
	}
}

func main() {
	initDB()

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(":3000", r)
}
