package tenants

import (
	"time"

	"github.com/google/uuid"
)

// Enum for ApprovalStatus
type ApprovalStatus string

const (
	pending     ApprovalStatus = "pending"
	approved    ApprovalStatus = "approved"
	rejected    ApprovalStatus = "rejected"
	suspended   ApprovalStatus = "suspended"
	active      ApprovalStatus = "active"
	deactivated ApprovalStatus = "deactivated"
)

// Struct for Tenants
type Tenants struct {
	ID             uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name           string         `json:"name" gorm:"type:varchar(255);not null"`
	GovtRegisterID string         `json:"govt_register_id" gorm:"type:varchar(100);unique;not null"`
	ApprovalStatus ApprovalStatus `json:"approval_status" gorm:"type:varchar(50);default:'pending'"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      *time.Time     `json:"deleted_at,omitempty" gorm:"index"`
}
