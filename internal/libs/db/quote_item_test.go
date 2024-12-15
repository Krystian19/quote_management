package db_test

import (
	"testing"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestCreateQuoteItem(t *testing.T) {
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
	createdInventoryItemPrice, err := test_db.CreateInventoryItemPrice(dummyInventoryItemPrice, nil)
	if err != nil {
		t.Fatal(err)
	}

	dummyQuote := db.QuoteFactory.MustCreate().(db.Quote)
	createdQuote, err := test_db.CreateQuote(dummyQuote, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = test_db.CreateQuoteItems([]*db.QuoteItem{
		{
			QuoteId:     createdQuote.ID,
			ItemId:      createdInventoryItem.ID,
			ItemPriceId: createdInventoryItemPrice.ID,
			Quantity:    1,
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}

	foundQuoteItems, err := test_db.GetQuoteItems(createdQuote.ID, nil)
	if err != nil {
		t.Fatal(err)
	}

	require.NotEmpty(t, foundQuoteItems)

	require.True(t,
		lo.ContainsBy(foundQuoteItems, func(item *db.QuoteItem) bool {
			return item.QuoteId == createdQuote.ID
		}),
	)
}
