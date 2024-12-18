package external_bff_test

import (
	"context"
	"testing"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/Krystian19/quote_management/internal/libs/utils"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestCreateTax(t *testing.T) {
	demoPort := utils.GetRandOpenPort()

	srv, err := getBffInstance(
		demoPort,
	)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		if err := srv.Run(); err != nil {
			t.Fatal(err)
		}
	}()

	test_db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	dummyTax := db.TaxFactory.MustCreate().(db.Tax)

	createTaxRes, err := CreateTax(context.Background(), getGqlClient(demoPort), CreateTaxInput{
		Name:    dummyTax.Name,
		TaxRate: dummyTax.TaxRate,
	})
	if err != nil {
		t.Fatal(err)
	}

	createdTax := createTaxRes.CreateTax
	require.NotEmpty(t, createdTax.GetId().String())
	require.Equal(t, dummyTax.Name, createdTax.Name)
	require.Equal(t, dummyTax.TaxRate, createdTax.TaxRate)

	foundTax, err := test_db.GetTax(createdTax.GetId(), nil)
	if err != nil {
		t.Fatal(err)
	}

	require.NotNil(t, foundTax)
}

func TestGetAllTaxes(t *testing.T) {
	demoPort := utils.GetRandOpenPort()

	srv, err := getBffInstance(
		demoPort,
	)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		if err := srv.Run(); err != nil {
			t.Fatal(err)
		}
	}()

	dummyTax := db.TaxFactory.MustCreate().(db.Tax)
	createTaxRes, err := CreateTax(context.Background(), getGqlClient(demoPort), CreateTaxInput{
		Name:    dummyTax.Name,
		TaxRate: dummyTax.TaxRate,
	})
	if err != nil {
		t.Fatal(err)
	}

	createdTax := createTaxRes.CreateTax
	getAllTaxesRes, err := getAllTaxes(context.Background(), getGqlClient(demoPort))
	if err != nil {
		t.Fatal(err)
	}

	foundTaxes := getAllTaxesRes.GetAllTaxes

	require.NotEmpty(t, foundTaxes)
	require.True(t,
		lo.ContainsBy(foundTaxes, func(tax getAllTaxesGetAllTaxesTax) bool {
			return tax.Id == createdTax.Id
		}),
	)
}
