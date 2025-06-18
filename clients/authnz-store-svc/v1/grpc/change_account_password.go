package authnz_store_svc_grpc_v1

import (
	"context"

	"github.com/bufbuild/connect-go"
	pkg_grpc "github.com/goauthnz/pkg/grpc"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	authnz_store_svc_v1_entities "github.com/goauthnz/contracts/clients/authnz-store-svc/v1/entities"
	pb "github.com/goauthnz/contracts/generated/services/authnz/store/svc/v1"
)

func (c *AuthnzStoreSvcClient) ChangeAccountPassword(ctx context.Context, id string, params *authnz_store_svc_v1_entities.ChangeAccountPassword) error {
	req := connect.NewRequest(&pb.ChangeAccountPasswordRequest{
		Id:          wrapperspb.String(id),
		OldPassword: wrapperspb.String(params.OldPassword),
		NewPassword: wrapperspb.String(params.NewPassword),
	})

	_, err := c.client.ChangeAccountPassword(ctx, req)
	if err != nil {
		return pkg_grpc.TranslateFromGRPCError(ctx, err)
	}

	return nil
}
