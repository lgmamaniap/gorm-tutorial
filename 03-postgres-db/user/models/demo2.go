package models

type Demo2 struct {
	ID    uint64 `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"size:255"`
	Data  string `json:"data" gorm:"size:255"`
	Demo1 *Demo1 `json:"demo1" gorm:"foreignKey:ID"`
}
