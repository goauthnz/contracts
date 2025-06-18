package authnz_store_svc_grpc_v1

import (
	"context"

	"github.com/bufbuild/connect-go"
	pkg_grpc "github.com/goauthnz/pkg/grpc"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/goauthnz/contracts/generated/services/authnz/store/svc/v1"
)

func (c *AuthnzStoreSvcClient) SetAccountPassword(ctx context.Context, id string, password string) error {
	req := connect.NewRequest(&pb.SetAccountPasswordRequest{
		Id:       wrapperspb.String(id),
		Password: wrapperspb.String(password),
	})

	_, err := c.client.SetAccountPassword(ctx, req)
	if err != nil {
		return pkg_grpc.TranslateFromGRPCError(ctx, err)
	}

	return nil
}
