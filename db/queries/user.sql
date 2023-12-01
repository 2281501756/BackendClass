-- name: GetUser :one
SELECT * FROM "user"
WHERE id = $1 LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM "user"
WHERE id = $1 LIMIT 1 for update;

-- name: ListUsers :many
SELECT * FROM "user"
ORDER BY name;

-- name: CreateUser :one
INSERT INTO "user" (
    id, name, email, money
) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;

-- name: UpdateUser :exec
UPDATE "user"
set name = $2,
    email = $3
WHERE id = $1;

-- name: UpdateUserMoney :one
update "user"
set money = money + sqlc.arg(money)
where id = sqlc.arg(id) returning *;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;