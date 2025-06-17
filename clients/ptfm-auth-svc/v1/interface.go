package ptfm_auth_svc_v1

import (
	"context"

	ptfm_auth_svc_v1_entities "github.com/goauthnz/contracts/clients/ptfm-auth-svc/v1/entities"
)

type PtfmAuthSvc interface {
	CreateAccountWithEmail(ctx context.Context, email string) (*ptfm_auth_svc_v1_entities.Account, string, error)
	CreateAccountWithEmailPassword(ctx context.Context, email string, password string) (*ptfm_auth_svc_v1_entities.Account, string, error)
	CreateAccountWithPhone(ctx context.Context, phone string) (*ptfm_auth_svc_v1_entities.Account, string, error)
	CreateAccountWithPhonePassword(ctx context.Context, phone string, password string) (*ptfm_auth_svc_v1_entities.Account, string, error)

	SendIdentityVerificationCode(ctx context.Context, id string) error
	VerifyIdentity(ctx context.Context, id string, code string) error

	CreateAccountWithSSO(ctx context.Context, provider string, code string) (*ptfm_auth_svc_v1_entities.Account, error)

	GetAccountMFASetup(ctx context.Context, id string) (string, string, error)
	EnableAccountMFA(ctx context.Context, id string, code string) error
	DisableAccountMFA(ctx context.Context, id string, code string) error

	GetAccountByID(ctx context.Context, id string) (*ptfm_auth_svc_v1_entities.Account, error)
	GetAccountByIdentityID(ctx context.Context, identityID string) (*ptfm_auth_svc_v1_entities.Account, error)
	GetAccountByEmail(ctx context.Context, email string) (*ptfm_auth_svc_v1_entities.Account, error)
	GetAccountByPhone(ctx context.Context, phone string) (*ptfm_auth_svc_v1_entities.Account, error)

	GetVerifiedEmailIdentityByEmail(ctx context.Context, email string) (*ptfm_auth_svc_v1_entities.EmailIdentity, error)
	GetVerifiedPhoneIdentityByPhone(ctx context.Context, phone string) (*ptfm_auth_svc_v1_entities.PhoneIdentity, error)

	ChangeAccountPassword(ctx context.Context, id string, params *ptfm_auth_svc_v1_entities.ChangeAccountPassword) error
	SetAccountPassword(ctx context.Context, id string, password string) error

	CreateSessionWithEmailPassword(ctx context.Context, email string, password string) (*ptfm_auth_svc_v1_entities.Session, error)
	CreateSessionWithPhonePassword(ctx context.Context, phone string, password string) (*ptfm_auth_svc_v1_entities.Session, error)

	DeleteSessionFromID(ctx context.Context, id string) error
	GetSessionByID(ctx context.Context, id string) (*ptfm_auth_svc_v1_entities.Session, error)
	GetSessionByToken(ctx context.Context, token string) (*ptfm_auth_svc_v1_entities.Session, error)
}
