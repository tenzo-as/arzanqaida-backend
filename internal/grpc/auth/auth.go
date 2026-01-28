package auth

import (
	"context"

	arzanqaidav1 "github.com/tenzo-as/arzanqaida-proto/gen/go/arzanqaida"
	"google.golang.org/grpc"
)

type serverAPI struct {
	arzanqaidav1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	arzanqaidav1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *arzanqaidav1.LoginRequest,
) (*arzanqaidav1.LoginResponse, error) {
	panic("implement me")
}