package endpoint

import (
	"context"

	"github.com/Job-Finder-Network/api-gateway/reqresp"
	"github.com/Job-Finder-Network/api-gateway/service"
	"github.com/go-kit/kit/endpoint"
)

type UserEndpoints struct {
	CreateUser endpoint.Endpoint
	Login      endpoint.Endpoint
}

func MakeEndpoints(userService service.UserService, loginService service.LoginServiceInterface) UserEndpoints {
	return UserEndpoints{
		CreateUser: makeCreateUserEndpoint(userService),
		Login:      makeLoginEndpoint(userService, loginService),
	}
}

func makeCreateUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(reqresp.CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password, req.Role)
		return reqresp.CreateUserResponse{Ok: ok}, err
	}
}
func makeLoginEndpoint(userService service.UserService, loginService service.LoginServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(reqresp.LoginRequest)
		ok, err := userService.AuthenticateUser(ctx, req.Email, req.Password)
		if !ok {
			return reqresp.LoginResponse{Token: ""}, err
		}
		user, err := userService.FindUserByEmail(ctx, req.Email)
		token, err := loginService.GenerateToken(req.Email, user.Role)
		return reqresp.LoginResponse{Token: token}, err
	}
}
