package db

import (
	"time"

	"github.com/bluele/factory-go/factory"
	"github.com/google/uuid"
)

type InventoryItemSupply struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	InventoryItemId uuid.UUID
	Quantity        int
	CreatedAt       time.Time
}

func (InventoryItemSupply) TableName() string {
	return "inventory_item_supply"
}

var InventoryItemSupplyFactory = factory.NewFactory(
	Quote{},
).
	Attr("InventoryItemId", func(args factory.Args) (interface{}, error) {
		return uuid.New(), nil
	}).
	Attr("Quantity", func(args factory.Args) (interface{}, error) {
		return 7, nil
	})

func (d *DB) CreateInventoryItemSupply(newItemSupply InventoryItemSupply, txn *Txn) (*InventoryItemSupply, error) {
	if err := d.getQuery(txn).Create(&newItemSupply).Error; err != nil {
		return nil, err
	}

	return &newItemSupply, nil
}

func (d *DB) GetAllInventoryItemSupply(txn *Txn) ([]*InventoryItemSupply, error) {
	var res []*InventoryItemSupply
	err := d.getQuery(txn).Find(&res).Error

	return res, err
}
