package external_bff

import (
	"context"

	"github.com/Krystian19/quote_management/internal/libs/db"
)

type quoteTaxResolver struct{ resolver }

func (r resolver) QuoteTax() QuoteTaxResolver {
	return quoteTaxResolver{r}
}

func (r quoteTaxResolver) Tax(ctx context.Context, obj *db.QuoteTax) (*db.Tax, error) {
	return r.db.GetTax(obj.TaxId, nil)
}
