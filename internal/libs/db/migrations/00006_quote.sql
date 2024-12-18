-- +goose Up
-- +goose StatementBegin
CREATE TABLE quote (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    account_id UUID DEFAULT uuid_generate_v4(),
    payment_id UUID DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO
    quote (id, account_id)
VALUES
    (
        'f3ae7d77-a325-4137-aae5-0f1e59db9665',
        'c1f39078-3043-4a0d-a4be-6e8c1a70e5a6'
    ) ON CONFLICT (id) DO NOTHING;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS quote;

-- +goose StatementEnd
