-- +goose Up
CREATE TABLE
    "rate_limit_buckets" (
        "id" TEXT NOT NULL PRIMARY KEY,
        "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
        "last_reset_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
        "version" INTEGER DEFAULT 1 NOT NULL,
        "consumed" BIGINT NOT NULL,
        "last_consumed_at" TIMESTAMPTZ
    );

CREATE INDEX "rate_limit_buckets_version_idx" ON "rate_limit_buckets" ("version");

-- +goose Down
DROP TABLE "rate_limit_buckets";