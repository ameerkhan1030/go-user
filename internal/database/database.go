package database

import (
	"fmt"
	"test/pkg/config"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// DataStore represents a database handle for persistent storage.
type DataStore struct {
	DB *gorm.DB
}

// New returns a new instance of the datastore.
func New(c config.Config) (*DataStore, error) {

	cfg := mysql.Config{
		User:                 c.DatabaseUser,
		Passwd:               c.DatabasePassword,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", c.DatabaseHost, c.DatabasePort),
		DBName:               c.DatabaseName,
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	db, err := gorm.Open("mysql",cfg.FormatDSN())
	if err != nil {
		panic("failed to connect database")
	}

	return &DataStore{DB: db}, nil
}

// NewTransaction returns a new transaction.
func (s DataStore) NewTransaction() (*gorm.DB, error) {
	tx := s.DB.Begin()
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "store: begin transaction failed")
	}

	return tx, nil
}

// Commit commits all the changes executed in passed transaction.
func (s DataStore) Commit(tx *gorm.DB) error {
	if tx == nil {
		return errors.New("store: cannot commit empty transaction")
	}

	if err := tx.Commit().Error; err != nil {
		return errors.Wrap(err, "store: commit failed")
	}

	return nil
}

// Rollback rollbacks all the changes executed in passed transaction.
func (s DataStore) Rollback(tx *gorm.DB) error {
	if tx == nil {
		return errors.New("store: cannot rollback empty transaction")
	}

	if err := tx.Rollback().Error; err != nil {
		return errors.Wrap(err, "store: rollback failed")
	}

	return nil
}