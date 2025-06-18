package authnz_store_svc_grpc_v1

import (
	"context"

	"github.com/bufbuild/connect-go"
	pkg_grpc "github.com/goauthnz/pkg/grpc"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	authnz_store_svc_v1_entities "github.com/goauthnz/contracts/clients/authnz-store-svc/v1/entities"
	pb "github.com/goauthnz/contracts/generated/services/authnz/store/svc/v1"
)

func (c *AuthnzStoreSvcClient) CreateSessionWithPhonePassword(ctx context.Context, phone string, password string) (*authnz_store_svc_v1_entities.Session, error) {
	req := connect.NewRequest(&pb.CreateSessionWithPhonePasswordRequest{
		Phone:    wrapperspb.String(phone),
		Password: wrapperspb.String(password),
	})

	res, err := c.client.CreateSessionWithPhonePassword(ctx, req)
	if err != nil {
		return nil, pkg_grpc.TranslateFromGRPCError(ctx, err)
	}

	return &authnz_store_svc_v1_entities.Session{
		ID:         res.Msg.GetSession().GetId().GetValue(),
		AccountID:  res.Msg.GetSession().GetAccountId().GetValue(),
		IdentityID: res.Msg.GetSession().GetIdentityId().GetValue(),
		Token:      res.Msg.GetSession().GetToken().GetValue(),
		CreatedAt:  res.Msg.GetSession().GetCreatedAt().AsTime(),
		UpdatedAt:  res.Msg.GetSession().GetUpdatedAt().AsTime(),
	}, nil
}
