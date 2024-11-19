-- +goose Up
CREATE TABLE
    "phone_numbers" (
        "id" ulid PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
        "updated_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
        "version" INTEGER DEFAULT 1 NOT NULL,
        "user_id" ulid NOT NULL,
        "value" citext NOT NULL UNIQUE,
        "is_primary" BOOLEAN DEFAULT FALSE NOT NULL,
        "is_verified" BOOLEAN DEFAULT FALSE NOT NULL,
        "last_verified_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
        "verification_code_sent_count" INTEGER NOT NULL,
        "verification_code_last_sent_at" TIMESTAMPTZ,
        "verification_code_sent_count_reset_at" TIMESTAMPTZ
    );

-- +goose Down
DROP TABLE "phone_numbers";