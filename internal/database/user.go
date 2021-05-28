package database

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// User struct is used to managing data for users
type User struct {
	gorm.Model
	Password      string
	Email         string `gorm:"size:256;not null"`
}


func AddUser(user *User,ds *DataStore) error {

	tx,err := ds.NewTransaction();
	defer func(ds *DataStore, tx *gorm.DB) {
		if err != nil && tx != nil {
			if txErr := ds.Rollback(tx); txErr != nil {
				err = errors.Wrap(err, txErr.Error())
			}
		}
	}(ds, tx)
	if err != nil {
		// log
	}
	tx.Create(user)
	tx.Commit();

	return err
}

func UserList(ds *DataStore) ([]User,error) {

	var users []User
	ds.DB.Find(&users)

	

	return users,nil
}