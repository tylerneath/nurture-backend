-- +goose Up
-- +goose StatementBegin
INSERT INTO
    message_types (id, created_at, updated_at, type)
VALUES
    (uuid_generate_v4(), now(), now(), 'Log'),
    (uuid_generate_v4(), now(), now(), 'Reminder'),
    (uuid_generate_v4(), now(), now(), 'Reflect');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DELETE FROM
    message_types
WHERE
    type = 'Log'
    AND type = 'Reminder'
    AND type = 'Reflect';

-- +goose StatementEnd