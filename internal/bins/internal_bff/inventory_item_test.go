package internal_bff_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/Krystian19/quote_management/internal/libs/utils"
	"github.com/Krystian19/quote_management/proto"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestCreateTax(t *testing.T) {
	demoPort, _ := strconv.Atoi(utils.GetRandOpenPort())

	srv, err := getInternalBffInstance(demoPort)
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

	client, err := getInternalBffClient(demoPort)
	if err != nil {
		t.Fatal(err)
	}

	stream, err := client.BulkCreateInventoryItems(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	dummyInventoryItem := db.InventoryItemFactory.MustCreate().(db.InventoryItem)
	dummyInventoryItemPrice := db.InventoryItemPriceFactory.MustCreate().(db.InventoryItemPrice)

	stream.Send(&proto.NewInventoryItem{
		Name:              dummyInventoryItem.Name,
		IntroductionPrice: float32(dummyInventoryItemPrice.Price),
	})

	stream.CloseSend()

	// we have to give some time for the stream to be processed
	time.Sleep(time.Second)

	foundItems, err := test_db.GetAllInventoryItems(nil)
	if err != nil {
		t.Fatal(err)
	}

	require.True(t,
		lo.ContainsBy(foundItems, func(item *db.InventoryItem) bool {
			return item.Name == dummyInventoryItem.Name
		}),
	)
}
