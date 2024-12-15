package db

import (
	"github.com/google/uuid"
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
	quoteTaxes []*QuoteTax,
	txn *Txn,
) error {
	return d.getQuery(txn).Model(QuoteTax{}).Create(quoteTaxes).Error
}

func (d *DB) GetQuoteTaxes(
	quoteId uuid.UUID,
	txn *Txn,
) ([]*QuoteTax, error) {
	var res []*QuoteTax
	err := d.getQuery(txn).Where("quote_id = ?", quoteId).Find(&res).Error

	return res, err
}
