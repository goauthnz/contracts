package authnz_store_svc_grpc_v1

import (
	"context"
	"time"

	"github.com/bufbuild/connect-go"
	pkg_grpc "github.com/goauthnz/pkg/grpc"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	authnz_store_svc_v1_entities "github.com/goauthnz/contracts/clients/authnz-store-svc/v1/entities"
	pb "github.com/goauthnz/contracts/generated/services/authnz/store/svc/v1"
)

func (c *AuthnzStoreSvcClient) GetAccountByID(ctx context.Context, id string) (*authnz_store_svc_v1_entities.Account, error) {
	req := connect.NewRequest(&pb.GetAccountByIDRequest{
		Id: wrapperspb.String(id),
	})

	res, err := c.client.GetAccountByID(ctx, req)
	if err != nil {
		return nil, pkg_grpc.TranslateFromGRPCError(ctx, err)
	}

	return &authnz_store_svc_v1_entities.Account{
		ID: res.Msg.GetAccount().GetId().GetValue(),
		Identities: func() *authnz_store_svc_v1_entities.Identities {
			returnedIdentities := &authnz_store_svc_v1_entities.Identities{
				EmailIdentities: make([]*authnz_store_svc_v1_entities.EmailIdentity, 0),
				PhoneIdentities: make([]*authnz_store_svc_v1_entities.PhoneIdentity, 0),
				SSOIdentities:   make([]*authnz_store_svc_v1_entities.SSOIdentity, 0),
			}

			if res.Msg.GetAccount().GetIdentities() == nil {
				return returnedIdentities
			}

			for _, identity := range res.Msg.GetAccount().GetIdentities().GetEmailIdentities() {
				returnedIdentities.EmailIdentities = append(returnedIdentities.EmailIdentities, &authnz_store_svc_v1_entities.EmailIdentity{
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
				returnedIdentities.PhoneIdentities = append(returnedIdentities.PhoneIdentities, &authnz_store_svc_v1_entities.PhoneIdentity{
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
				returnedIdentities.SSOIdentities = append(returnedIdentities.SSOIdentities, &authnz_store_svc_v1_entities.SSOIdentity{
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
