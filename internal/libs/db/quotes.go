package db

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Quote struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()"`
	AccountId uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()"`
	PaymentId *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d *DB) CreateQuote(newQuote Quote, txn *Txn) (*Quote, error) {
	if err := d.getQuery(txn).Create(&newQuote).Error; err != nil {
		return nil, err
	}

	return &newQuote, nil
}

func (d *DB) GetQuote(quoteId uuid.UUID, txn *Txn) (*Quote, error) {
	var res Quote

	if err := d.getQuery(txn).Where("id = ?", quoteId).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}

func (d *DB) GetAllQuotes(txn *Txn) ([]*Quote, error) {
	var res []*Quote
	err := d.getQuery(txn).Find(&res).Error

	return res, err
}
