-- +goose Up
-- +goose StatementBegin
create table users (
                       id serial primary key,
                       email text unique,
                       phone text unique,
                       password text not null,
                       created_at timestamp not null default now(),
                       updated_at timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop talse if exists users;
-- +goose StatementEnd
