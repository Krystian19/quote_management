package external_bff

import (
	"context"
	"fmt"

	"github.com/Krystian19/quote_management/internal/libs/db"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func (r queryResolver) GetQuote(ctx context.Context, id uuid.UUID) (*db.Quote, error) {
	return r.db.GetQuote(id, nil)
}

func (r queryResolver) GetQuotes(ctx context.Context) ([]*db.Quote, error) {
	return r.db.GetAllQuotes(nil)
}

func (r mutationResolver) CreateQuote(ctx context.Context, fields CreateQuoteInput) (*db.Quote, error) {
	var res *db.Quote

	err := r.db.GetConn().Transaction(func(tx *gorm.DB) error {
		createdQuote, err := r.db.CreateQuote(db.Quote{
			AccountId: fields.AccountID,
		}, tx)
		if err != nil {
			return err
		}

		var quoteItemsToCreate []*db.QuoteItem
		for _, it := range fields.Items {
			foundItem, err := r.db.GetInventoryItem(it.ItemID, tx)
			if err != nil {
				return err
			}

			if foundItem == nil {
				return fmt.Errorf("no inventory item found with id %v", it.ItemID.String())
			}

			foundPrice, err := r.db.GetLatestInventoryItemPrice(foundItem.ID, tx)
			if err != nil {
				return err
			}

			quoteItemsToCreate = append(quoteItemsToCreate, &db.QuoteItem{
				QuoteId:     createdQuote.ID,
				ItemId:      it.ItemID,
				ItemPriceId: foundPrice.ID,
				Quantity:    it.Quantity,
			})
		}

		if err := r.db.CreateQuoteItems(quoteItemsToCreate, tx); err != nil {
			return err
		}

		foundTaxes, err := r.db.GetAllTaxes(tx)
		if err != nil {
			return err
		}

		if err := r.db.CreateQuoteTaxes(
			lo.Map(foundTaxes, func(t *db.Tax, idx int) *db.QuoteTax {
				return &db.QuoteTax{
					QuoteId: createdQuote.ID,
					TaxId:   t.ID,
				}
			}), tx,
		); err != nil {
			return err
		}

		res = createdQuote
		return nil
	})

	return res, err
}

type quoteResolver struct{ resolver }

func (r resolver) Quote() QuoteResolver { return quoteResolver{r} }

func (r quoteResolver) CreatedAt(ctx context.Context, obj *db.Quote) (string, error) {
	return obj.CreatedAt.String(), nil
}

func (r quoteResolver) UpdatedAt(ctx context.Context, obj *db.Quote) (string, error) {
	return obj.UpdatedAt.String(), nil
}

func (r quoteResolver) Items(ctx context.Context, obj *db.Quote) ([]*db.QuoteItem, error) {
	return r.db.GetQuoteItems(obj.ID, nil)
}

func (r quoteResolver) Taxes(ctx context.Context, obj *db.Quote) ([]*db.QuoteTax, error) {
	return r.db.GetQuoteTaxes(obj.ID, nil)
}

func (r quoteResolver) Conflicts(ctx context.Context, obj *db.Quote) ([]*QuoteConflict, error) {
	var res []*QuoteConflict

	foundQuoteItems, err := r.db.GetQuoteItems(obj.ID, nil)
	if err != nil {
		return res, err
	}

	for _, qit := range foundQuoteItems {
		foundLatestPrice, err := r.db.GetLatestInventoryItemPrice(qit.ItemId, nil)
		if err != nil {
			return []*QuoteConflict{}, err
		}

		if foundLatestPrice.ID != qit.ItemPriceId {
			res = append(res, &QuoteConflict{
				ItemID: qit.ItemId,
				Reason: QuoteConflictTypePriceChanged,
			})
		}
	}

	return res, nil
}

func (r quoteResolver) Total(ctx context.Context, obj *db.Quote) (float64, error) {
	var total float64

	foundQuoteItems, err := r.db.GetQuoteItems(obj.ID, nil)
	if err != nil {
		return 0, err
	}

	for _, qit := range foundQuoteItems {
		foundPrice, err := r.db.GetInventoryItemPrice(qit.ItemPriceId, nil)
		if err != nil {
			return 0, err
		}

		total += foundPrice.Price * float64(qit.Quantity)
	}

	return total, nil
}
