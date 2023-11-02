package controller

import (
	"context"

	"github.com/math-88/medic-online/internal/entity"
	"github.com/math-88/medic-online/internal/usecase"
	pb "github.com/math-88/medic-online/proto/v1/gen/go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MedicalServer struct {
	pb.UnimplementedMedicalServiceServer
	userStore       UserStore
	protocolUseCase *usecase.ProtocolUseCase
}

// NewMedicalServer returns a new medical server
func NewMedicalServer(userStore UserStore, protocolUseCase *usecase.ProtocolUseCase) pb.MedicalServiceServer {
	return &MedicalServer{
		userStore:       userStore,
		protocolUseCase: protocolUseCase,
	}
}

func (server *MedicalServer) UserGet(ctx context.Context, req *pb.UserGetRequest) (*pb.UserGetReply, error) {

	user, err := server.userStore.GetByID(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	reply := &pb.UserGetReply{
		User: &pb.User{
			Id:   user.ID.String(),
			Name: user.Username,
			Role: pb.Role_user,
		},
	}

	return reply, nil
}

func (server *MedicalServer) ProtocolGet(ctx context.Context, req *pb.ProtocolGetRequest) (*pb.ProtocolGetReply, error) {

	anamnesis := entity.Anamnesis{
		Description: req.Anamnesis.Description,
	}

	protocol, err := server.protocolUseCase.MakeProtocol(ctx, anamnesis)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot make protocol: %v", err)
	}

	replyActions := make([]*pb.Actions, 0, len(protocol.Actions))
	for _, a := range protocol.Actions {
		act := &pb.Actions{
			Id:          a.ID.String(),
			Drug:        a.Drug,
			Dosage:      a.Dosage,
			Description: a.Description,
		}
		replyActions = append(replyActions, act)
	}

	reply := &pb.ProtocolGetReply{
		Protocol: &pb.Protocol{
			Actions: replyActions,
		},
	}

	return reply, nil
}
