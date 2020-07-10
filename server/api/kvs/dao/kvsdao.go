package dao

import (
	"fmt"
	"net/http"
	"strings"

	"kvs-go/server/api/kvs/model"
	"kvs-go/server/api/utils"
)

func (dao *dao) GetValueByKey(key string) (*model.KeyValue, *utils.ApiError) {
	var keyvalue model.KeyValue
	if err := dao.db.Where(&model.KeyValue{Key: key}).Find(&keyvalue).Error; err != nil {
		rawtype := http.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			rawtype = http.StatusNotFound
		}
		return nil, &utils.ApiError{
			Error: err,
			Type:  rawtype,
		}
	}
	return &keyvalue, nil
}

func (dao *dao) SetKeyValue(key, value string) *utils.ApiError {
	keyvalue := &model.KeyValue{
		Key:   key,
		Value: value,
	}

	if err := dao.db.Create(keyvalue).Error; err != nil {
		return &utils.ApiError{
			Error: fmt.Errorf("couldn't create keyvalue: %s", err.Error()),
			Type:  http.StatusInternalServerError,
		}
	}

	return nil
}
