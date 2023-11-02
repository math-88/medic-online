// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO user (
  name, role
) VALUES (
  ?, ?
)
`

type CreateUserParams struct {
	Name string
	Role string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.Name, arg.Role)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id []byte) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getProtocolActions = `-- name: GetProtocolActions :many
SELECT id, protocol_id, drug, dosage, time, duration, description FROM action
WHERE protocol_id = ?
ORDER BY id
`

func (q *Queries) GetProtocolActions(ctx context.Context, protocolID sql.NullString) ([]Action, error) {
	rows, err := q.db.QueryContext(ctx, getProtocolActions, protocolID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Action
	for rows.Next() {
		var i Action
		if err := rows.Scan(
			&i.ID,
			&i.ProtocolID,
			&i.Drug,
			&i.Dosage,
			&i.Time,
			&i.Duration,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProtocolByDiagnosis = `-- name: GetProtocolByDiagnosis :one
SELECT id, diagnosis, name, description FROM Protocol
WHERE diagnosis = ? LIMIT 1
`

func (q *Queries) GetProtocolByDiagnosis(ctx context.Context, diagnosis string) (Protocol, error) {
	row := q.db.QueryRowContext(ctx, getProtocolByDiagnosis, diagnosis)
	var i Protocol
	err := row.Scan(
		&i.ID,
		&i.Diagnosis,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const getProtocolByID = `-- name: GetProtocolByID :one
SELECT id, diagnosis, name, description FROM Protocol
WHERE id = ? LIMIT 1
`

// Protocol
func (q *Queries) GetProtocolByID(ctx context.Context, id []byte) (Protocol, error) {
	row := q.db.QueryRowContext(ctx, getProtocolByID, id)
	var i Protocol
	err := row.Scan(
		&i.ID,
		&i.Diagnosis,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, name, role, is_active, created_at, updated_at, deleted_at FROM user
WHERE id = ? LIMIT 1
`

// USER
func (q *Queries) GetUser(ctx context.Context, id []byte) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Role,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listProtocol = `-- name: ListProtocol :many
SELECT id, diagnosis, name, description FROM Protocol
ORDER BY name
`

func (q *Queries) ListProtocol(ctx context.Context) ([]Protocol, error) {
	rows, err := q.db.QueryContext(ctx, listProtocol)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Protocol
	for rows.Next() {
		var i Protocol
		if err := rows.Scan(
			&i.ID,
			&i.Diagnosis,
			&i.Name,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUser = `-- name: ListUser :many
SELECT id, name, role, is_active, created_at, updated_at, deleted_at FROM user
ORDER BY name
`

func (q *Queries) ListUser(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Role,
			&i.IsActive,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeUser = `-- name: RemoveUser :exec
UPDATE user SET name = ?, deleted_at = ?
WHERE id = ?
`

type RemoveUserParams struct {
	Name      string
	DeletedAt sql.NullTime
	ID        []byte
}

func (q *Queries) RemoveUser(ctx context.Context, arg RemoveUserParams) error {
	_, err := q.db.ExecContext(ctx, removeUser, arg.Name, arg.DeletedAt, arg.ID)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE user SET name = ?
WHERE id = ?
`

type UpdateUserParams struct {
	Name string
	ID   []byte
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser, arg.Name, arg.ID)
	return err
}