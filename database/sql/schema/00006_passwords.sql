-- +goose Up
CREATE TABLE "passwords" (
    "id" BIGSERIAL PRIMARY KEY NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    "version" INTEGER DEFAULT 1 NOT NULL,
    "abandoned_at" TIMESTAMPTZ,
    "is_active" BOOLEAN NOT NULL,
    "user_id" ulid NOT NULL,
    "value" TEXT NOT NULL
);

CREATE INDEX "passwords_version_idx" ON "passwords" ("version");

CREATE INDEX "passwords_user_id_idx" ON "passwords" ("user_id");

CREATE INDEX "passwords_is_active_idx" ON "passwords" ("is_active");

-- +goose Down
DROP TABLE "passwords";