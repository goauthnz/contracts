package user_store_svc_v1_entities

import "time"

type User struct {
	ID                string    `json:"id"`
	AccountID         string    `json:"account_id"`
	ProfilePictureURL string    `json:"profile_picture_url"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type User_Create struct {
	AccountID string `json:"account_id"`
}
