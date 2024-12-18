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

INSERT INTO
    inventory_item_price (id, inventory_item_id, price, version)
VALUES
    (
        '9a11cfde-76eb-41d3-bab2-4c8f7790fa56',
        'b00bdb3c-655b-4244-828f-0e518c2fc9ad',
        19.99,
        1
    ) ON CONFLICT (id) DO NOTHING;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS inventory_item_price;

-- +goose StatementEnd
