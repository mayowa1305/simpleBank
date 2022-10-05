CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "account_number" bigserial UNIQUE NOT NULL,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_number" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_number" bigint NOT NULL,
  "to_account_number" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
); 

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("account_number");

CREATE INDEX ON "transfers" ("from_account_number");

CREATE INDEX ON "transfers" ("to_account_number");

CREATE INDEX ON "transfers" ("from_account_number", "to_account_number");

COMMENT ON COLUMN "entries"."amount" IS 'it can be positve or negative value';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

ALTER TABLE "entries" ADD FOREIGN KEY ("account_number") REFERENCES "accounts" ("account_number");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_number") REFERENCES "accounts" ("account_number");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_number") REFERENCES "accounts" ("account_number");
