package controller

import (
	"github.com/jinzhu/gorm"

	"kvs-go/server/api/kvs/dao"
)

type Controller struct {
	dao dao.Dao
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		dao: dao.NewDao(db),
	}
}
