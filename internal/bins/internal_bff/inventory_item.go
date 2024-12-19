package internal_bff

import (
	"io"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/Krystian19/quote_management/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

func (s server) BulkCreateInventoryItems(stream proto.InternalBffAPI_BulkCreateInventoryItemsServer) error {
	newItemsCh := make(chan *proto.NewInventoryItem)

	go func() {
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				break
			}

			newItemsCh <- req
		}

		close(newItemsCh)
	}()

	for rawItem := range newItemsCh {
		err := s.db.GetConn().Transaction(func(tx *gorm.DB) error {
			item, err := s.db.CreateInventoryItem(db.InventoryItem{
				Name: rawItem.Name,
			}, nil)
			if err != nil {
				return err
			}

			_, err = s.db.CreateInventoryItemPrice(db.InventoryItemPrice{
				InventoryItemId: item.ID,
				Price:           float64(rawItem.IntroductionPrice),
				Version:         1,
			}, tx)
			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			return err
		}
	}

	return stream.SendAndClose(&emptypb.Empty{})
}
