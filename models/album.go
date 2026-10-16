package models

type Album struct {
	ID     uint    `json:"id" gorm:"primarykey;autoIncrement"`
	Title  string  `json:"title" gorm:"not null"`
	Artist string  `json:"artist" gorm:"not null"`
	Price  float64 `jaon:"price" gorm:"type:decimal(10,2)"`
}
