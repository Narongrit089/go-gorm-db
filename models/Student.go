package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model //gorm create ID, CreatedAt,UpdatedAt, DeletedAt,
	Code       string
	Name       string
}
