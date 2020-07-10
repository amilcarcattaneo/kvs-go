package model

type KeyValue struct {
	Key   string `gorm:"column:rawkey" json:"key"`
	Value string `gorm:"column:rawvalue" json:"value"`
}

func (KeyValue) TableName() string {
	return "keyvalue"
}
