-- +goose Up
-- +goose StatementBegin
CREATE TABLE inventory_item (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO
    inventory_item (id, name)
VALUES
    (
        'b00bdb3c-655b-4244-828f-0e518c2fc9ad',
        'Iphone 15'
    ) ON CONFLICT (id) DO NOTHING;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS inventory_item;

-- +goose StatementEnd
