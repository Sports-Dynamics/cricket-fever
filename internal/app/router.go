package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/sports-dynamics/cricket-fever/internal/handlers"
	auth "github.com/sports-dynamics/cricket-fever/internal/middleware"
	"github.com/sports-dynamics/cricket-fever/internal/modo"
	"github.com/sports-dynamics/cricket-fever/internal/teams"
	"go.uber.org/zap"
)

func newRouter(config Config) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Get("/heartbeat", handlers.HeartbeatHandler)

	r.Route("/v1", func(r chi.Router) {

		r.Use(auth.AuthMiddleware(config.ValidAPIToken))

		db, err := NewPostgresDBConn()
		if err != nil {
			modo.Logger(context.Background()).Error("failed to connect with postres db", zap.Error(err))
			fmt.Println(" err : ", err)
		}

		teamRepo := teams.NewCreateTeamRepo(db)
		teamService := teams.NewTeamService(teamRepo)

		r.Route("/team", func(r chi.Router) {
			r.Post("/", teams.NewCreateTeamHandler(teamService))
		})
	})

	return r
}

func NewPostgresDBConn() (*sql.DB, error) {

	// Replace the connection string with your PostgreSQL database connection details
	connStr := "user=dev password=strongone dbname=postgres sslmode=require"

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the PostgreSQL database!")

	return db, nil
}
