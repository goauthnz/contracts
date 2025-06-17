package ptfm_auth_svc_grpc_v1

import (
	"context"

	"github.com/bufbuild/connect-go"
	pkg_grpc "github.com/goauthnz/pkg/grpc"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	ptfm_auth_svc_v1_entities "github.com/goauthnz/contracts/clients/ptfm-auth-svc/v1/entities"
	pb "github.com/goauthnz/contracts/generated/services/ptfm/auth/svc/v1"
)

func (c *PtfmAuthSvcClient) CreateSessionWithEmailPassword(ctx context.Context, email string, password string) (*ptfm_auth_svc_v1_entities.Session, error) {
	req := connect.NewRequest(&pb.CreateSessionWithEmailPasswordRequest{
		Email:    wrapperspb.String(email),
		Password: wrapperspb.String(password),
	})

	res, err := c.client.CreateSessionWithEmailPassword(ctx, req)
	if err != nil {
		return nil, pkg_grpc.TranslateFromGRPCError(ctx, err)
	}

	return &ptfm_auth_svc_v1_entities.Session{
		ID:         res.Msg.GetSession().GetId().GetValue(),
		AccountID:  res.Msg.GetSession().GetAccountId().GetValue(),
		IdentityID: res.Msg.GetSession().GetIdentityId().GetValue(),
		Token:      res.Msg.GetSession().GetToken().GetValue(),
		CreatedAt:  res.Msg.GetSession().GetCreatedAt().AsTime(),
		UpdatedAt:  res.Msg.GetSession().GetUpdatedAt().AsTime(),
	}, nil
}
