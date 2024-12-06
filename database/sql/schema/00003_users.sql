-- +goose Up
CREATE TABLE "users" (
    "id" ulid PRIMARY KEY,
    "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    "version" INTEGER DEFAULT 1 NOT NULL,
    "is_registered" BOOLEAN DEFAULT FALSE NOT NULL,
    "name" TEXT NOT NULL,
    "username" citext NOT NULL,
    "dob" date NOT NULL,
    "bio" TEXT,
    "website" TEXT,
    "location" TEXT,
    "avatar_path" TEXT,
    "banner_path" TEXT,
    "last_active_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    "primary_email_id" ulid UNIQUE NULLS NOT DISTINCT,
    "primary_phone_number_id" ulid UNIQUE NULLS NOT DISTINCT,
    "is_mfa_enabled" BOOLEAN DEFAULT FALSE NOT NULL,
    "is_bot" BOOLEAN NOT NULL DEFAULT FALSE,
    "profile_locked" BOOLEAN NOT NULL DEFAULT FALSE,
    "show_dob" BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE INDEX "users_version_idx" ON "users" ("version");

CREATE UNIQUE INDEX "users_username_registered_unique_idx" ON "users" ("username", "is_registered")
WHERE
    "is_registered";

-- +goose Down
DROP TABLE "users";