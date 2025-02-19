package http

import (
	"md-tnt-mgmt/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUpRouter(controller *controller.Controller) http.Handler {
	var (
		router = mux.NewRouter()
	)
	router.HandleFunc("/tnt/api/v1/heartbeat", controller.HeartBeatHandler).Methods("GET")

	// Tenants API Routes
	router.HandleFunc("/tnt/api/v1/tenants", controller.CreateTenantHandler).Methods("POST")
	router.HandleFunc("/tnt/api/v1/tenants/{id}", controller.UpdateTenantHandler).Methods("PUT")
	router.HandleFunc("/tnt/api/v1/tenants/{id}/deactivate", controller.DeactivateTenantHandler).Methods("PATCH")
	router.HandleFunc("/tnt/api/v1/tenants/{id}", controller.DeleteTenantHandler).Methods("DELETE")
	router.HandleFunc("/tnt/api/v1/tenants", controller.GetTenantsHandler).Methods("GET")
	router.HandleFunc("/tnt/api/v1/tenants/{id}", controller.GetTenantByIDHandler).Methods("GET")

	return router
}
