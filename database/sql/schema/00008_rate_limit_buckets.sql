-- +goose Up
CREATE TABLE "rate_limit_buckets" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    "last_reset_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    "consumed" BIGINT NOT NULL,
    "last_consumed_at" TIMESTAMPTZ
);

-- +goose Down
DROP TABLE "rate_limit_buckets";