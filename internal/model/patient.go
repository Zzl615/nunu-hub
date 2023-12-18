package model

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
}

// func (m *Patient) TableName() string {
// 	return "patient"
// }
