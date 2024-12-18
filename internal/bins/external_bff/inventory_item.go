package external_bff

import (
	"context"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r mutationResolver) CreateInventoryItem(ctx context.Context, fields CreateInventoryItemInput) (*db.InventoryItem, error) {
	var createdInventoryItem *db.InventoryItem

	err := r.db.GetConn().Transaction(func(tx *gorm.DB) error {
		item, err := r.db.CreateInventoryItem(db.InventoryItem{
			Name: fields.Name,
		}, nil)
		if err != nil {
			return err
		}

		_, err = r.db.CreateInventoryItemPrice(db.InventoryItemPrice{
			InventoryItemId: item.ID,
			Price:           fields.IntroductionPrice,
			Version:         1,
		}, tx)
		if err != nil {
			return err
		}

		createdInventoryItem = item
		return nil
	})

	return createdInventoryItem, err
}

func (r mutationResolver) CreateInventoryItems(ctx context.Context, fields []*CreateInventoryItemInput) (bool, error) {
	err := r.db.GetConn().Transaction(func(tx *gorm.DB) error {
		for _, f := range fields {
			item, err := r.db.CreateInventoryItem(db.InventoryItem{
				Name: f.Name,
			}, nil)
			if err != nil {
				return err
			}

			_, err = r.db.CreateInventoryItemPrice(db.InventoryItemPrice{
				InventoryItemId: item.ID,
				Price:           f.IntroductionPrice,
				Version:         1,
			}, tx)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return true, err
}

func (r queryResolver) GetInventoryItem(ctx context.Context, id uuid.UUID) (*db.InventoryItem, error) {
	return r.db.GetInventoryItem(id, nil)
}

func (r queryResolver) GetAllInventoryItems(ctx context.Context) ([]*db.InventoryItem, error) {
	return r.db.GetAllInventoryItems(nil)
}

type inventoryResolver struct{ resolver }

func (r resolver) InventoryItem() InventoryItemResolver { return inventoryResolver{r} }

func (r inventoryResolver) CreatedAt(ctx context.Context, obj *db.InventoryItem) (string, error) {
	return obj.CreatedAt.String(), nil
}

func (r inventoryResolver) CurrentPrice(ctx context.Context, obj *db.InventoryItem) (*db.InventoryItemPrice, error) {
	return r.db.GetLatestInventoryItemPrice(obj.ID, nil)
}

func (r inventoryResolver) UpdatedAt(ctx context.Context, obj *db.InventoryItem) (string, error) {
	return obj.UpdatedAt.String(), nil
}
