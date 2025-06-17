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

func (c *PtfmAuthSvcClient) GetVerifiedEmailIdentityByEmail(ctx context.Context, email string) (*ptfm_auth_svc_v1_entities.EmailIdentity, error) {
	req := connect.NewRequest(&pb.GetVerifiedEmailIdentityByEmailRequest{
		Email: wrapperspb.String(email),
	})

	res, err := c.client.GetVerifiedEmailIdentityByEmail(ctx, req)
	if err != nil {
		return nil, pkg_grpc.TranslateFromGRPCError(ctx, err)
	}

	return &ptfm_auth_svc_v1_entities.EmailIdentity{
		ID:        res.Msg.GetIdentity().GetId().GetValue(),
		Email:     res.Msg.GetIdentity().GetEmail().GetValue(),
		CreatedAt: res.Msg.GetIdentity().GetCreatedAt().AsTime(),
		VerifiedAt: func() *time.Time {
			if res.Msg.GetIdentity().GetVerifiedAt() != nil {
				t := res.Msg.GetIdentity().GetVerifiedAt().AsTime()
				return &t
			}
			return nil
		}(),
	}, nil
}
