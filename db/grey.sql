CREATE TYPE "transaction_type" AS ENUM (
  'CREDIT',
  'DEBIT'
);

CREATE TYPE "transaction_status" AS ENUM (
  'PENDING',
  'FAILED',
  'SUCCESSFUL'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "wallets" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "balance" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "type" transaction_type NOT NULL,
  "status" transaction_status NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "wallets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "wallets" ("user_id");

CREATE INDEX ON "transactions" ("user_id");

CREATE INDEX ON "transactions" ("type");

CREATE INDEX ON "transactions" ("status");

-- Seed
INSERT INTO "users" ("id", "name") 
VALUES (1, 'Charles');

INSERT INTO "wallets" ("id", "user_id", "balance") 
VALUES (1, 1, 100000);
