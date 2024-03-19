package models

import (
	"goblog/models/ctype"
	"time"
)

type UserModel struct {
	Model
	NickName     string           `gorm:"column:nick_name;type:varchar(36);not null;comment:昵称" json:"nick_name"`
	UserName     string           `gorm:"column:user_name;type:varchar(36);unique;not null;comment:用户名" json:"user_name"`
	PassWord     string           `gorm:"column:pass_word;type:varchar(128);not null;comment:密码" json:"pass_word"`
	Avatar       string           `gorm:"column:avatar_id;type:varchar(256);comment:头像地址" json:"avatar_id"`
	Email        string           `gorm:"column:email;type:varchar(20);comment:邮箱" json:"email"`
	Tel          string           `gorm:"column:tel;type:varchar(18);comment:电话" json:"tel"`
	Addr         string           `gorm:"column:addr;type:varchar(64);comment:地址" json:"addr"`
	Token        string           `gorm:"column:token;type:varchar(64);comment:token" json:"token"`
	IP           string           `gorm:"column:ip;type:varchar(20);comment:ip地址" json:"ip"`
	Role         ctype.RoleType   `gorm:"column:role;type:int(4);default:1;comment:角色ID"`
	SignStatus   ctype.SignStatus `gorm:"column:sign_status;type:smallint;comment:登录方式"`
	ArticleModel []ArticleModel   `gorm:"foreignKey:UserID;comment:文章" json:"-"`
	Collections  []ArticleModel   `gorm:"many2many:user_collection;joinForeignKey:UserID;JoinReferences:ArticleID;comment:收藏文章" json:"-"`
}

type UserCollectionModel struct {
	UserID       uint         `gorm:"primaryKey;comment:关联用户ID"`
	UserModel    UserModel    `gorm:"foreignKey:UserID;comment:关联用户"`
	ArticleID    uint         `gorm:"primaryKey;comment:关联文章ID"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID;comment:关联文章"`
	CreateAt     time.Time
}
