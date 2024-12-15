package db

import (
	"errors"
	"time"

	"github.com/bluele/factory-go/factory"
	"github.com/google/uuid"
	"github.com/jaswdr/faker"
	"gorm.io/gorm"
)

type InventoryItem struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (InventoryItem) TableName() string {
	return "inventory_item"
}

var InventoryItemFactory = factory.NewFactory(
	InventoryItem{},
).
	Attr("Name", func(args factory.Args) (interface{}, error) {
		return faker.New().Address().City(), nil
	})

func (d *DB) CreateInventoryItem(newItem InventoryItem, txn *Txn) (*InventoryItem, error) {
	if err := d.getQuery(txn).Create(&newItem).Error; err != nil {
		return nil, err
	}

	return &newItem, nil
}

func (d *DB) GetInventoryItem(itemId uuid.UUID, txn *Txn) (*InventoryItem, error) {
	var res InventoryItem

	if err := d.getQuery(txn).Where("id = ?", itemId).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}

func (d *DB) GetAllInventoryItems(itemId uuid.UUID, txn *Txn) ([]*InventoryItem, error) {
	var res []*InventoryItem
	err := d.getQuery(txn).Find(&res).Error

	return res, err
}
