-- +goose Up
-- +goose StatementBegin
CREATE TABLE readers
(
    id       UUID NOT NULL PRIMARY KEY,
    email    TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    born_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE readers;
-- +goose StatementEnd