package user_store_svc_v1

import (
	"context"

	user_store_svc_v1_entities "github.com/goauthnz/contracts/clients/user-store-svc/v1/entities"
)

type UserStoreSvc interface {
	CreateUser(ctx context.Context, params *user_store_svc_v1_entities.User_Create) (*user_store_svc_v1_entities.User, error)
	GetUserByID(ctx context.Context, id string) (*user_store_svc_v1_entities.User, error)
	GetUserByAccountID(ctx context.Context, accountID string) (*user_store_svc_v1_entities.User, error)

	UpdateProfilePicture(ctx context.Context, id string, url string) (*user_store_svc_v1_entities.User, error)
	DeleteProfilePicture(ctx context.Context, id string) (*user_store_svc_v1_entities.User, error)
}
