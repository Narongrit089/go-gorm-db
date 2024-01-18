package models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model //gorm create ID, CreatedAt,UpdatedAt, DeletedAt,
	Code       string
	Name       string
}
