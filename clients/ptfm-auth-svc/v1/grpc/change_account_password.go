package ptfm_auth_svc_grpc_v1

import (
	"context"

	"github.com/bufbuild/connect-go"
	pkg_grpc "github.com/goauthnz/pkg/grpc"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	ptfm_auth_svc_v1_entities "github.com/goauthnz/contracts/clients/ptfm-auth-svc/v1/entities"
	pb "github.com/goauthnz/contracts/generated/services/ptfm/auth/svc/v1"
)

func (c *PtfmAuthSvcClient) ChangeAccountPassword(ctx context.Context, id string, params *ptfm_auth_svc_v1_entities.ChangeAccountPassword) error {
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
