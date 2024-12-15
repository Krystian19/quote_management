package db_test

import (
	"testing"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestInventoryItemSupply(t *testing.T) {
	test_db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	dummyInventoryItem := db.InventoryItemFactory.MustCreate().(db.InventoryItem)
	createdInventoryItem, err := test_db.CreateInventoryItem(dummyInventoryItem, nil)
	if err != nil {
		t.Fatal(err)
	}

	dummyInventoryItemSupply := db.InventoryItemSupplyFactory.MustCreate().(db.InventoryItemSupply)
	dummyInventoryItemSupply.InventoryItemId = createdInventoryItem.ID
	createdInventoryItemSupply, err := test_db.CreateInventoryItemSupply(dummyInventoryItemSupply, nil)
	if err != nil {
		t.Fatal(err)
	}

	foundSupplies, err := test_db.GetAllInventoryItemSupply(nil)
	if err != nil {
		t.Fatal(err)
	}

	require.NotEmpty(t, foundSupplies)
	require.True(t,
		lo.ContainsBy(foundSupplies, func(item *db.InventoryItemSupply) bool {
			return item.ID == createdInventoryItemSupply.ID
		}),
	)
}
