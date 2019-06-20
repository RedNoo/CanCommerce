package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User Model Description
type User struct {
	ID         int64 `gorm:"primary_key"`
	FullName   string
	Email      string
	Password   string
	UserType   int        //1 admin 2 user
	UserStatus UserStatus `gorm:"foreignkey:UserStatusRefer"`

	CreatedAt time.Time
	DeletedAt *time.Time
	UpdatedAt time.Time
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m User) TableName() string {
	return "users"
}

// UserDB is the implementation of the storage interface for
// User.
type UserDB struct {
	Db *gorm.DB
}

// DB returns the underlying database.
func (m *UserDB) DB() interface{} {
	return m.Db
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *UserDB) TableName() string {
	return "users"

}

// Get returns a single User as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *UserDB) Get(id int64) (*User, error) {
	var native User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of User
func (m *UserDB) List() ([]*User, error) {

	var users []*User
	err := m.Db.Table(m.TableName()).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return users, nil
}

// Add creates a new record.
func (m *UserDB) Add(model *User) error {

	err := m.Db.Create(model).Error
	if err != nil {
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *UserDB) Update(model *User) error {

	user, err := m.Get(model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(user).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *UserDB) Delete(id int) error {

	var user User

	err := m.Db.Delete(&user, id).Error

	if err != nil {
		return err
	}

	return nil
}
