package ptfm_auth_svc_v1_entities

import "time"

type Identities struct {
	EmailIdentities []*EmailIdentity `json:"email_identities"`
	PhoneIdentities []*PhoneIdentity `json:"phone_identities"`
	SSOIdentities   []*SSOIdentity   `json:"sso_identities"`
}

type EmailIdentity struct {
	ID         string     `json:"id"`
	Email      string     `json:"email"`
	CreatedAt  time.Time  `json:"created_at"`
	VerifiedAt *time.Time `json:"verified_at,omitempty"`
}

type PhoneIdentity struct {
	ID         string     `json:"id"`
	Phone      string     `json:"phone"`
	CreatedAt  time.Time  `json:"created_at"`
	VerifiedAt *time.Time `json:"verified_at,omitempty"`
}

type SSOIdentity struct {
	ID         string    `json:"id"`
	Provider   string    `json:"provider"`
	ExternalID string    `json:"external_id"`
	CreatedAt  time.Time `json:"created_at"`
}
