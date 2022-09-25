package model

import "github.com/jinzhu/gorm"

type Save struct {
	Request  string
	Response string
	gorm.Model
}
