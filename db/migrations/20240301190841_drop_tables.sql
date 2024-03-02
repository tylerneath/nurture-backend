-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS messages;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd