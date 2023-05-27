package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	MSISDN   string    `gorm:"type:varchar(32);unique;not null"`
	Name     string    `gorm:"type:varchar(255);not null"`
	Username string    `gorm:"type:varchar(255);unique;not null"`
	Password string    `gorm:"type:varchar(255);not null"`
}
