-- +goose Up
-- +goose StatementBegin
-- Allows usage of uuids
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP EXTENSION IF EXISTS "uuid-ossp";

DROP EXTENSION IF EXISTS "pgcrypto";

-- +goose StatementEnd
