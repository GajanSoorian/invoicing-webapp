package models

import "github.com/google/uuid"

// Model for Customer
type Customer struct {
	Id             uuid.UUID `json:"id"`                        // Unique identifier for the customer.
	OrganizationId uuid.UUID `json:"organization_id,omitempty"` // Unique identifier for the organization the customer works for.
	Email          string    `json:"type,omitempty"`            // Email address for the customer.
	Name           string    `json:"name,omitempty"`            // Customer name.
}

// Model for Organization entity
type Organization struct {
	OrganizationId     uuid.UUID    `json:"organization_id"`               // Unique identifier for an organization.
	RegisteredUsers    []uuid.UUID  `json:"registered_users,omitempty"`    // List of users associated this organization.
	RelationshipStatus clientStatus `json:"relationship_status,omitempty"` // Information about client based on purchasing history.
}

type clientStatus string

const (
	NEW       clientStatus = "NEW_CLIENT"
	REGULAR   clientStatus = "REGULAR_CLIENT"
	IRREGULAR clientStatus = "IRREGULAR_CLIENT"
)
