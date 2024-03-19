package models

import "goblog/models/ctype"

type MenuModel struct {
	Model
	MenuTitle    string        `gorm:"column:menu_title;type:varchar(32)" json:"menu_title"`
	MenuTitleEn  string        `gorm:"column:menu_title_en;type:varchar(32)" json:"menu_title_en"`
	Slogan       string        `gorm:"column:slogan;type:varchar(64)" json:"slogan"`
	Abstract     ctype.Array   `gorm:"type:string;" json:"abstract"`
	AbstractTime int           `json:"abstract_time"`
	MenuBanner   []BannerModel `gorm:"many2many:menu_image;joinForeignKey:MenuID;joinReferences:BannerID;" json:"menu_images"`
	MenuTime     int           `json:"menu_time"`
	Sort         int           `gorm:"size:10" json:"sort"`
}

type MenuBannerModel struct {
	MenuID      uint        `json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID" json:"menu_model"`
	BannerID    uint        `json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID" json:"banner_model"`
	Sort        int         `gorm:"size:10" json:"sort"`
}
