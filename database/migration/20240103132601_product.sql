-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE products (
  id UUID NOT NULL DEFAULT uuid_generate_v4(),
  name varchar(255),
  description text,
  price decimal,
  created_at int NOT NULL,
  updated_at int NOT NULL,
  deleted_at int NOT NULL,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd
