-- +goose Up
CREATE TABLE "verification_tokens" (
    "token" TEXT PRIMARY KEY NOT NULL,
    "id" ulid UNIQUE NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    "expire_at" TIMESTAMPTZ NOT NULL,
    "user_id" ulid NOT NULL,
    "email_id" ulid,
    "phone_number_id" ulid,
    "otp" TEXT NOT NULL
);

CREATE INDEX "verification_tokens_user_id_idx" ON "verification_tokens" ("user_id");

CREATE INDEX "verification_tokens_expire_at_idx" ON "verification_tokens" ("expire_at");

-- +goose Down
DROP TABLE "verification_tokens";