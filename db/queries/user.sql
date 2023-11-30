-- name: GetUser :one
SELECT * FROM "user"
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "user"
ORDER BY name;

-- name: CreateUser :one
INSERT INTO "user" (
    id, name, email
) VALUES (
             $1, $2, $3
         ) RETURNING *;

-- name: UpdateUser :exec
UPDATE "user"
set name = $2,
    email = $3
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;