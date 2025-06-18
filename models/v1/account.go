package models_v1

import "time"

type Account struct {
	ID           string             `json:"id"`
	Identities   *AccountIdentities `json:"identities"`
	IsMFAEnabled bool               `json:"is_mfa_enabled"`
	IsVerified   bool               `json:"is_verified"`
	HasPassword  bool               `json:"has_password"`
	CreatedAt    time.Time          `json:"created_at"`
}

type AccountIdentities struct {
	EmailIdentities []*EmailIdentity `json:"email_identities"`
	PhoneIdentities []*PhoneIdentity `json:"phone_identities"`
	SSOIdentities   []*SSOIdentity   `json:"sso_identities"`
}

type EmailIdentity struct {
	ID         string     `json:"id"`
	Email      string     `json:"email"`
	CreatedAt  time.Time  `json:"created_at"`
	VerifiedAt *time.Time `json:"verified_at"`
}

type PhoneIdentity struct {
	ID         string     `json:"id"`
	Phone      string     `json:"phone"`
	CreatedAt  time.Time  `json:"created_at"`
	VerifiedAt *time.Time `json:"verified_at"`
}

type SSOIdentity struct {
	ID        string    `json:"id"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
}

type PublicAccount struct {
	ID         string    `json:"id"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
}
