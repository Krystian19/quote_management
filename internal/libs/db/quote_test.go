package db_test

import (
	"testing"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestCreateQuote(t *testing.T) {
	test_db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	dummyQuote := db.QuoteFactory.MustCreate().(db.Quote)
	created_quote, err := test_db.CreateQuote(dummyQuote, nil)
	if err != nil {
		t.Fatal(err)
	}

	dummyQuote.ID = created_quote.ID
	dummyQuote.CreatedAt = created_quote.CreatedAt
	dummyQuote.UpdatedAt = created_quote.UpdatedAt

	require.Equal(t, dummyQuote, *created_quote)
}

func TestGetQuote(t *testing.T) {
	test_db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	dummyQuote := db.QuoteFactory.MustCreate().(db.Quote)
	createdQuote, err := test_db.CreateQuote(dummyQuote, nil)
	if err != nil {
		t.Fatal(err)
	}

	foundQuote, err := test_db.GetQuote(createdQuote.ID, nil)
	if err != nil {
		t.Fatal(err)
	}

	require.NotNil(t, foundQuote)
	require.Equal(t, *createdQuote, *foundQuote)
}

func TestGetQuotes(t *testing.T) {
	test_db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	dummyQuote := db.QuoteFactory.MustCreate().(db.Quote)
	createdQuote, err := test_db.CreateQuote(dummyQuote, nil)
	if err != nil {
		t.Fatal(err)
	}

	foundQuotes, err := test_db.GetAllQuotes(nil)
	if err != nil {
		t.Fatal(err)
	}

	require.NotEmpty(t, foundQuotes)

	require.True(t,
		lo.ContainsBy(foundQuotes, func(quote *db.Quote) bool {
			return quote.ID == createdQuote.ID
		}),
	)
}
