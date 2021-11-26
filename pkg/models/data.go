package models

type Data struct {
	Key   string `gorm:"not null;unique" json:"key"`
	Value string `json:"value"`
}
