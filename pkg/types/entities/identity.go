package entities

import (
	"time"
)

type Identity struct {
	Alias               string
	Address             string
	PublicKey           string
	CompressedPublicKey string
	Active              bool
	TenantID            string
	Attributes          map[string]string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}