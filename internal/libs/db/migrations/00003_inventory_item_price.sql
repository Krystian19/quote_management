-- +goose Up
-- +goose StatementBegin
CREATE TABLE inventory_item_price (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    inventory_item_id UUID NOT NULL REFERENCES inventory_item(id),
    price NUMERIC NOT NULL,
    version BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON inventory_item_price (inventory_item_id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS inventory_item_price;

-- +goose StatementEnd
