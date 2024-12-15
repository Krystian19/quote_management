package db_test

import (
	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/Krystian19/quote_management/internal/libs/env"
)

func getDB() (*db.DB, error) {
	ev := env.GetEnv()

	return db.New(db.Opts{
		Url: ev.DB_TEST_URL,
	})
}
