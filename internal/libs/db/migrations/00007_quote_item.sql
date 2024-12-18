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

INSERT INTO
    quote_item (id, quote_id, item_id, item_price_id, quantity)
VALUES
    (
        'a13bb390-d78c-495b-b4ee-1fa6820b4423',
        'f3ae7d77-a325-4137-aae5-0f1e59db9665',
        'b00bdb3c-655b-4244-828f-0e518c2fc9ad',
        '9a11cfde-76eb-41d3-bab2-4c8f7790fa56',
        1
    ) ON CONFLICT (id) DO NOTHING;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS quote_item;

-- +goose StatementEnd
