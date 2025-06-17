package user_store_svc_grpc_v1

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	grpc_interceptors "github.com/goauthnz/pkg/grpc/interceptors"

	user_store_svc_v1 "github.com/goauthnz/contracts/clients/user-store-svc/v1"
	"github.com/goauthnz/contracts/generated/services/user/store/svc/v1/svcv1connect"
)

type UserStoreSvcClient struct {
	client svcv1connect.UserStoreSvcClient
}

func NewUserStoreSvcClient(ctx context.Context, userStoreSvcClientURL string) user_store_svc_v1.UserStoreSvc {
	iceptorsChain := grpc_interceptors.ClientDefaultInterceptors(grpc_interceptors.DefaultTimeout)
	httpClient := &http.Client{
		Timeout: grpc_interceptors.DefaultTimeout,
	}

	return &UserStoreSvcClient{
		client: svcv1connect.NewUserStoreSvcClient(httpClient, userStoreSvcClientURL, connect.WithInterceptors(iceptorsChain...)),
	}
}
