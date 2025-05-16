package models

type Album struct {
	ID     uint    `json:"id" gorm:"primarykey;autoIncrement"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Year   uint    `json:"year"`
	Price  float64 `json:"price"`
}
