package ptfm_auth_svc_v1_entities

import "time"

type Session struct {
	ID         string    `json:"id"`
	AccountID  string    `json:"account_id"`
	IdentityID string    `json:"identity_id"`
	Token      string    `json:"token"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
