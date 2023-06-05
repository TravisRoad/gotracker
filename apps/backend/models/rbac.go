package models

import "gorm.io/gorm"

type Policy struct {
	gorm.Model
	Sub string
	Obj string
	Act string // Allow or Deny
}
