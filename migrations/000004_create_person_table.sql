-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS persons
(
    id         SERIAL PRIMARY KEY NOT NULL,
    name       VARCHAR(255)       NOT NULL,
    surname    VARCHAR(255)       NOT NULL,
    patronymic VARCHAR(255)       NOT NULL,
    age        INTEGER            NOT NULL,
    sex        VARCHAR(255)       NOT NULL,
    nation     VARCHAR(255)       NOT NULL,

    created_at timestamp(0) NOT NULL DEFAULT now(),
    updated_at timestamp(0) NOT NULL DEFAULT now(),
    deleted_at timestamp(0)          DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS persons;
-- +goose StatementEnd
