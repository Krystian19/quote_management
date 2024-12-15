package db_test

import (
	"testing"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestCreateInventoryItem(t *testing.T) {
	test_db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	dummyInventoryItem := db.InventoryItemFactory.MustCreate().(db.InventoryItem)
	createdInventoryItem, err := test_db.CreateInventoryItem(dummyInventoryItem, nil)
	if err != nil {
		t.Fatal(err)
	}

	dummyInventoryItem.ID = createdInventoryItem.ID
	dummyInventoryItem.CreatedAt = createdInventoryItem.CreatedAt
	dummyInventoryItem.UpdatedAt = createdInventoryItem.UpdatedAt

	require.Equal(t, dummyInventoryItem, *createdInventoryItem)
}

func TestGetInventoryItem(t *testing.T) {
	test_db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	dummyInventoryItem := db.InventoryItemFactory.MustCreate().(db.InventoryItem)
	createdInventoryItem, err := test_db.CreateInventoryItem(dummyInventoryItem, nil)
	if err != nil {
		t.Fatal(err)
	}

	foundInventoryItem, err := test_db.GetInventoryItem(createdInventoryItem.ID, nil)
	if err != nil {
		t.Fatal(err)
	}

	require.NotNil(t, foundInventoryItem)
	require.Equal(t, *createdInventoryItem, *foundInventoryItem)
}

func TestGetAllInventoryItems(t *testing.T) {
	test_db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	dummyInventoryItem := db.InventoryItemFactory.MustCreate().(db.InventoryItem)
	createdInventoryItem, err := test_db.CreateInventoryItem(dummyInventoryItem, nil)
	if err != nil {
		t.Fatal(err)
	}

	foundItems, err := test_db.GetAllInventoryItems(nil)
	if err != nil {
		t.Fatal(err)
	}

	require.NotEmpty(t, foundItems)
	require.True(t,
		lo.ContainsBy(foundItems, func(item *db.InventoryItem) bool {
			return item.ID == createdInventoryItem.ID
		}),
	)
}
