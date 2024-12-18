-- +goose Up
-- +goose StatementBegin
CREATE TABLE quote_tax (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    quote_id UUID NOT NULL REFERENCES quote(id),
    tax_id UUID NOT NULL REFERENCES tax(id)
);

CREATE INDEX ON quote_tax (quote_id);

CREATE INDEX ON quote_tax (tax_id);

INSERT INTO
    quote_tax (id, quote_id, tax_id)
VALUES
    (
        'a13bb390-d78c-495b-b4ee-1fa6820b4423',
        'f3ae7d77-a325-4137-aae5-0f1e59db9665',
        '94b6c275-0293-477a-be70-2f05484f9f0a'
    ) ON CONFLICT (id) DO NOTHING;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS quote_tax;

-- +goose StatementEnd
