package models

import "goblog/models/ctype"

type ArticleModel struct {
	Model
	Title           string         `gorm:"column:title;type:varchar(32);not null;comment:标题" json:"title"`
	Abstract        string         `gorm:"column:abstract;type:varchar(128);not null;comment:摘要" json:"abstract"`
	Content         string         `gorm:"column:content;type:text;not null;comment:内容" json:"content"`
	LookCount       int            `gorm:"column:look_count;type:int;not null;default:0;comment:浏览次数" json:"look_count"`
	DiggCount       int            `gorm:"column:digg_count;type:int;not null;default:0;comment:点赞次数" json:"digg_count"`
	CollectionCount int            `gorm:"column:collection_count;type:int;not null;default:0;comment:收藏次数" json:"collection_count"`
	TagModels       []TagModel     `gorm:"many2many:article_tag;" json:"tag_models"`
	CommentModels   []CommentModel `gorm:"foreignkey:ArticleID;" json:"-"`
	UserModel       UserModel      `gorm:"foreignkey:UserID;" json:"-"`
	UserID          uint           `json:"user_id"`
	Catagory        string         `gorm:"column:Category;type:varchar(32);not null;comment:分类" json:"category"`
	WordCount       int            `json:"word_count"`
	BannerModel     BannerModel    `gorm:"foreignKey:BannerID" json:"-"`
	BannerID        uint           `json:"banner_id"`
	BannerPath      string         `json:"banner_path"`
	NickName        string         `gorm:"column:nick_name;type:varchar(42);comment:用户昵称" json:"nick_name"`
	Tags            ctype.Array    `gorm:"type:string;size:64"`
}
