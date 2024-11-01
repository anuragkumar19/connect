-- +goose Up
CREATE TYPE "registration_flows_state" AS ENUM('invalid_fields', 'verification_code_sent', 'registration_completed');

CREATE TABLE
    "registration_flows" (
        "id" BIGSERIAL PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
        "updated_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
        "version" INTEGER DEFAULT 1 NOT NULL,
        "expire_at" TIMESTAMPTZ NOT NULL,
        "csrf_token" TEXT NOT NULL,
        "state" "registration_flows_state" NOT NULL,
        -- email
        "email_value" TEXT,
        "email_is_valid" BOOLEAN NOT NULL DEFAULT FALSE,
        "email_error_message" TEXT,
        -- email_code
        "email_code" TEXT,
        "email_code_expiry" TIMESTAMPTZ,
        "email_code_incorrect_verification_attempts" INTEGER DEFAULT 0 NOT NULL,
        "email_code_remaining_verification_attempts" INTEGER DEFAULT 0 NOT NULL,
        "email_code_is_verified" BOOLEAN NOT NULL DEFAULT FALSE,
        "email_code_error_message" TEXT,
        -- phone_number
        "phone_number_value" TEXT,
        "phone_number_is_valid" BOOLEAN NOT NULL DEFAULT FALSE,
        "phone_number_error_message" TEXT,
        -- phone_number_code
        "phone_number_code" TEXT,
        "phone_number_code_expiry" TIMESTAMPTZ,
        "phone_number_code_incorrect_verification_attempts" INTEGER DEFAULT 0 NOT NULL,
        "phone_number_code_remaining_verification_attempts" INTEGER DEFAULT 0 NOT NULL,
        "phone_number_code_is_verified" BOOLEAN NOT NULL DEFAULT FALSE,
        "phone_number_code_error_message" TEXT,
        -- name
        "name_value" TEXT,
        "name_is_valid" BOOLEAN NOT NULL DEFAULT FALSE,
        "name_error_message" TEXT,
        -- username
        "username_value" TEXT,
        "username_is_valid" BOOLEAN NOT NULL DEFAULT FALSE,
        "username_error_message" TEXT,
        -- password
        "password_value" TEXT,
        "password_is_valid" BOOLEAN NOT NULL DEFAULT FALSE,
        "password_error_message" TEXT
    );

-- +goose Down
DROP TABLE "registration_flows";

DROP TYPE "registration_flows_state";