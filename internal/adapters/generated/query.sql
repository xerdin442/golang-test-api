-- name: ListEvents :many
SELECT
  *
FROM
  events
ORDER BY
  name;

-- name: GetEvent :one
SELECT
  *
FROM
  events
WHERE
  id = ?
LIMIT
  1;

-- name: CreateEvent :execresult
INSERT INTO
  events (name, description, location, datetime, owner_id)
VALUES
  (?, ?, ?, ?, ?);

-- name: UpdateEvent :execresult
UPDATE events
SET
  name = ?,
  description = ?,
  location = ?,
  datetime = ?,
  owner_id = ?
WHERE
  id = ?;

-- name: DeleteEvent :exec
DELETE FROM events
WHERE
  id = ?;

-- name: CreateUser :execresult
INSERT INTO
  users (name, email, password)
VALUES
  (?, ?, ?);

-- name: GetUserByID :one
SELECT
  *
FROM
  users
WHERE
  id = ?
LIMIT
  1;  

-- name: GetUserByEmail :one
SELECT
  *
FROM
  users
WHERE
  email = ?
LIMIT
  1;  

-- name: AddAttendee :execresult
INSERT INTO
  attendees (user_id, event_id)
VALUES
  (?, ?);

-- name: RemoveAttendee :exec
DELETE FROM attendees
WHERE
  user_id = ?
  AND event_id = ?;

-- name: GetEventAttendees :many
SELECT
  u.name,
  u.email
FROM
  users u
  JOIN attendees a ON u.id = a.user_id
WHERE
  a.event_id = ?;