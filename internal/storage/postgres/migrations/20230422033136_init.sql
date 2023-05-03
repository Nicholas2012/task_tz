-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	login VARCHAR(255) NOT NULL,
	data JSONB NOT NULL DEFAULT '{}',
	CONSTRAINT users_pk PRIMARY KEY (login)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
