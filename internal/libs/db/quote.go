package db

import (
	"errors"
	"time"

	"github.com/Krystian19/quote_management/internal/libs/utils"
	"github.com/bluele/factory-go/factory"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Quote struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	AccountId uuid.UUID
	PaymentId *uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Quote) TableName() string {
	return "quote"
}

var QuoteFactory = factory.NewFactory(
	Quote{},
).
	Attr("AccountId", func(args factory.Args) (interface{}, error) {
		return uuid.New(), nil
	}).
	Attr("PaymentId", func(args factory.Args) (interface{}, error) {
		return utils.GetPtr(uuid.New()), nil
	})

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
