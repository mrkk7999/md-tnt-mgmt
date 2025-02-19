package iface

import "md-tnt-mgmt/request_response/tenants"

type Repository interface {
	HeartBeat() map[string]string
	CreateTenant(tenant tenants.Tenants) (tenants.Tenants, error)
	UpdateTenant(tenant tenants.Tenants) (tenants.Tenants, error)
	DeactivateTenant(id string) (tenants.Tenants, error)
	DeleteTenant(id string) error
	GetTenants() ([]tenants.Tenants, error)
	GetTenantByID(id string) (tenants.Tenants, error)
	GetTenantInfo(filters map[string]interface{}) (*tenants.Tenants, error)
}
