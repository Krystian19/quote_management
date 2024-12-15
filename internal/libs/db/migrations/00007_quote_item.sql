-- +goose Up
-- +goose StatementBegin
CREATE TABLE quote_item (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quote_id UUID NOT NULL REFERENCES quote(id),
    item_id UUID NOT NULL REFERENCES inventory_item(id),
    item_price_id UUID NOT NULL REFERENCES inventory_item_price(id),
    quantity BIGINT NOT NULL
);

CREATE INDEX ON quote_item (quote_id);

CREATE INDEX ON quote_item (item_id);

CREATE INDEX ON quote_item (item_price_id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS quote_item;

-- +goose StatementEnd
