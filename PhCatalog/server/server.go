package server

import (
	"context"
	"github.com/DudzinDzmitry/CourseProj/PhCatalog/catalog"
	pb "github.com/DudzinDzmitry/CourseProj/PhCatalog/proto"
	"github.com/DudzinDzmitry/CourseProj/PhCatalog/service"
	__ "github.com/DudzinDzmitry/CourseProj/UserService/proto"
)

// Server struct
type Server struct {
	pb.UnimplementedPharmacyCatalogServer
	se *service.Service
}

func (s *Server) Registration(ctx context.Context, request *__.RegistrationRequest) (*__.RegistrationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetUser(ctx context.Context, request *__.GetUserRequest) (*__.GetUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetAllUsers(ctx context.Context, request *__.GetAllUsersRequest) (*__.GetAllUsersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeleteUser(ctx context.Context, request *__.DeleteUserRequest) (*__.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UpdateUser(ctx context.Context, request *__.UpdateUserRequest) (*__.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Authentication(ctx context.Context, request *__.AuthenticationRequest) (*__.AuthenticationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) RefreshMyTokens(ctx context.Context, request *__.RefreshTokensRequest) (*__.RefreshTokensResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Logout(ctx context.Context, request *__.LogoutRequest) (*__.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) mustEmbedUnimplementedCRUDServer() {
	//TODO implement me
	panic("implement me")
}

// NewServer create new server connection
func NewServer(serv *service.Service) *Server {
	return &Server{se: serv}
}

// CreateMedicine create new medicine
func (s *Server) CreateMedicine(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	m := catalog.Medicine{
		Name:  request.Name,
		Count: request.Count,
		Price: request.Price,
	}
	newID, err := s.se.CreateMedicine(ctx, &m)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{Id: newID}, nil
}

// GetMedicine get medicine by id from db
func (s *Server) GetMedicine(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	idMedicine := request.GetId()
	medicineDB, err := s.se.GetMedicine(ctx, idMedicine)
	if err != nil {
		return nil, err
	}
	medicineProto := &pb.GetResponse{
		Med: &pb.Medicine{
			Id:    medicineDB.Id,
			Name:  medicineDB.Name,
			Count: medicineDB.Count,
			Price: medicineDB.Price,
		},
	}
	return medicineProto, nil
}

// GetAllMedicine get all medicine from db
func (s *Server) GetAllMedicine(ctx context.Context, _ *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	medicines, err := s.se.GetAllMedicine(ctx)
	if err != nil {
		return nil, err
	}
	var list []*pb.Medicine
	for _, medicine := range medicines {
		medicineProto := new(pb.Medicine)
		medicineProto.Id = medicine.Id
		medicineProto.Name = medicine.Name
		medicineProto.Count = medicine.Count
		medicineProto.Price = medicine.Price
		list = append(list, medicineProto)
	}
	return &pb.GetAllResponse{Med: list}, nil
}

// DeleteMedicine delete medicine by id
func (s *Server) DeleteMedicine(ctx context.Context, request *pb.DelRequest) (*pb.Response, error) {
	idMed := request.GetId()
	err := s.se.DeleteMedicine(ctx, idMed)
	if err != nil {
		return nil, err
	}
	return new(pb.Response), nil
}

// ChangeMedicine update medicine with new parameters
func (s *Server) ChangeMedicine(ctx context.Context, request *pb.ChRequest) (*pb.Response, error) {
	med := &catalog.Medicine{
		Name:  request.Med.Name,
		Count: request.Med.Count,
		Price: request.Med.Price,
	}
	idMed := request.GetId()
	err := s.se.ChangeMedicine(ctx, idMed, med)
	if err != nil {
		return nil, err
	}
	return new(pb.Response), nil
}
