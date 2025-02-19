package controller

import (
	"encoding/json"
	"errors"
	"md-tnt-mgmt/request_response/tenants"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// CreateTenantHandler
func (c *Controller) CreateTenantHandler(w http.ResponseWriter, r *http.Request) {
	var tenant tenants.Tenants
	if err := json.NewDecoder(r.Body).Decode(&tenant); err != nil {
		encodeJSONResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	createdTenant, err := c.svc.CreateTenant(tenant)
	if err != nil {
		encodeJSONResponse(w, http.StatusConflict, nil, err) // Conflict status for duplicate tenant
		return
	}
	encodeJSONResponse(w, http.StatusOK, createdTenant, nil)
}

// UpdateTenantHandler
func (c *Controller) UpdateTenantHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		encodeJSONResponse(w, http.StatusBadRequest, nil, errors.New("Empty id"))
		return
	}
	var tenant tenants.Tenants
	if err := json.NewDecoder(r.Body).Decode(&tenant); err != nil {
		encodeJSONResponse(w, http.StatusBadRequest, nil, err)
		return
	}
	tenantID, err := uuid.Parse(id)
	if err != nil {
		encodeJSONResponse(w, http.StatusBadRequest, nil, errors.New("Invalid UUID"))
		return
	}
	tenant.ID = tenantID

	if tenant.GovtRegisterID != "" {
		encodeJSONResponse(w, http.StatusBadRequest, nil, errors.New("Not allowed to update GovtRegisterID"))
		return
	}
	updatedTenant, err := c.svc.UpdateTenant(tenant)
	if err != nil {
		encodeJSONResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	encodeJSONResponse(w, http.StatusOK, updatedTenant, nil)
}

// DeactivateTenantHandler
func (c *Controller) DeactivateTenantHandler(w http.ResponseWriter, r *http.Request) {
	// Extract path parameter using mux.Vars
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		encodeJSONResponse(w, http.StatusBadRequest, nil, errors.New("Empty id"))
		return
	}
	deactivatedTenant, err := c.svc.DeactivateTenant(id)
	if err != nil {
		encodeJSONResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	encodeJSONResponse(w, http.StatusOK, deactivatedTenant, nil)
}

// DeleteTenantHandler
func (c *Controller) DeleteTenantHandler(w http.ResponseWriter, r *http.Request) {
	// Extract path parameter using mux.Vars
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		encodeJSONResponse(w, http.StatusBadRequest, nil, errors.New("Empty id"))
		return
	}
	_, err := c.svc.GetTenantByID(id)
	if err != nil {
		encodeJSONResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "User record not found",
		}, err)
		return
	}
	err = c.svc.DeleteTenant(id)
	if err != nil {
		encodeJSONResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	successMessage := map[string]string{
		"message": "Tenant successfully deleted",
	}
	encodeJSONResponse(w, http.StatusOK, successMessage, nil)
}

// GetTenantsHandler
func (c *Controller) GetTenantsHandler(w http.ResponseWriter, r *http.Request) {
	tenantsList, err := c.svc.GetTenants()
	if err != nil {
		encodeJSONResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	encodeJSONResponse(w, http.StatusOK, tenantsList, nil)
}

// GetTenantByIDHandler
func (c *Controller) GetTenantByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Extract path parameter using mux.Vars
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		encodeJSONResponse(w, http.StatusBadRequest, nil, errors.New("Empty id"))
		return
	}
	tenant, err := c.svc.GetTenantByID(id)
	if err != nil {
		encodeJSONResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	encodeJSONResponse(w, http.StatusOK, tenant, nil)
}
