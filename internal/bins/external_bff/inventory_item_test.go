package external_bff_test

import (
	"context"
	"testing"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/Krystian19/quote_management/internal/libs/utils"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestCreateInventoryItemTest(t *testing.T) {
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
	dummyInventoryItemPrice := db.InventoryItemPriceFactory.MustCreate().(db.InventoryItemPrice)

	createdItemRes, err := CreateInventoryItem(context.Background(), getGqlClient(demoPort), CreateInventoryItemInput{
		Name:              dummyInventoryItem.Name,
		IntroductionPrice: dummyInventoryItemPrice.Price,
	})
	if err != nil {
		t.Fatal(err)
	}

	createdItem := createdItemRes.CreateInventoryItem

	foundItem, err := test_db.GetInventoryItem(createdItem.Id, nil)
	if err != nil {
		t.Fatal(err)
	}

	require.NotNil(t, foundItem)
	require.Equal(t, foundItem.Name, dummyInventoryItem.Name)

	foundPrice, err := test_db.GetLatestInventoryItemPrice(createdItem.Id, nil)
	if err != nil {
		t.Fatal(err)
	}

	require.NotNil(t, foundPrice)
	require.Equal(t, foundPrice.Price, dummyInventoryItemPrice.Price)
}

func TestGetInventoryItemTest(t *testing.T) {
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
	dummyInventoryItemPrice := db.InventoryItemPriceFactory.MustCreate().(db.InventoryItemPrice)

	createdItemRes, err := CreateInventoryItem(context.Background(), getGqlClient(demoPort), CreateInventoryItemInput{
		Name:              dummyInventoryItem.Name,
		IntroductionPrice: dummyInventoryItemPrice.Price,
	})
	if err != nil {
		t.Fatal(err)
	}

	createdItem := createdItemRes.CreateInventoryItem
	foundItemRes, err := GetInventoryItem(context.Background(), getGqlClient(demoPort), createdItem.Id)
	if err != nil {
		t.Fatal(err)
	}

	foundItem := foundItemRes.GetInventoryItem
	require.NotNil(t, foundItem)

	foundDBItem, err := test_db.GetInventoryItem(createdItem.Id, nil)
	if err != nil {
		t.Fatal(err)
	}

	foundPrice, err := test_db.GetLatestInventoryItemPrice(createdItem.Id, nil)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, foundDBItem.ID, foundItem.Id)
	require.Equal(t, foundDBItem.Name, foundItem.Name)
	require.Equal(t, foundPrice.ID, foundItem.CurrentPrice.Id)
	require.Equal(t, foundDBItem.ID, foundItem.CurrentPrice.InventoryItemId)
	require.Equal(t, foundPrice.Price, foundItem.CurrentPrice.Price)
	require.Equal(t, foundPrice.Version, foundItem.CurrentPrice.Version)
}

func TestCreateInventoryItemsTest(t *testing.T) {
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

	dummyInventoryItem := db.InventoryItemFactory.MustCreate().(db.InventoryItem)
	dummyInventoryItemPrice := db.InventoryItemPriceFactory.MustCreate().(db.InventoryItemPrice)

	_, err = CreateInventoryItems(context.Background(), getGqlClient(demoPort), []CreateInventoryItemInput{
		{
			Name:              dummyInventoryItem.Name,
			IntroductionPrice: dummyInventoryItemPrice.Price,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	foundItemsRes, err := GetAllInventoryItems(context.Background(), getGqlClient(demoPort))
	if err != nil {
		t.Fatal(err)
	}

	foundItems := foundItemsRes.GetAllInventoryItems
	require.NotEmpty(t, foundItems)

	require.True(t,
		lo.ContainsBy(foundItems, func(item GetAllInventoryItemsGetAllInventoryItemsInventoryItem) bool {
			return item.Name == dummyInventoryItem.Name
		}),
	)
}

func TestUpdateInventoryItemPrice(t *testing.T) {
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
	dummyInventoryItemPrice := db.InventoryItemPriceFactory.MustCreate().(db.InventoryItemPrice)

	createdItemRes, err := CreateInventoryItem(context.Background(), getGqlClient(demoPort), CreateInventoryItemInput{
		Name:              dummyInventoryItem.Name,
		IntroductionPrice: dummyInventoryItemPrice.Price,
	})
	if err != nil {
		t.Fatal(err)
	}

	createdItem := createdItemRes.CreateInventoryItem
	newPrice := 19.2

	_, err = UpdateInventoryItemPrice(context.Background(), getGqlClient(demoPort), createdItem.Id, newPrice)
	if err != nil {
		t.Fatal(err)
	}

	foundPrice, err := test_db.GetLatestInventoryItemPrice(createdItem.Id, nil)
	if err != nil {
		t.Fatal(err)
	}

	require.NotNil(t, foundPrice)
	require.Equal(t, foundPrice.Price, newPrice)
	require.Equal(t, foundPrice.Version, 2)
}
