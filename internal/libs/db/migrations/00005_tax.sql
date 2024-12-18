-- +goose Up
-- +goose StatementBegin
CREATE TABLE tax (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    tax_rate NUMERIC NOT NULL,
    effective_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO
    tax (id, name, tax_rate)
VALUES
    (
        '94b6c275-0293-477a-be70-2f05484f9f0a',
        'Dummy state tax',
        12.5
    ) ON CONFLICT (id) DO NOTHING;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tax;

-- +goose StatementEnd
