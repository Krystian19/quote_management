package db_test

import (
	"testing"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/stretchr/testify/require"
)

func TestCreateInventoryItemPrice(t *testing.T) {
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

	dummyInventoryItemPrice.ID = createdInventoryItemPrice.ID
	dummyInventoryItemPrice.CreatedAt = createdInventoryItemPrice.CreatedAt

	require.Equal(t, dummyInventoryItemPrice, *createdInventoryItemPrice)
}

func TestGetLatestsInventoryItemPrice(t *testing.T) {
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

	foundItemPrice, err := test_db.GetLatestInventoryItemPrice(createdInventoryItem.ID, nil)
	if err != nil {
		t.Fatal(err)
	}

	require.NotNil(t, foundItemPrice)
	require.Equal(t, *createdInventoryItemPrice, *foundItemPrice)
}
