package in_memory

import (
	"context"
	"errors"
	"sync"

	"github.com/math-88/medic-online/internal/entity"
)

// InMemoryStore stores users in memory
type InMemoryStore struct {
	mutex    sync.RWMutex
	users    map[string]*entity.User
	protocol map[string]*entity.Protocol
	actions  map[string]*entity.Action
}

// NewInMemoryStore returns a new in-memory user store
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		users: make(map[string]*entity.User),
	}
}

// Save saves a user to the store
func (store *InMemoryStore) Save(ctx context.Context, user *entity.User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.users[user.Username] != nil {
		return errors.New("user already exist")
	}

	store.users[user.Username] = user.Clone()
	return nil
}

// Find finds a user by ID
func (store *InMemoryStore) GetByID(ctx context.Context, id string) (*entity.User, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	user := store.users[id]
	if user == nil {
		return nil, errors.New("user not found")
	}

	return user.Clone(), nil
}

// Find finds a user by username
func (store *InMemoryStore) GetByName(ctx context.Context, username string) (*entity.User, error) {
	return store.GetByID(ctx, username)
}
