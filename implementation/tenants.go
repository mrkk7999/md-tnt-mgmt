package implementation

import (
	"errors"
	"md-tnt-mgmt/request_response/tenants"
)

// CreateTenant
func (s *service) CreateTenant(tenant tenants.Tenants) (tenants.Tenants, error) {
	existingTenant, err := s.repo.GetTenantInfo(map[string]interface{}{
		"govt_register_id": tenant.GovtRegisterID,
	})
	if err == nil && existingTenant != nil {
		return tenants.Tenants{}, errors.New("tenant already exists")
	}

	return s.repo.CreateTenant(tenant)
}

// UpdateTenant
func (s *service) UpdateTenant(tenant tenants.Tenants) (tenants.Tenants, error) {
	return s.repo.UpdateTenant(tenant)
}

// DeactivateTenant
func (s *service) DeactivateTenant(id string) (tenants.Tenants, error) {
	return s.repo.DeactivateTenant(id)
}

// DeleteTenant
func (s *service) DeleteTenant(id string) error {
	return s.repo.DeleteTenant(id)
}

// GetTenants
func (s *service) GetTenants() ([]tenants.Tenants, error) {
	return s.repo.GetTenants()
}

// GetTenantByID
func (s *service) GetTenantByID(id string) (tenants.Tenants, error) {
	return s.repo.GetTenantByID(id)
}
