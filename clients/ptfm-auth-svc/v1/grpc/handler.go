package ptfm_auth_svc_grpc_v1

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/goauthnz/contracts/generated/services/ptfm/auth/svc/v1/svcv1connect"
	grpc_interceptors "github.com/goauthnz/pkg/grpc/interceptors"

	ptfm_auth_svc_v1 "github.com/goauthnz/contracts/clients/ptfm-auth-svc/v1"
)

type PtfmAuthSvcClient struct {
	client svcv1connect.PtfmAuthSvcClient
}

func NewPtfmAuthSvcClient(ctx context.Context, ptfmAuthSvcClientURL string) ptfm_auth_svc_v1.PtfmAuthSvc {
	iceptorsChain := grpc_interceptors.ClientDefaultInterceptors(grpc_interceptors.DefaultTimeout)
	httpClient := &http.Client{
		Timeout: grpc_interceptors.DefaultTimeout,
	}

	return &PtfmAuthSvcClient{
		client: svcv1connect.NewPtfmAuthSvcClient(httpClient, ptfmAuthSvcClientURL, connect.WithInterceptors(iceptorsChain...)),
	}
}
