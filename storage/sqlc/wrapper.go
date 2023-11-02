package sqlc

import (
	"context"
	"database/sql"
	"errors"

	"github.com/math-88/medic-online/internal/entity"

	"github.com/google/uuid"
)

type Storage struct {
	*Queries
}

func NewWrapper(q *Queries) *Storage {
	return &Storage{
		q,
	}
}

func (s *Storage) Save(ctx context.Context, user *entity.User) error {

	arg := CreateUserParams{
		user.Username,
		user.Role,
	}
	s.Queries.CreateUser(ctx, arg)

	return nil
}

func (s *Storage) GetByID(ctx context.Context, id string) (*entity.User, error) {

	user, err := s.GetUser(ctx, []byte(id))
	if err != nil {
		return nil, err
	}

	userID, err := uuid.ParseBytes(user.ID)
	if err != nil {
		return nil, err
	}

	entityUser := &entity.User{
		ID: userID,
	}

	return entityUser, nil
}

// @TODO
func (s *Storage) GetByName(ctx context.Context, username string) (*entity.User, error) {
	return nil, errors.New("method not implemented")
}

func (s *Storage) GetProtocol(ctx context.Context, diagnosis string) (*entity.Protocol, error) {

	protocol, err := s.Queries.GetProtocolByDiagnosis(ctx, diagnosis)
	if err != nil {
		return nil, err
	}

	protocolID, err := uuid.ParseBytes(protocol.ID)
	if err != nil {
		return nil, err
	}

	entityProtocol := &entity.Protocol{
		ID:          protocolID,
		Name:        protocol.Name,
		Description: protocol.Diagnosis,
	}

	actions, err := s.GetProtocolActions(ctx, protocol.ID)
	if err != nil {
		return nil, err
	}
	entityProtocol.Actions = actions

	return entityProtocol, nil
}

func (s *Storage) GetProtocolActions(ctx context.Context, protocolID []byte) ([]entity.Action, error) {

	id := sql.NullString{
		String: string(protocolID),
		Valid:  true,
	}

	actions, err := s.Queries.GetProtocolActions(ctx, id)
	if err != nil {
		return nil, err
	}

	entityActons := make([]entity.Action, 0, len(actions))
	for _, action := range actions {

		entityActon := entity.Action{
			// ID:
			Drug:     action.Drug,
			Dosage:   action.Dosage.String,
			Time:     action.Time.String,
			Duration: action.Description.String,
		}
		entityActons = append(entityActons, entityActon)
	}

	return entityActons, nil
}
