-- +goose Up
-- +goose StatementBegin
CREATE TABLE quote_tax (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quote_id UUID NOT NULL REFERENCES quote(id),
    tax_id UUID NOT NULL REFERENCES tax_id(id),
);

CREATE INDEX ON quote_tax (tax_id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS quote_tax;

-- +goose StatementEnd
