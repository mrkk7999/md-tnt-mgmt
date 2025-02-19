package tenants

import "time"

// TenantReq
type TenantReq struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	ApprovalStatus string    `json:"approval_status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at,omitempty"`
}
