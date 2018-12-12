package backend

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// Server represents grpc infrastructure for setting up a server
type Server struct {
	Health  *health.Server
	Server  *grpc.Server
	options []grpc.ServerOption
}

// Handler to register grpc server
type Handler interface {
	Register(server *grpc.Server)
}

// NewGrpcServer creates a new server
func NewGrpcServer(tracer opentracing.Tracer) *Server {
	s := Server{}
	s.Health = health.NewServer()

	// Add interceptors for instrumentation
	s.options = append(s.options,grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads())))
	s.options = append(s.options,grpc.StreamInterceptor(otgrpc.OpenTracingStreamServerInterceptor(tracer, otgrpc.LogPayloads())))


	s.Server = grpc.NewServer(s.options...)
	grpc_health_v1.RegisterHealthServer(s.Server, s.Health)

	reflection.Register(s.Server)
	return &s
}

// Start runs a server on a specific port
func (s *Server) Start() error {
	port := "6565"
	addr := fmt.Sprintf(":%s", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error starting server :: ", err.Error())
		return err
	}

	fmt.Println("Initiating gRPC server")

	err = s.Server.Serve(lis)
	if err != nil {
		fmt.Println("Error creating listener :: ", err.Error())
	}

	fmt.Println("gRPC server started")
	return nil
}

// Add a handler to server
func (s *Server) Add(handler Handler) {
	handler.Register(s.Server)
}
