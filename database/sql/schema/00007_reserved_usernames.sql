-- +goose Up
CREATE TABLE "reserved_usernames" ("username" citext PRIMARY KEY NOT NULL);

-- +goose Down
DROP TABLE "reserved_usernames";