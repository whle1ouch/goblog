package models

import "goblog/models/ctype"

type MenuModel struct {
	Model
	MenuTitle    string        `gorm:"column:menu_title;type:varchar(32);comment:菜单名" json:"menu_title"`
	MenuTitleEn  string        `gorm:"column:menu_title_en;type:varchar(32);comment:菜单名英文" json:"menu_title_en"`
	Slogan       string        `gorm:"column:slogan;type:varchar(64);comment:标语" json:"slogan"`
	Abstract     ctype.Array   `gorm:"type:string;comment:简介" json:"abstract"`
	AbstractTime int           `json:"abstract_time"`
	MenuBanner   []BannerModel `gorm:"many2many:menu_image;joinForeignKey:MenuID;joinReferences:BannerID;comment:菜单图片" json:"menu_images"`
	MenuTime     int           `json:"menu_time"`
	Sort         int           `gorm:"size:10;comment:排序值" json:"sort"`
}

type MenuBannerModel struct {
	MenuID      uint        `json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID;comment:关联菜单" json:"menu_model"`
	BannerID    uint        `json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID;comment:关联图片" json:"banner_model"`
	Sort        int         `gorm:"size:10;comment:排序" json:"sort"`
}
