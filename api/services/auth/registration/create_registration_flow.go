package registration

import (
	"context"
	"time"

	"connectrpc.com/connect"
	authv1 "github.com/anuragkumar19/connect/api/gen/auth/v1"
	"github.com/anuragkumar19/connect/database"
	"github.com/anuragkumar19/connect/token"
)

func (s *Service) CreateRegistrationFlow(ctx context.Context, r *connect.Request[authv1.CreateRegistrationFlowRequest]) (*connect.Response[authv1.CreateRegistrationFlowResponse], error) {
	now := time.Now()
	expiryTime := now.Add(flowLifeTime)

	csrfToken, err := token.Opaque()
	if err != nil {
		return nil, err
	}
	flow, err := s.store.CreateRegistrationFlow(ctx, &database.CreateRegistrationFlowParams{
		CreatedAt: now,
		UpdatedAt: now,
		ExpireAt:  expiryTime,
		CsrfToken: csrfToken,
		State:     database.RegistrationFlowsStateInvalidFields,
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&authv1.CreateRegistrationFlowResponse{
		Flow:  mapRegistrationFlow(&flow),
		Token: "todo", // TODO:
	})

	return res, nil
}
