-- +goose Up
CREATE TABLE
    "login_flows" ("id" SERIAL NOT NULL PRIMARY KEY, "name" TEXT NOT NULL);

-- +goose Down
DROP TABLE "login_flows";