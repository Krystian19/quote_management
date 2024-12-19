package internal_bff_test

import (
	"fmt"

	"github.com/Krystian19/quote_management/internal/bins/internal_bff"
	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/Krystian19/quote_management/internal/libs/env"
	"github.com/Krystian19/quote_management/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getDB() (*db.DB, error) {
	ev := env.GetEnv()

	return db.New(db.Opts{
		Url: ev.DB_TEST_URL,
	})
}

func getInternalBffInstance(
	port int,
) (*internal_bff.InternalBff, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	return internal_bff.New(internal_bff.Opts{
		Port: port,
		DB:   db,
	})
}

func getInternalBffClient(port int) (proto.InternalBffAPIClient, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return proto.NewInternalBffAPIClient(conn), nil
}
