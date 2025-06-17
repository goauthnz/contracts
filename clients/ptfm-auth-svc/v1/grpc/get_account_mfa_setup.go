package ptfm_auth_svc_grpc_v1

import (
	"context"

	"github.com/bufbuild/connect-go"
	pkg_grpc "github.com/goauthnz/pkg/grpc"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/goauthnz/contracts/generated/services/ptfm/auth/svc/v1"
)

func (c *PtfmAuthSvcClient) GetAccountMFASetup(ctx context.Context, accountID string) (string, string, error) {
	req := connect.NewRequest(&pb.GetAccountMFASetupRequest{
		Id: wrapperspb.String(accountID),
	})

	res, err := c.client.GetAccountMFASetup(ctx, req)
	if err != nil {
		return "", "", pkg_grpc.TranslateFromGRPCError(ctx, err)
	}

	return res.Msg.GetSecret().GetValue(), res.Msg.GetQrCodeUrl().GetValue(), nil
}
