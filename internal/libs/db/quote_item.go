package db

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type QuoteItem struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	QuoteId     uuid.UUID
	ItemId      uuid.UUID
	ItemPriceId uuid.UUID
	Quantity    int
}

func (QuoteItem) TableName() string {
	return "quote_item"
}

func (d *DB) CreateQuoteItems(
	taintScanId uuid.UUID,
	quoteItems []*QuoteItem,
	txn *Txn,
) error {
	items :=
		lo.Map(quoteItems, func(c *QuoteItem, idx int) *QuoteItem {
			c.QuoteId = taintScanId

			return c
		})

	if err := d.getQuery(txn).Model(QuoteItem{}).Create(items).Error; err != nil {
		return err
	}

	return nil
}
