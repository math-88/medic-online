package entity

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User contains user's information
type User struct {
	ID             uuid.UUID
	Username       string
	HashedPassword string
	Role           string
}

// NewUser returns a new user
func NewUser(id uuid.UUID, username string, password string, role string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	user := &User{
		ID:             id,
		Username:       username,
		HashedPassword: string(hashedPassword),
		Role:           role,
	}

	return user, nil
}

// IsCorrectPassword checks if the provided password is correct or not
func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}

// Clone returns a clone of this user
func (user *User) Clone() *User {
	return &User{
		ID:             user.ID,
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		Role:           user.Role,
	}
}
