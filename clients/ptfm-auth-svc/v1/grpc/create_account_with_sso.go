package ptfm_auth_svc_grpc_v1

import (
	"context"
	"time"

	"github.com/bufbuild/connect-go"
	pkg_grpc "github.com/goauthnz/pkg/grpc"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	ptfm_auth_svc_v1_entities "github.com/goauthnz/contracts/clients/ptfm-auth-svc/v1/entities"
	pb "github.com/goauthnz/contracts/generated/services/ptfm/auth/svc/v1"
)

func (c *PtfmAuthSvcClient) CreateAccountWithSSO(ctx context.Context, provider string, code string) (*ptfm_auth_svc_v1_entities.Account, error) {
	req := connect.NewRequest(&pb.CreateAccountWithSSORequest{
		Provider: wrapperspb.String(provider),
		Code:     wrapperspb.String(code),
	})

	res, err := c.client.CreateAccountWithSSO(ctx, req)
	if err != nil {
		return nil, pkg_grpc.TranslateFromGRPCError(ctx, err)
	}

	return &ptfm_auth_svc_v1_entities.Account{
		ID: res.Msg.GetAccount().GetId().GetValue(),
		Identities: func() *ptfm_auth_svc_v1_entities.Identities {
			returnedIdentities := &ptfm_auth_svc_v1_entities.Identities{
				EmailIdentities: make([]*ptfm_auth_svc_v1_entities.EmailIdentity, 0),
				PhoneIdentities: make([]*ptfm_auth_svc_v1_entities.PhoneIdentity, 0),
				SSOIdentities:   make([]*ptfm_auth_svc_v1_entities.SSOIdentity, 0),
			}

			if res.Msg.GetAccount().GetIdentities() == nil {
				return returnedIdentities
			}

			for _, identity := range res.Msg.GetAccount().GetIdentities().GetEmailIdentities() {
				returnedIdentities.EmailIdentities = append(returnedIdentities.EmailIdentities, &ptfm_auth_svc_v1_entities.EmailIdentity{
					ID:        identity.GetId().GetValue(),
					Email:     identity.GetEmail().GetValue(),
					CreatedAt: identity.GetCreatedAt().AsTime(),
					VerifiedAt: func() *time.Time {
						if identity.GetVerifiedAt() != nil {
							t := identity.GetVerifiedAt().AsTime()
							return &t
						}
						return nil
					}(),
				})
			}

			for _, identity := range res.Msg.GetAccount().GetIdentities().GetPhoneIdentities() {
				returnedIdentities.PhoneIdentities = append(returnedIdentities.PhoneIdentities, &ptfm_auth_svc_v1_entities.PhoneIdentity{
					ID:        identity.GetId().GetValue(),
					Phone:     identity.GetPhone().GetValue(),
					CreatedAt: identity.GetCreatedAt().AsTime(),
					VerifiedAt: func() *time.Time {
						if identity.GetVerifiedAt() != nil {
							t := identity.GetVerifiedAt().AsTime()
							return &t
						}
						return nil
					}(),
				})
			}

			for _, identity := range res.Msg.GetAccount().GetIdentities().GetSsoIdentities() {
				returnedIdentities.SSOIdentities = append(returnedIdentities.SSOIdentities, &ptfm_auth_svc_v1_entities.SSOIdentity{
					ID:         identity.GetId().GetValue(),
					ExternalID: identity.GetExternalId().GetValue(),
					Provider:   identity.GetProvider().GetValue(),
					CreatedAt:  identity.GetCreatedAt().AsTime(),
				})
			}

			return returnedIdentities
		}(),
		IsMFAEnabled: res.Msg.GetAccount().GetIsMfaEnabled().GetValue(),
		Password:     res.Msg.GetAccount().GetPassword().GetValue(),
		CreatedAt:    res.Msg.GetAccount().GetCreatedAt().AsTime(),
		UpdatedAt:    res.Msg.GetAccount().GetUpdatedAt().AsTime(),
	}, nil
}
