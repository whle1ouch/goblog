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
	Email        string           `gorm:"column:email;type:varchar(20)" json:"email"`
	Tel          string           `gorm:"column:tel;type:varchar(18)" json:"tel"`
	Addr         string           `gorm:"column:addr;type:varchar(64)" json:"addr"`
	Token        string           `gorm:"column:token;type:varchar(64)" json:"token"`
	IP           string           `gorm:"column:ip;type:varchar(20)" json:"ip"`
	Role         ctype.RoleType   `gorm:"column:role;type:int(4);default:1;comment:角色"`
	SignStatus   ctype.SignStatus `gorm:"column:sign_status;type:smallint;comment:登录方式"`
	ArticleModel []ArticleModel   `gorm:"foreignKey:UserID" json:"-"`
	Collections  []ArticleModel   `gorm:"many2many:user_collection;joinForeignKey:UserID;JoinReferences:ArticleID" json:"-"`
}

type UserCollectionModel struct {
	UserID       uint         `gorm:"primaryKey"`
	UserModel    UserModel    `gorm:"foreignKey:UserID"`
	ArticleID    uint         `gorm:"primaryKey"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
	CreateAt     time.Time
}
