package server

import (
	"context"
	"github.com/DudzinDzmitry/CourseProj/AuthService/internal/service"
	"github.com/DudzinDzmitry/CourseProj/AuthService/internal/user"
	pb "github.com/DudzinDzmitry/CourseProj/AuthService/proto"
)

type Server struct {
	currentService *service.Service
}

func NewServerConnect(server *service.Service) *Server {
	return &Server{currentService: server}
}

func (s *Server) CreateAccount(ctx context.Context, request *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	err := s.currentService.CreateAccount(ctx, new(user.AccountIfo))
	if err != nil {
		return nil, err
	}
	return new(pb.CreateAccountResponse), nil
}

func (s *Server) DeleteAccount(ctx context.Context, request *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	err := s.currentService.DeleteAccount(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return new(pb.DeleteAccountResponse), nil
}

func (s *Server) LogIn(ctx context.Context, request *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	err := s.currentService.CreateAccount(ctx, new(user.AccountIfo))
	if err != nil {
		return nil, err
	}
	return new(pb.CreateAccountResponse), nil
}

func (s *Server) LogOut(ctx context.Context, request *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	err := s.currentService.DeleteAccount(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return new(pb.DeleteAccountResponse), nil
}

func (s *Server) UpdateAccount(ctx context.Context, request *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error) {
	if err := s.currentService.UpdateAccount(ctx); err != nil {
		return nil, err
	}
	newInfo := &user.AccountIfo{
		UserName: request.UserName,
	}
	err := s.currentService.UpdateAccount(ctx, request.Id, newInfo)
	if err != nil {
		return nil, err
	}
	return new(pb.UpdateAccountResponse), nil
}
