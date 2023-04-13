package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Title       string     `json:"title" form:"title" valid:"required~Title pf ypur product is required"`
	Description string     `json:"description" form:"description" valid:"required~Description of your product is required"`
	UserID      uint
	User        *User
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
