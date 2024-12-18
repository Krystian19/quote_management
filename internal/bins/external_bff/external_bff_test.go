package external_bff_test

import (
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Krystian19/quote_management/internal/bins/external_bff"
	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/Krystian19/quote_management/internal/libs/env"
	"github.com/Krystian19/quote_management/internal/libs/utils"
)

func getGqlClient(port string) graphql.Client {
	return graphql.NewClient(fmt.Sprintf("http://localhost:%s%s", port, external_bff.GQLEndpoint), nil)
}

func getDB() (*db.DB, error) {
	ev := env.GetEnv()

	return db.New(db.Opts{
		Url: ev.DB_TEST_URL,
	})
}

func getBffInstance(
	port string,
) (*external_bff.Server, error) {
	test_db, err := getDB()
	if err != nil {
		return nil, err
	}

	return utils.GetPtr(external_bff.New(external_bff.Opts{
		Port: port,
		DB:   test_db,
	})), nil
}
