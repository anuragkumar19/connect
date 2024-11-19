-- +goose Up
CREATE TABLE
    "passwords" (
        "id" ulid PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
        "abandoned_at" TIMESTAMPTZ,
        "is_active" BOOLEAN NOT NULL,
        "user_id" ulid NOT NULL,
        "value" TEXT NOT NULL,
        "incorrect_attempt_count" INTEGER NOT NULL,
        "last_incorrect_attempt_at" TIMESTAMPTZ,
        "incorrect_attempt_count_reset_at" TIMESTAMPTZ
    );

-- +goose Down
DROP TABLE "passwords";