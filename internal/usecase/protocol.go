package usecase

import (
	"context"

	"github.com/math-88/medic-online/internal/entity"
)

type Repo interface {
	GetProtocol(ctx context.Context, diagnosis string) (*entity.Protocol, error)
	GetProtocolActions(ctx context.Context, protocolID []byte) ([]entity.Action, error)
}

type WebAPI interface {
	Ask(context.Context, string) (string, error)
}

// ProtocolUseCase - протокол лечения
type ProtocolUseCase struct {
	repo   Repo
	webAPI WebAPI
}

// New -.
func New(
	repo Repo,
	webAPI WebAPI,
) *ProtocolUseCase {
	return &ProtocolUseCase{
		repo:   repo,
		webAPI: webAPI,
	}
}

// Make protocol -.
func (pc *ProtocolUseCase) MakeProtocol(ctx context.Context, anamnesis entity.Anamnesis) (*entity.Protocol, error) {

	// anamnesis -> diagnosis -> protocol

	// protocol := entity.Protocol{
	// 	Actions: []entity.Action{
	// 		{
	// 			Drug:     "paracetomoll",
	// 			Dosage:   "200mg",
	// 			Time:     "twice of day",
	// 			Duration: "3 days",
	// 		},
	// 		{
	// 			Drug:     "water",
	// 			Dosage:   "2000 ml",
	// 			Time:     "per day",
	// 			Duration: "always",
	// 		},
	// 	},
	// }

	diagnosis, err := pc.webAPI.Ask(ctx, anamnesis.Description)
	if err != nil {
		return nil, err
	}

	protocol, err := pc.repo.GetProtocol(ctx, diagnosis)
	if err != nil {
		return nil, err
	}

	actions, err := pc.repo.GetProtocolActions(ctx, []byte(protocol.ID.String()))
	if err != nil {
		return nil, err
	}

	protocol.Actions = actions

	return protocol, nil
}
