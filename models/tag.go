package models

type TagModel struct {
	Model
	Title    string         `gorm:"column:title;type:varchar(16)" json:"title"`
	Articles []ArticleModel `gorm:"many2many:article_tag;" json:"-"`
}
