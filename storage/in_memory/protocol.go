package in_memory

import (
	"context"
	"errors"

	"github.com/math-88/medic-online/internal/entity"
)

// Find protocols
func (store *InMemoryStore) GetProtocol(ctx context.Context, diagnosis string) (*entity.Protocol, error) {
	return nil, errors.New("not implemented")
}

// Find protocol actions
func (store *InMemoryStore) GetProtocolActions(ctx context.Context, protocolID []byte) ([]entity.Action, error) {
	return nil, errors.New("not implemented")
}
