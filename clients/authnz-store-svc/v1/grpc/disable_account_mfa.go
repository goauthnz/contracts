package authnz_store_svc_grpc_v1

import (
	"context"

	"github.com/bufbuild/connect-go"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/goauthnz/contracts/generated/services/authnz/store/svc/v1"
	pkg_grpc "github.com/goauthnz/pkg/grpc"
)

func (c *AuthnzStoreSvcClient) DisableAccountMFA(ctx context.Context, accountID string, code string) error {
	req := connect.NewRequest(&pb.DisableAccountMFARequest{
		Id:               wrapperspb.String(accountID),
		VerificationCode: wrapperspb.String(code),
	})

	_, err := c.client.DisableAccountMFA(ctx, req)
	if err != nil {
		return pkg_grpc.TranslateFromGRPCError(ctx, err)
	}

	return nil
}
