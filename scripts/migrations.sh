#!/bin/sh

# Build the migrations image
docker build -f ./docker/migrations/Dockerfile -t quote/migrations .

export MIGRATIONS_PATH="./internal/libs/db/migrations"
export MAIN_DBSTRING="host=quote.postgres port=5432 user=postgres password=pass dbname=quote sslmode=disable"
export TEST_DBSTRING="host=quote.postgres port=5432 user=postgres password=pass dbname=quote_test sslmode=disable"

case $1 in
    "test")
        echo "===================================================================="
        echo "DB test migrations"
        echo "===================================================================="

        export CMD="./internal/libs/db/scripts/run_migrations.sh"
        export GOOSE_DBSTRING="$TEST_DBSTRING"
    ;;
    "test_redo")
        echo "===================================================================="
        echo "DB up/down/up test migrations"
        echo "===================================================================="

        export CMD="./internal/libs/db/scripts/run_redo_migrations.sh"
        export GOOSE_DBSTRING="$TEST_DBSTRING"
    ;;
    "redo")
        echo "===================================================================="
        echo "DB up/down/up migrations"
        echo "===================================================================="

        export CMD="./internal/libs/db/scripts/run_redo_migrations.sh"
        export GOOSE_DBSTRING="$MAIN_DBSTRING"
    ;;
    *)
        echo "===================================================================="
        echo "DB migrations"
        echo "===================================================================="

        export CMD="./internal/libs/db/scripts/run_migrations.sh"
        export GOOSE_DBSTRING="$MAIN_DBSTRING"
    ;;
esac

docker run -t -v ${PWD}:/go/src/app --network="canary_net" -e MIGRATIONS_PATH=${MIGRATIONS_PATH} -e GOOSE_DBSTRING="${GOOSE_DBSTRING}" quote/migrations $CMD 
