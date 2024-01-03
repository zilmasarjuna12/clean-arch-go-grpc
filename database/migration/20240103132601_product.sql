-- +goose Up
-- +goose StatementBegin
CREATE TABLE products (
  id varchar(36) NOT NULL,
  name varchar(255),
  description text,
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
