package registration

import (
	"context"

	"connectrpc.com/connect"
	authv1 "github.com/anuragkumar19/connect/api/gen/auth/v1"
)

func (s *Service) UpdateRegistrationFlow(ctx context.Context, r *connect.Request[authv1.UpdateRegistrationFlowRequest]) (*connect.Response[authv1.UpdateRegistrationFlowResponse], error) {
	return nil, connect.NewError(connect.CodeUnavailable, nil)
}
