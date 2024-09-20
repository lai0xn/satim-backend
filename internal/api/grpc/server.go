package grpc

import (
	"net"

	"github.com/charmbracelet/log"
	"github.com/laix0n/satim/internal/api/grpc/handler"
	"github.com/laix0n/satim/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	ADDR string
}

func (s *Server) Run() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterTestServiceServer(server, handler.Handler{})
	reflection.Register(server)
	log.Info("Server running on port 8080")
	if err := server.Serve(l); err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	return &Server{}
}
