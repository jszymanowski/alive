package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/jszymanowski/alive/handlers"
	"github.com/jszymanowski/alive/models"
	"github.com/jszymanowski/alive/repositories"
)

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db/development.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	migrateErr := db.AutoMigrate(&models.Feature{})
	if migrateErr != nil {
		panic("failed to migrate database: " + migrateErr.Error())
	}

	return db
}

func main() {
	db := initDB()

	// Setup repositories and handlers
	userRepo := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	// Routes
	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetAll)
		r.Post("/", userHandler.Create)
		r.Get("/{id}", userHandler.GetByID)
		// Add more routes...
	})

	http.ListenAndServe(":3000", r)
}
