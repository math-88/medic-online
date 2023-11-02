-- USER
-- name: GetUser :one
SELECT * FROM user
WHERE id = ? LIMIT 1;

-- name: ListUser :many
SELECT * FROM user
ORDER BY name;

-- name: CreateUser :exec
INSERT INTO user (
  name, role
) VALUES (
  ?, ?
);

-- name: UpdateUser :exec
UPDATE user SET name = ?
WHERE id = ?;

-- name: RemoveUser :exec
UPDATE user SET name = ?, deleted_at = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?;



-- Protocol
-- name: GetProtocolByID :one
SELECT * FROM Protocol
WHERE id = ? LIMIT 1;

-- name: GetProtocolByDiagnosis :one
SELECT * FROM Protocol
WHERE diagnosis = ? LIMIT 1;

-- name: ListProtocol :many
SELECT * FROM Protocol
ORDER BY name;

-- name: GetProtocolActions :many
SELECT * FROM action
WHERE protocol_id = ?
ORDER BY id;