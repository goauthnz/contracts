package authnz_store_svc_v1_entities

import "time"

type Account struct {
	ID           string      `json:"id"`
	Identities   *Identities `json:"identities"`
	IsMFAEnabled bool        `json:"is_mfa_enabled"`
	Password     string      `json:"password"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

type ChangeAccountPassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
