#!/bin/sh

goose -dir=$MIGRATIONS_PATH up
goose -dir=$MIGRATIONS_PATH reset 
goose -dir=$MIGRATIONS_PATH up
