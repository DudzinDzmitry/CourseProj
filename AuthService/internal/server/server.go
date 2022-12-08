package server

import (
	"context"
	"github.com/DudzinDzmitry/CourseProj/AuthService/internal/service"
	"github.com/DudzinDzmitry/CourseProj/AuthService/internal/user"
	pb "github.com/DudzinDzmitry/CourseProj/AuthService/proto"
)

type Server struct {
	currentService *service.Service
	pb.UnimplementedCRUDServer
}

func NewServerConnect(server *service.Service) *Server {
	return &Server{currentService: server}
}

func (s *Server) CreateAccount(ctx context.Context, request *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	newAcc := new(user.AccountIfo)
	newAcc.ID = request.Id
	newAcc.UserName = request.UserName
	newAcc.Password = request.Password
	err := s.currentService.CreateAccount(ctx, newAcc)
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
	newAcc := new(user.AccountIfo)
	newAcc.ID = request.Id
	newAcc.UserName = request.UserName
	newAcc.Password = request.Password
	if err := s.currentService.UpdateAccount(ctx, newAcc); err != nil {
		return nil, err
	}
	return new(pb.UpdateAccountResponse), nil
}
