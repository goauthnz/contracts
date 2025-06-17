package user_store_svc_grpc_v1

import (
	"context"

	"github.com/bufbuild/connect-go"
	pkg_grpc "github.com/goauthnz/pkg/grpc"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	user_store_svc_v1_entities "github.com/goauthnz/contracts/clients/user-store-svc/v1/entities"
	pb "github.com/goauthnz/contracts/generated/services/user/store/svc/v1"
)

func (c *UserStoreSvcClient) UpdateProfilePicture(ctx context.Context, id string, url string) (*user_store_svc_v1_entities.User, error) {
	req := connect.NewRequest(&pb.UpdateProfilePictureRequest{
		Id:  wrapperspb.String(id),
		Url: wrapperspb.String(url),
	})

	res, err := c.client.UpdateProfilePicture(ctx, req)
	if err != nil {
		return nil, pkg_grpc.TranslateFromGRPCError(ctx, err)
	}

	return &user_store_svc_v1_entities.User{
		ID:                res.Msg.GetUser().GetId().GetValue(),
		AccountID:         res.Msg.GetUser().GetAccountId().GetValue(),
		ProfilePictureURL: res.Msg.GetUser().GetProfilePictureUrl().GetValue(),
		CreatedAt:         res.Msg.GetUser().GetCreatedAt().AsTime(),
		UpdatedAt:         res.Msg.GetUser().GetUpdatedAt().AsTime(),
	}, nil
}
