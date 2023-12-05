package app

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/sports-dynamics/cricket-fever/internal/handlers"
	auth "github.com/sports-dynamics/cricket-fever/internal/middleware"
	"github.com/sports-dynamics/cricket-fever/internal/repositories"
	"github.com/sports-dynamics/cricket-fever/internal/services"
)

func newRouter(config Config) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Get("/heartbeat", handlers.HeartbeatHandler)

	r.Route("/v1", func(r chi.Router) {

		r.Use(auth.AuthMiddleware(config.ValidAPIToken))

		teamRepo := repositories.TeamRepo{}
		teamService := services.NewTeamService(teamRepo)

		r.Route("/team", func(r chi.Router) {
			r.Post("/", handlers.NewCreateTeamHandler(teamService))
		})

		// r.Route("/accounts/{accountUUID}/sub_accounts", func(r chi.Router) {
		// 	r.Post("/", handlers.NewCreateSubAccountHandler(sovService))
		// 	r.Get("/", handlers.NewListSubAccountsByAccountIDHandler(subAccountRepo))
		// 	r.Get("/{subAccountUUID}", handlers.NewGetSubAccountByIdHandler(subAccountRepo))
		// })

		// r.Route("/records", func(r chi.Router) {
		// 	r.Post("/", handlers.NewCreateRecordHandler(sovService))
		// 	r.Get("/{recordUUID}", handlers.NewGetRecordByIdHandler(recordRepo))
		// })

		// r.Route("/disbursements", func(r chi.Router) {
		// 	r.Post("/", handlers.NewCreateDisbursementHandler(disbursementService))
		// 	r.Get("/{disbursementUUID}", handlers.NewGetDisbursementByIdHandler(disbursementRepo))
		// 	r.Patch("/{disbursementUUID}", handlers.NewProcessDisbursementHandler(disbursementService))
		// })

	})

	return r
}
