-- +goose Up
-- +goose StatementBegin
CREATE TABLE quote (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    account_id UUID DEFAULT uuid_generate_v4(),
    payment_id UUID DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS quote;

-- +goose StatementEnd
