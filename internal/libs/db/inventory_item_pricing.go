package db

import (
	"errors"
	"time"

	"github.com/bluele/factory-go/factory"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InventoryItemPrice struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	InventoryItemId uuid.UUID
	Price           float64
	Version         uint
	CreatedAt       time.Time
}

func (InventoryItemPrice) TableName() string {
	return "inventory_item_price"
}

var InventoryItemPriceFactory = factory.NewFactory(
	InventoryItemPrice{},
).
	Attr("InventoryItemId", func(args factory.Args) (interface{}, error) {
		return uuid.New(), nil
	}).
	Attr("Price", func(args factory.Args) (interface{}, error) {
		return float64(1.0), nil
	}).
	Attr("Version", func(args factory.Args) (interface{}, error) {
		return uint(1), nil
	})

func (d *DB) CreateInventoryItemPrice(newItemPrice InventoryItemPrice, txn *Txn) (*InventoryItemPrice, error) {
	if err := d.getQuery(txn).Create(&newItemPrice).Error; err != nil {
		return nil, err
	}

	return &newItemPrice, nil
}

func (d *DB) GetLatestInventoryItemPrice(inventoryItemId uuid.UUID, txn *Txn) (*InventoryItemPrice, error) {
	var res InventoryItemPrice

	if err := d.getQuery(txn).Where("inventory_item_id = ?", inventoryItemId).Order("version DESC").First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}
