// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: auth/v1/registration_service.proto

package authv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/anuragkumar19/connect/api/gen/auth/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// RegistrationServiceName is the fully-qualified name of the RegistrationService service.
	RegistrationServiceName = "auth.v1.RegistrationService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RegistrationServiceRegisterProcedure is the fully-qualified name of the RegistrationService's
	// Register RPC.
	RegistrationServiceRegisterProcedure = "/auth.v1.RegistrationService/Register"
	// RegistrationServiceResendCodeProcedure is the fully-qualified name of the RegistrationService's
	// ResendCode RPC.
	RegistrationServiceResendCodeProcedure = "/auth.v1.RegistrationService/ResendCode"
	// RegistrationServiceVerifyProcedure is the fully-qualified name of the RegistrationService's
	// Verify RPC.
	RegistrationServiceVerifyProcedure = "/auth.v1.RegistrationService/Verify"
)

// RegistrationServiceClient is a client for the auth.v1.RegistrationService service.
type RegistrationServiceClient interface {
	// Register new user with email or phone number. Sends a verification code to email/phone number. It returns a token which can be used for further requests
	Register(context.Context, *connect.Request[v1.RegisterRequest]) (*connect.Response[v1.RegisterResponse], error)
	// Resend Code to email or phone number
	ResendCode(context.Context, *connect.Request[v1.ResendCodeRequest]) (*connect.Response[v1.ResendCodeResponse], error)
	// Verify emails or phone number with the code and token. If it return successful response you can login
	Verify(context.Context, *connect.Request[v1.VerifyRequest]) (*connect.Response[v1.VerifyResponse], error)
}

// NewRegistrationServiceClient constructs a client for the auth.v1.RegistrationService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRegistrationServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) RegistrationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	registrationServiceMethods := v1.File_auth_v1_registration_service_proto.Services().ByName("RegistrationService").Methods()
	return &registrationServiceClient{
		register: connect.NewClient[v1.RegisterRequest, v1.RegisterResponse](
			httpClient,
			baseURL+RegistrationServiceRegisterProcedure,
			connect.WithSchema(registrationServiceMethods.ByName("Register")),
			connect.WithIdempotency(connect.IdempotencyIdempotent),
			connect.WithClientOptions(opts...),
		),
		resendCode: connect.NewClient[v1.ResendCodeRequest, v1.ResendCodeResponse](
			httpClient,
			baseURL+RegistrationServiceResendCodeProcedure,
			connect.WithSchema(registrationServiceMethods.ByName("ResendCode")),
			connect.WithIdempotency(connect.IdempotencyIdempotent),
			connect.WithClientOptions(opts...),
		),
		verify: connect.NewClient[v1.VerifyRequest, v1.VerifyResponse](
			httpClient,
			baseURL+RegistrationServiceVerifyProcedure,
			connect.WithSchema(registrationServiceMethods.ByName("Verify")),
			connect.WithIdempotency(connect.IdempotencyIdempotent),
			connect.WithClientOptions(opts...),
		),
	}
}

// registrationServiceClient implements RegistrationServiceClient.
type registrationServiceClient struct {
	register   *connect.Client[v1.RegisterRequest, v1.RegisterResponse]
	resendCode *connect.Client[v1.ResendCodeRequest, v1.ResendCodeResponse]
	verify     *connect.Client[v1.VerifyRequest, v1.VerifyResponse]
}

// Register calls auth.v1.RegistrationService.Register.
func (c *registrationServiceClient) Register(ctx context.Context, req *connect.Request[v1.RegisterRequest]) (*connect.Response[v1.RegisterResponse], error) {
	return c.register.CallUnary(ctx, req)
}

// ResendCode calls auth.v1.RegistrationService.ResendCode.
func (c *registrationServiceClient) ResendCode(ctx context.Context, req *connect.Request[v1.ResendCodeRequest]) (*connect.Response[v1.ResendCodeResponse], error) {
	return c.resendCode.CallUnary(ctx, req)
}

// Verify calls auth.v1.RegistrationService.Verify.
func (c *registrationServiceClient) Verify(ctx context.Context, req *connect.Request[v1.VerifyRequest]) (*connect.Response[v1.VerifyResponse], error) {
	return c.verify.CallUnary(ctx, req)
}

// RegistrationServiceHandler is an implementation of the auth.v1.RegistrationService service.
type RegistrationServiceHandler interface {
	// Register new user with email or phone number. Sends a verification code to email/phone number. It returns a token which can be used for further requests
	Register(context.Context, *connect.Request[v1.RegisterRequest]) (*connect.Response[v1.RegisterResponse], error)
	// Resend Code to email or phone number
	ResendCode(context.Context, *connect.Request[v1.ResendCodeRequest]) (*connect.Response[v1.ResendCodeResponse], error)
	// Verify emails or phone number with the code and token. If it return successful response you can login
	Verify(context.Context, *connect.Request[v1.VerifyRequest]) (*connect.Response[v1.VerifyResponse], error)
}

// NewRegistrationServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRegistrationServiceHandler(svc RegistrationServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	registrationServiceMethods := v1.File_auth_v1_registration_service_proto.Services().ByName("RegistrationService").Methods()
	registrationServiceRegisterHandler := connect.NewUnaryHandler(
		RegistrationServiceRegisterProcedure,
		svc.Register,
		connect.WithSchema(registrationServiceMethods.ByName("Register")),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
		connect.WithHandlerOptions(opts...),
	)
	registrationServiceResendCodeHandler := connect.NewUnaryHandler(
		RegistrationServiceResendCodeProcedure,
		svc.ResendCode,
		connect.WithSchema(registrationServiceMethods.ByName("ResendCode")),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
		connect.WithHandlerOptions(opts...),
	)
	registrationServiceVerifyHandler := connect.NewUnaryHandler(
		RegistrationServiceVerifyProcedure,
		svc.Verify,
		connect.WithSchema(registrationServiceMethods.ByName("Verify")),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
		connect.WithHandlerOptions(opts...),
	)
	return "/auth.v1.RegistrationService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RegistrationServiceRegisterProcedure:
			registrationServiceRegisterHandler.ServeHTTP(w, r)
		case RegistrationServiceResendCodeProcedure:
			registrationServiceResendCodeHandler.ServeHTTP(w, r)
		case RegistrationServiceVerifyProcedure:
			registrationServiceVerifyHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRegistrationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRegistrationServiceHandler struct{}

func (UnimplementedRegistrationServiceHandler) Register(context.Context, *connect.Request[v1.RegisterRequest]) (*connect.Response[v1.RegisterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.RegistrationService.Register is not implemented"))
}

func (UnimplementedRegistrationServiceHandler) ResendCode(context.Context, *connect.Request[v1.ResendCodeRequest]) (*connect.Response[v1.ResendCodeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.RegistrationService.ResendCode is not implemented"))
}

func (UnimplementedRegistrationServiceHandler) Verify(context.Context, *connect.Request[v1.VerifyRequest]) (*connect.Response[v1.VerifyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.RegistrationService.Verify is not implemented"))
}
