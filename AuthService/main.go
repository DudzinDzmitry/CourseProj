package AuthService

import (
	"context"
	"fmt"
	"github.com/DudzinDzmitry/CourseProj/AuthService/internal/repository"
	"github.com/DudzinDzmitry/CourseProj/AuthService/internal/server"
	"github.com/DudzinDzmitry/CourseProj/AuthService/internal/service"
	"github.com/DudzinDzmitry/CourseProj/AuthService/internal/user"
	pb "github.com/DudzinDzmitry/CourseProj/AuthService/proto"
	"github.com/caarlos0/env/v6"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		defer log.Fatalf("error while listening port: %e", err)
	}
	fmt.Println("Server successfully started on port :50051...")
	key := []byte("super-key")
	cfg := user.Config{JwtKey: key}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("failed to start service, %e", err)
	}
	conn := DBConnection(&cfg)
	fmt.Println("DB successfully connected...")

	ns := grpc.NewServer()
	newService := service.NewServiceConnect(conn, cfg.JwtKey)
	srv := server.NewServerConnect(newService)
	pb.RegisterCRUDServer(ns, srv)

	if err = ns.Serve(listen); err != nil {
		defer log.Fatalf("error while listening server: %e", err)
	}

}

// DBConnection create connection with db
func DBConnection(cfg *user.Config) repository.Repo {

	log.Info(cfg.PostgresDBURL)
	poolP, err := pgxpool.Connect(context.Background(), cfg.PostgresDBURL) // "postgresql://postgres:123@localhost:5432/person"
	if err != nil {
		log.Fatalf("bad connection with postgresql: %v", err)
		return nil
	}
	return &repository.PRepository{Pool: poolP}

}
