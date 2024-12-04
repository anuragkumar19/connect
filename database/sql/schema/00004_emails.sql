-- +goose Up
CREATE TABLE "emails" (
    "id" ulid PRIMARY KEY NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    "version" INTEGER DEFAULT 1 NOT NULL,
    "user_id" ulid NOT NULL,
    "value" citext NOT NULL UNIQUE,
    "is_primary" BOOLEAN DEFAULT FALSE NOT NULL,
    "is_verified" BOOLEAN DEFAULT FALSE NOT NULL,
    "last_verified_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

CREATE INDEX "emails_version_idx" ON "emails" ("version");

CREATE INDEX "emails_user_id_idx" ON "emails" ("user_id");

CREATE INDEX "emails_is_primary_idx" ON "emails" ("is_primary");

CREATE INDEX "emails_is_verified_idx" ON "emails" ("is_verified");

-- +goose Down
DROP TABLE "emails";