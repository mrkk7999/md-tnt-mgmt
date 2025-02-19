package repository

import (
	"errors"
	"log"
	"md-tnt-mgmt/request_response/tenants"
	"time"

	"github.com/google/uuid"
)

// CreateTenant
func (r *repository) CreateTenant(tenant tenants.Tenants) (tenants.Tenants, error) {
	currentTime := time.Now().UTC()
	tenant.CreatedAt = currentTime
	tenant.UpdatedAt = currentTime

	if err := r.db.Create(&tenant).Error; err != nil {
		log.Println("Error creating tenant:", err)
		return tenant, err
	}
	return tenant, nil
}

// UpdateTenant
func (r *repository) UpdateTenant(tenant tenants.Tenants) (tenants.Tenants, error) {
	updateFields := map[string]interface{}{}

	if tenant.Name != "" {
		updateFields["name"] = tenant.Name
	}
	if tenant.ApprovalStatus != "" {
		updateFields["approval_status"] = tenant.ApprovalStatus
	}
	updateFields["updated_at"] = time.Now()

	if len(updateFields) > 0 {
		if err := r.db.Model(&tenant).Where("id = ?", tenant.ID).Updates(updateFields).Error; err != nil {
			log.Println("Error updating tenant:", err)
			return tenant, err
		}
	}
	// Use GetTenantByID to fetch the updated tenant
	updatedTenant, err := r.GetTenantByID(tenant.ID.String())
	if err != nil {
		log.Println("Error fetching updated tenant:", err)
		return tenant, err
	}

	return updatedTenant, nil
}

// DeactivateTenant
func (r *repository) DeactivateTenant(id string) (tenants.Tenants, error) {
	var tenant tenants.Tenants

	// string ID to UUID
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Println("Error parsing UUID:", err)
		return tenant, errors.New("Invalid UUID format")
	}

	if err := r.db.Model(&tenant).Where("id = ?", uuid).Updates(map[string]interface{}{
		"approval_status": "deactivated",
		"updated_at":      time.Now(),
	}).Error; err != nil {
		log.Println("Error deactivating tenant:", err)
		return tenant, err
	}

	// Fetch the updated tenant data after the update
	updatedTenant, err := r.GetTenantByID(id)
	if err != nil {
		log.Println("Error fetching updated tenant:", err)
		return tenant, err
	}

	return updatedTenant, nil
}

// DeleteTenant
func (r *repository) DeleteTenant(id string) error {
	// string ID to UUID
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Println("Error parsing UUID:", err)
		return errors.New("Invalid UUID format")
	}

	if err := r.db.Where("id = ?", uuid).Delete(&tenants.Tenants{}).Error; err != nil {
		log.Println("Error deleting tenant:", err)
		return err
	}

	return nil
}

// GetTenants
func (r *repository) GetTenants() ([]tenants.Tenants, error) {
	var tenantsList []tenants.Tenants
	if err := r.db.Find(&tenantsList).Error; err != nil {
		log.Println("Error getting tenants:", err)
		return nil, err
	}

	return tenantsList, nil
}

// GetTenantByID
func (r *repository) GetTenantByID(id string) (tenants.Tenants, error) {
	var tenant tenants.Tenants
	if err := r.db.Where("id = ?", id).First(&tenant).Error; err != nil {
		log.Println("Error getting tenant by ID:", err)
		return tenant, err
	}
	return tenant, nil
}

// GetTenantInfo
func (r *repository) GetTenantInfo(filters map[string]interface{}) (*tenants.Tenants, error) {
	var tenant tenants.Tenants
	query := r.db.Model(&tenant)

	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	if err := query.First(&tenant).Error; err != nil {
		log.Println("Error fetching tenant:", err)
		return nil, err
	}
	return &tenant, nil
}
