package controller

import (
	"context"

	"github.com/math-88/medic-online/internal/entity"
	pb "github.com/math-88/medic-online/proto/v1/gen/go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserStore is an interface to store users
type UserStore interface {
	Save(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByName(ctx context.Context, username string) (*entity.User, error)
}

// AuthServer is the server for authentication
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	userStore  UserStore
	jwtManager *JWTManager
}

// NewAuthServer returns a new auth server
func NewAuthServer(userStore UserStore, jwtManager *JWTManager) pb.AuthServiceServer {
	return &AuthServer{userStore: userStore, jwtManager: jwtManager}
}

// Login is a unary RPC to login user
func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	user, err := server.userStore.GetByName(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.LoginResponse{AccessToken: token}

	return res, nil
}
