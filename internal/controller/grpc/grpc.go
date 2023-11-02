package controller

import (
	"net"

	"github.com/math-88/medic-online/pkg/logger"
	pb "github.com/math-88/medic-online/proto/v1/gen/go"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// func unaryInterceptor(
// 	ctx context.Context,
// 	req interface{},
// 	info *grpc.UnaryServerInfo,
// 	handler grpc.UnaryHandler,
// ) (
// 	resp interface{},
// 	err error,
// ) {
// 	log.Println("--> unary interceptor", info.FullMethod)
// 	return handler(ctx, req)
// }

// func streamInterceptor(
// 	srv interface{},
// 	stream grpc.ServerStream,
// 	info *grpc.StreamServerInfo,
// 	handler grpc.StreamHandler,
// ) error {
// 	log.Println("--> stream interceptor", info.FullMethod)
// 	return handler(srv, stream)
// }

func accessibleRoles() map[string][]string {
	const medicalServicePath = "/v1.MedicalService/"

	return map[string][]string{
		medicalServicePath + "UploadImage": {"admin"},
		medicalServicePath + "CreateUser":  {"admin"},
		medicalServicePath + "GetUser":     {"admin", "user"},
	}
}

func RunGRPC(
	log *logger.Logger,
	authServer pb.AuthServiceServer,
	medicalServer pb.MedicalServiceServer,
	jwtManager *JWTManager,
	listener net.Listener,
) error {

	// controller := &Controller{
	// 	App: app,
	// }

	autInterceptor := NewAuthInterceptor(jwtManager, accessibleRoles())

	// create new gRPC server
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_prometheus.StreamServerInterceptor,
			// grpc_auth.StreamServerInterceptor(exampleAuthFunc),
			// streamInterceptor,
			autInterceptor.Stream(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			// grpc_auth.UnaryServerInterceptor(exampleAuthFunc),
			// unaryInterceptor,
			autInterceptor.Unary(),
		)),
	)

	// register the Controller on the gRPC server
	pb.RegisterAuthServiceServer(grpcServer, authServer)
	pb.RegisterMedicalServiceServer(grpcServer, medicalServer)

	grpc_prometheus.Register(grpcServer)

	// documentation
	reflection.Register(grpcServer)

	log.Info("Start GRPC server at %s", listener.Addr().String())

	return grpcServer.Serve(listener)
}
