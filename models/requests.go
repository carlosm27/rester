package models

import (
	"gorm.io/gorm"
)

type Requests struct {
	gorm.Model
	Uri      string
	Response string
}
