-- +goose Up
CREATE TABLE
    "phone_numbers" (
        "id" ulid PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
        "version" INTEGER DEFAULT 1 NOT NULL,
        "user_id" ulid NOT NULL,
        "value" citext NOT NULL UNIQUE,
        "is_primary" BOOLEAN DEFAULT FALSE NOT NULL,
        "is_verified" BOOLEAN DEFAULT FALSE NOT NULL,
        "last_verified_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL
    );

CREATE INDEX "phone_numbers_version_idx" ON "phone_numbers" ("version");

CREATE INDEX "phone_numbers_user_id_idx" ON "phone_numbers" ("user_id");

CREATE INDEX "phone_numbers_is_primary_idx" ON "phone_numbers" ("is_primary");

CREATE INDEX "phone_numbers_is_verified_idx" ON "phone_numbers" ("is_verified");

-- +goose Down
DROP TABLE "phone_numbers";