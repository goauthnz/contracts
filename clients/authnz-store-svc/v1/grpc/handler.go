package authnz_store_svc_grpc_v1

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/goauthnz/contracts/generated/services/authnz/store/svc/v1/svcv1connect"
	grpc_interceptors "github.com/goauthnz/pkg/grpc/interceptors"

	ptfm_auth_svc_v1 "github.com/goauthnz/contracts/clients/authnz-store-svc/v1"
)

type AuthnzStoreSvcClient struct {
	client svcv1connect.AuthnzStoreSvcClient
}

func NewAuthnzStoreSvcClient(ctx context.Context, AuthnzStoreSvcClientURL string) ptfm_auth_svc_v1.PtfmAuthSvc {
	iceptorsChain := grpc_interceptors.ClientDefaultInterceptors(grpc_interceptors.DefaultTimeout)
	httpClient := &http.Client{
		Timeout: grpc_interceptors.DefaultTimeout,
	}

	return &AuthnzStoreSvcClient{
		client: svcv1connect.NewAuthnzStoreSvcClient(httpClient, AuthnzStoreSvcClientURL, connect.WithInterceptors(iceptorsChain...)),
	}
}
