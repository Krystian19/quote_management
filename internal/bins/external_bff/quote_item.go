package external_bff

import (
	"context"

	"github.com/Krystian19/quote_management/internal/libs/db"
)

type quoteItemResolver struct{ resolver }

func (r resolver) QuoteItem() QuoteItemResolver {
	return quoteItemResolver{r}
}

func (r quoteItemResolver) Item(ctx context.Context, obj *db.QuoteItem) (*db.InventoryItem, error) {
	return r.db.GetInventoryItem(obj.ItemId, nil)
}

func (r quoteItemResolver) ItemPrice(ctx context.Context, obj *db.QuoteItem) (*db.InventoryItemPrice, error) {
	return r.db.GetInventoryItemPrice(obj.ItemPriceId, nil)
}
