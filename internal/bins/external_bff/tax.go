package external_bff

import (
	"context"
	"time"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/samber/lo"
)

func (r queryResolver) GetAllTaxes(ctx context.Context) ([]*db.Tax, error) {
	return r.db.GetAllTaxes(nil)
}

func (r mutationResolver) CreateTax(ctx context.Context, fields CreateTaxInput) (*db.Tax, error) {
	return r.db.CreateTax(db.Tax{
		Name:        fields.Name,
		TaxRate:     fields.TaxRate,
		EffectiveAt: time.Now(),
	}, nil)
}

func (r mutationResolver) CreateTaxes(ctx context.Context, fields []*CreateTaxInput) (bool, error) {
	taxesToCreate := lo.Map(fields, func(f *CreateTaxInput, idx int) *db.Tax {
		return &db.Tax{
			Name:        f.Name,
			TaxRate:     f.TaxRate,
			EffectiveAt: time.Now(),
		}
	})

	err := r.db.CreateTaxes(taxesToCreate, nil)
	return true, err
}

type taxResolver struct{ resolver }

func (r resolver) Tax() TaxResolver { return taxResolver{r} }

func (r taxResolver) CreatedAt(ctx context.Context, obj *db.Tax) (string, error) {
	return obj.CreatedAt.String(), nil
}

func (r taxResolver) UpdatedAt(ctx context.Context, obj *db.Tax) (string, error) {
	return obj.UpdatedAt.String(), nil
}

func (r taxResolver) EffectiveAt(ctx context.Context, obj *db.Tax) (string, error) {
	return obj.EffectiveAt.String(), nil
}
