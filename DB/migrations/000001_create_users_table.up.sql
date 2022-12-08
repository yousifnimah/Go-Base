CREATE TABLE IF NOT EXISTS users
(
    id         bigserial           not null unique primary key,
    username   varchar(255) unique null,
    password   varchar(255)        null,
    full_name  varchar(255)        null,
    created_at timestamp           null,
    updated_at timestamp           null,
    deleted_at timestamp           null
);

CREATE INDEX index_username_users ON users (username);
CREATE INDEX index_created_at_users ON users (created_at);