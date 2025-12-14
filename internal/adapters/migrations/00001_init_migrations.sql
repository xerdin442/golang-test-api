-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
  );

CREATE TABLE
  IF NOT EXISTS events (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    location VARCHAR(255) NOT NULL,
    datetime DATETIME NOT NULL,
    owner_id INT NOT NULL,
    CONSTRAINT fk_event_owner FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
  );

CREATE TABLE
  IF NOT EXISTS attendees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    event_id INT NOT NULL,
    CONSTRAINT fk_attendee_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_attendee_event FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE,
    UNIQUE KEY unique_attendance (user_id, event_id)
  );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS attendees;

-- +goose StatementEnd