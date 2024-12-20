package db

import (
	"errors"
	"time"

	"github.com/bluele/factory-go/factory"
	"github.com/google/uuid"
	"github.com/jaswdr/faker"
	"gorm.io/gorm"
)

type Tax struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string
	TaxRate     float64
	EffectiveAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Tax) TableName() string {
	return "tax"
}

var TaxFactory = factory.NewFactory(
	Tax{},
).
	Attr("Name", func(args factory.Args) (interface{}, error) {
		return faker.New().Address().City(), nil
	}).
	Attr("TaxRate", func(args factory.Args) (interface{}, error) {
		return float64(1.0), nil
	}).
	Attr("EffectiveAt", func(args factory.Args) (interface{}, error) {
		return time.Now(), nil
	})

func (d *DB) CreateTax(newTax Tax, txn *Txn) (*Tax, error) {
	if err := d.getQuery(txn).Create(&newTax).Error; err != nil {
		return nil, err
	}

	return &newTax, nil
}

func (d *DB) GetTax(id uuid.UUID, txn *Txn) (*Tax, error) {
	var res *Tax

	if err := d.getQuery(txn).Where("id = ?", id).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return res, nil
}

func (d *DB) CreateTaxes(newTaxes []*Tax, txn *Txn) error {
	return d.getQuery(txn).Create(newTaxes).Error
}

func (d *DB) GetAllTaxes(txn *Txn) ([]*Tax, error) {
	var res []*Tax
	err := d.getQuery(txn).Where("effective_at < ?", time.Now()).Find(&res).Error

	return res, err
}
