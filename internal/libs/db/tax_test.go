package db_test

import (
	"testing"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestCreateTax(t *testing.T) {
	test_db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	dummyTax := db.TaxFactory.MustCreate().(db.Tax)
	createdTax, err := test_db.CreateTax(dummyTax, nil)
	if err != nil {
		t.Fatal(err)
	}

	dummyTax.ID = createdTax.ID
	dummyTax.EffectiveAt = createdTax.EffectiveAt
	dummyTax.CreatedAt = createdTax.CreatedAt
	dummyTax.UpdatedAt = createdTax.UpdatedAt

	require.Equal(t, dummyTax, *createdTax)
}

func TestGetAllTaxes(t *testing.T) {
	test_db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	dummyTax := db.TaxFactory.MustCreate().(db.Tax)
	createdTax, err := test_db.CreateTax(dummyTax, nil)
	if err != nil {
		t.Fatal(err)
	}

	found_taxes, err := test_db.GetAllTaxes(nil)
	if err != nil {
		t.Fatal(err)
	}

	require.NotEmpty(t, found_taxes)

	require.True(t,
		lo.ContainsBy(found_taxes, func(tax *db.Tax) bool {
			return tax.ID == createdTax.ID
		}),
	)
}