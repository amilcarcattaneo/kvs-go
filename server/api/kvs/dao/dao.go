package dao

import (
	"github.com/jinzhu/gorm"

	"kvs-go/server/api/kvs/model"
	"kvs-go/server/api/utils"
)

type Dao interface {
	GetValueByKey(key string) (*model.KeyValue, *utils.ApiError)
	SetKeyValue(key, value string) *utils.ApiError
}

type dao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) Dao {
	return &dao{
		db: db,
	}
}
