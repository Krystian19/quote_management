package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	db *gorm.DB
}

type Opts struct {
	Url string
}

// Represents a database transaction
type Txn = gorm.DB

func New(opts Opts) (*DB, error) {
	// NOTE we have to make sure that all the code that is touching the db
	// is using UTC time
	time.Local = time.UTC

	db, err := gorm.Open(postgres.Open(opts.Url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}

func (d *DB) getQuery(txn *Txn) *gorm.DB {
	if txn != nil {
		return txn
	}

	return d.db
}

func (d *DB) GetConn() *gorm.DB {
	return d.db
}
