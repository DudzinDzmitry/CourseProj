package server

import (
	"context"
	"github.com/DudzinDzmitry/CourseProj/AuthService/internal/service"
	"github.com/DudzinDzmitry/CourseProj/AuthService/internal/user"
	pb "github.com/DudzinDzmitry/CourseProj/AuthService/proto"
)

type Server struct {
	pb.UnimplementedCRUDServer
	currentService *service.Service
}

// NewServerConnect create new server connection
func NewServerConnect(server *service.Service) *Server {
	return &Server{currentService: server}
}

// DeleteUser delete user by id
func (s *Server) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.Response, error) {
	idUser := request.GetId()
	err := s.currentService.DeleteUser(ctx, idUser)
	if err != nil {
		return nil, err
	}
	return new(pb.Response), nil
}

// UpdateUser update user with new parameters
func (s *Server) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.Response, error) {
	accessToken := request.GetAccessToken()
	if err := s.currentService.Verify(accessToken); err != nil {
		return nil, err
	}
	user := &user.User{
		Name:     request.Person.Name,
		Position: request.Person.Position,
	}
	idUser := request.GetId()
	err := s.currentService.UpdateUser(ctx, idUser, user)
	if err != nil {
		return nil, err
	}
	return new(pb.Response), nil
}
