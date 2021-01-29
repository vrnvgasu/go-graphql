-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS notes (
    id         bigserial primary key,
    text       text not null default '',
    user_id    integer not null
);

CREATE TABLE IF NOT EXISTS users (
    id         bigserial primary key,
    name       text not null,
    age        integer not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notes;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
