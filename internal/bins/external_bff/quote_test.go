package external_bff_test

import (
	"context"
	"testing"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/Krystian19/quote_management/internal/libs/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateQuote(t *testing.T) {
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

	dummyInventoryItem := db.InventoryItemFactory.MustCreate().(db.InventoryItem)
	createdInventoryItem, err := test_db.CreateInventoryItem(dummyInventoryItem, nil)
	if err != nil {
		t.Fatal(err)
	}

	dummyInventoryItemPrice := db.InventoryItemPriceFactory.MustCreate().(db.InventoryItemPrice)
	dummyInventoryItemPrice.InventoryItemId = createdInventoryItem.ID
	_, err = test_db.CreateInventoryItemPrice(dummyInventoryItemPrice, nil)
	if err != nil {
		t.Fatal(err)
	}

	dummyQuote := db.QuoteFactory.MustCreate().(db.Quote)

	createdQuoteRes, err := CreateQuote(context.Background(), getGqlClient(demoPort), CreateQuoteInput{
		AccountId: dummyQuote.AccountId,
		Items: []CreateQuoteItemsInput{
			{
				ItemId:   createdInventoryItem.ID,
				Quantity: 1,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	createdQuote := createdQuoteRes.CreateQuote
	foundQuote, err := test_db.GetQuote(createdQuote.Id, nil)
	if err != nil {
		t.Fatal(err)
	}

	require.NotNil(t, foundQuote)

	foundQuoteItems, err := test_db.GetQuoteItems(foundQuote.ID, nil)
	if err != nil {
		t.Fatal(err)
	}

	require.NotEmpty(t, foundQuoteItems)

	foundQuoteTaxes, err := test_db.GetQuoteTaxes(foundQuote.ID, nil)
	if err != nil {
		t.Fatal(err)
	}

	require.NotEmpty(t, foundQuoteTaxes)
}

func TestGetQuote(t *testing.T) {
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

	dummyInventoryItem := db.InventoryItemFactory.MustCreate().(db.InventoryItem)
	createdInventoryItem, err := test_db.CreateInventoryItem(dummyInventoryItem, nil)
	if err != nil {
		t.Fatal(err)
	}

	dummyInventoryItemPrice := db.InventoryItemPriceFactory.MustCreate().(db.InventoryItemPrice)
	dummyInventoryItemPrice.InventoryItemId = createdInventoryItem.ID
	_, err = test_db.CreateInventoryItemPrice(dummyInventoryItemPrice, nil)
	if err != nil {
		t.Fatal(err)
	}

	dummyQuote := db.QuoteFactory.MustCreate().(db.Quote)

	createdQuoteRes, err := CreateQuote(context.Background(), getGqlClient(demoPort), CreateQuoteInput{
		AccountId: dummyQuote.AccountId,
		Items: []CreateQuoteItemsInput{
			{
				ItemId:   createdInventoryItem.ID,
				Quantity: 1,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	createdQuote := createdQuoteRes.CreateQuote
	getQuoteRes, err := GetQuote(context.Background(), getGqlClient(demoPort), createdQuote.Id)
	if err != nil {
		t.Fatal(err)
	}

	foundQuote := getQuoteRes.GetQuote
	require.NotEmpty(t, foundQuote)
	require.Equal(nil, foundQuote.AccountId, dummyQuote.AccountId)
}
