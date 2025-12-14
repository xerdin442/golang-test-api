-- name: ListEvents :many
SELECT * FROM events
ORDER BY name;

-- name: GetEvent :one
SELECT * FROM events
WHERE id = ? LIMIT 1;

-- name: CreateEvent :execresult
INSERT INTO events (
  name, description, location, datetime
) VALUES (
  ?, ?, ?, ?
);

-- name: UpdateEvent :exec
UPDATE events
  SET name = ?,
  description = ?,
  location = ?,
  datetime = ?
WHERE id = ?;

-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = ?;