package authnz_store_svc_grpc_v1

import (
	"context"

	"github.com/bufbuild/connect-go"
	pkg_grpc "github.com/goauthnz/pkg/grpc"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/goauthnz/contracts/generated/services/authnz/store/svc/v1"
)

func (c *AuthnzStoreSvcClient) SendIdentityVerificationCode(ctx context.Context, identityID string) error {
	req := connect.NewRequest(&pb.SendIdentityVerificationCodeRequest{
		IdentityId: wrapperspb.String(identityID),
	})

	_, err := c.client.SendIdentityVerificationCode(ctx, req)
	if err != nil {
		return pkg_grpc.TranslateFromGRPCError(ctx, err)
	}

	return nil
}
