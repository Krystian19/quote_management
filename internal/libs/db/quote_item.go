package db

import (
	"github.com/google/uuid"
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
	items []*QuoteItem,
	txn *Txn,
) error {
	return d.getQuery(txn).Model(QuoteItem{}).Create(items).Error
}

func (d *DB) GetQuoteItems(
	quoteId uuid.UUID,
	txn *Txn,
) ([]*QuoteItem, error) {
	var res []*QuoteItem
	err := d.getQuery(txn).Where("quote_id = ?", quoteId).Find(&res).Error

	return res, err
}
