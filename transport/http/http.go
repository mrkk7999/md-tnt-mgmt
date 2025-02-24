package http

import (
	"md-tnt-mgmt/controller"
	"md-tnt-mgmt/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func SetUpRouter(controller *controller.Controller, log *logrus.Logger) http.Handler {
	var (
		router = mux.NewRouter()
	)
	// logging middleware
	router.Use(func(next http.Handler) http.Handler {
		return middleware.LoggingMiddleware(next, log)
	})

	router.HandleFunc("/tnt/api/v1/heartbeat", controller.HeartBeatHandler).Methods("GET")

	// Tenants API Routes
	router.HandleFunc("/tnt/api/v1/tenants/create", controller.CreateTenantHandler).Methods("POST")
	router.HandleFunc("/tnt/api/v1/tenants/update/{id}", controller.UpdateTenantHandler).Methods("PUT")
	router.HandleFunc("/tnt/api/v1/tenants/deactivate/{id}", controller.DeactivateTenantHandler).Methods("PATCH")
	router.HandleFunc("/tnt/api/v1/tenants/delete/{id}", controller.DeleteTenantHandler).Methods("DELETE")
	router.HandleFunc("/tnt/api/v1/tenants/all", controller.GetTenantsHandler).Methods("GET")
	router.HandleFunc("/tnt/api/v1/tenants/get/{id}", controller.GetTenantByIDHandler).Methods("GET")

	return router
}
