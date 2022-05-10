package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Requests struct {
	gorm.Model
	Uri      string
	Response datatypes.JSON
}
