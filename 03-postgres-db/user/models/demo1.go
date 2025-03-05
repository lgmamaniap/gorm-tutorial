package models

type Demo1 struct {
	ID    uint64   `json:"id" gorm:"primaryKey"`
	Name  string   `json:"name" gorm:"size:255"`
	Data  string   `json:"data" gorm:"size:255"`
	Demo2 *[]Demo2 `json:"demo2" gorm:"foreignKey:ID"`
}
