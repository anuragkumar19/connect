package registration

import (
	"context"

	"connectrpc.com/connect"
	authv1 "github.com/anuragkumar19/connect/api/gen/auth/v1"
)

func (s *Service) DeleteRegistrationFlow(ctx context.Context, r *connect.Request[authv1.DeleteRegistrationFlowRequest]) (*connect.Response[authv1.DeleteRegistrationFlowResponse], error) {
	return nil, connect.NewError(connect.CodeUnavailable, nil)
}
