package external_bff

import (
	"context"

	"github.com/Krystian19/quote_management/internal/libs/db"
)

type inventoryItemPriceResolver struct{ resolver }

func (r resolver) InventoryItemPrice() InventoryItemPriceResolver {
	return inventoryItemPriceResolver{r}
}

func (r inventoryItemPriceResolver) CreatedAt(ctx context.Context, obj *db.InventoryItemPrice) (string, error) {
	return obj.CreatedAt.String(), nil
}
