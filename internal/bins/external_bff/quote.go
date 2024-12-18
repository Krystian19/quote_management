package external_bff

import (
	"context"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/google/uuid"
)

func (r queryResolver) GetQuote(ctx context.Context, id uuid.UUID) (*db.Quote, error) {
	return r.db.GetQuote(id, nil)
}

type quoteResolver struct{ resolver }

func (r resolver) Quote() QuoteResolver { return quoteResolver{r} }

func (r quoteResolver) CreatedAt(ctx context.Context, obj *db.Quote) (string, error) {
	return obj.CreatedAt.String(), nil
}

func (r quoteResolver) UpdatedAt(ctx context.Context, obj *db.Quote) (string, error) {
	return obj.UpdatedAt.String(), nil
}
