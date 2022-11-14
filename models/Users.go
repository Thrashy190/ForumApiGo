package models

import "gorm.io/gorm"

type Users struct {
    gorm.Model
    Username string  `gorm:"not null; uniqueIndex"`
    Password string	 `gorm:"not null"`
}
