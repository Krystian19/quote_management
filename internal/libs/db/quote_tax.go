package db

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type QuoteTax struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	QuoteId uuid.UUID
	TaxId   uuid.UUID
}

func (QuoteTax) TableName() string {
	return "quote_tax"
}

func (d *DB) CreateQuoteTaxes(
	taintScanId uuid.UUID,
	quoteTaxes []*QuoteTax,
	txn *Txn,
) error {
	items :=
		lo.Map(quoteTaxes, func(c *QuoteTax, idx int) *QuoteTax {
			c.QuoteId = taintScanId

			return c
		})

	if err := d.getQuery(txn).Model(QuoteTax{}).Create(items).Error; err != nil {
		return err
	}

	return nil
}
