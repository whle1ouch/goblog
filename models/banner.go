package models

type BannerModel struct {
	Model
	Path string `json:"path"`
	Hash string `json:"hash"`
	Name string `gorm:"column:name;type:varchar(36)" json:"name"`
}
