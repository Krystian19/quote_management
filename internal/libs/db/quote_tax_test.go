package db_test

import (
	"testing"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestCreateQuoteTax(t *testing.T) {
	test_db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	dummyTax := db.TaxFactory.MustCreate().(db.Tax)
	createdTax, err := test_db.CreateTax(dummyTax, nil)
	if err != nil {
		t.Fatal(err)
	}

	dummyQuote := db.QuoteFactory.MustCreate().(db.Quote)
	createdQuote, err := test_db.CreateQuote(dummyQuote, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = test_db.CreateQuoteTaxes([]*db.QuoteTax{
		{
			QuoteId: createdQuote.ID,
			TaxId:   createdTax.ID,
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}

	foundQuoteTaxes, err := test_db.GetQuoteTaxes(createdQuote.ID, nil)
	if err != nil {
		t.Fatal(err)
	}

	require.NotEmpty(t, foundQuoteTaxes)

	require.True(t,
		lo.ContainsBy(foundQuoteTaxes, func(tax *db.QuoteTax) bool {
			return tax.QuoteId == createdQuote.ID
		}),
	)
}
