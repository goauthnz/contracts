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

func (c *AuthnzStoreSvcClient) GetVerifiedPhoneIdentityByPhone(ctx context.Context, phone string) (*authnz_store_svc_v1_entities.PhoneIdentity, error) {
	req := connect.NewRequest(&pb.GetVerifiedPhoneIdentityByPhoneRequest{
		Phone: wrapperspb.String(phone),
	})

	res, err := c.client.GetVerifiedPhoneIdentityByPhone(ctx, req)
	if err != nil {
		return nil, pkg_grpc.TranslateFromGRPCError(ctx, err)
	}

	return &authnz_store_svc_v1_entities.PhoneIdentity{
		ID:        res.Msg.GetIdentity().GetId().GetValue(),
		Phone:     res.Msg.GetIdentity().GetPhone().GetValue(),
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
