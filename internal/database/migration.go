package database

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"gopkg.in/gormigrate.v1"
)

// MigrateDB executes the migrations.
func MigrateDB(db *gorm.DB) error {
    
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
        {
            ID: "20210528100",
            Migrate: func(tx *gorm.DB) error {
                return tx.AutoMigrate(&User{}).Error
            },
            Rollback: func(tx *gorm.DB) error {
                return tx.DropTable("users").Error
            },
        },
    })

	if err := m.Migrate(); err != nil {
		return errors.Wrap(err, "database: migration failed")
	}

	return nil
}
