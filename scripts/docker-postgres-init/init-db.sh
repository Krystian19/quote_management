#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE quote;
    CREATE DATABASE quote_test;

    GRANT ALL PRIVILEGES ON DATABASE quote TO postgres;
    GRANT ALL PRIVILEGES ON DATABASE quote_test TO postgres;
EOSQL
