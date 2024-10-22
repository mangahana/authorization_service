package grpc

import (
	"authorization_service/internal/application"
	pb "authorization_service/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type grpcServer struct {
	pb.UnimplementedAuthorizationServer
	server  *grpc.Server
	useCase application.UseCase
}

func New(useCase application.UseCase) *grpcServer {
	return &grpcServer{
		useCase: useCase,
		server:  grpc.NewServer(),
	}
}

func (s *grpcServer) Run(socket string) {
	listener, err := net.Listen("tcp", socket)
	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterAuthorizationServer(s.server, s)
	s.server.Serve(listener)
}

func (s *grpcServer) GetUser(ctx context.Context, r *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.useCase.GetMe(ctx, r.AccessToken)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		ID:          int32(user.ID),
		Username:    user.Username,
		Photo:       user.Photo,
		IsBanned:    user.IsBanned,
		Permissions: user.Permissions,
	}, nil
}
