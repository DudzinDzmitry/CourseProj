package server

import (
	"context"
	pb "github.com/AHacTacIA/KP/PhCatalog/proto"
	"github.com/DudzinDzmitry/CourseProj/OrderService/internal/order"
	"github.com/DudzinDzmitry/CourseProj/OrderService/internal/service"
)

// Server struct
type Server struct {
	pb.UnimplementedPharmacyCatalogServer
	se *service.Service
}

// NewServer create new server connection
func NewServer(serv *service.Service) *Server {
	return &Server{se: serv}
}

// CreateMedicine create new medicine
func (s *Server) CreateMedicine(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	m := order.Item{
		Name:  request.Name,
		Count: request.Count,
		Price: request.Price,
	}
	newID, err := s.se.CreateItem(ctx, &m)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{Id: newID}, nil
}

// GetMedicine get medicine by id from db
func (s *Server) GetMedicine(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	idMedicine := request.GetId()
	medicineDB, err := s.se.GetItem(ctx, idMedicine)
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
	medicines, err := s.se.GetAllItems(ctx)
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
	err := s.se.DeleteItem(ctx, idMed)
	if err != nil {
		return nil, err
	}
	return new(pb.Response), nil
}

// ChangeMedicine update medicine with new parameters
func (s *Server) ChangeMedicine(ctx context.Context, request *pb.ChRequest) (*pb.Response, error) {
	med := &order.Item{
		Name:  request.Med.Name,
		Count: request.Med.Count,
		Price: request.Med.Price,
	}
	idMed := request.GetId()
	err := s.se.UpdateItem(ctx, idMed, med)
	if err != nil {
		return nil, err
	}
	return new(pb.Response), nil
}
