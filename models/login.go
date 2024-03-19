package models

type LoginDataModel struct {
	Model
	UserID    uint      `json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID;comment:登录用户" json:"user_model"`
	IP        string    `gorm:"size:20;comment:登录IP" json:"ip"`
	NickName  string    `gorm:"size:42;comment:用户昵称" json:"nick_name"`
	Token     string    `gorm:"size:256;comment:登录token" json:"token"`
	Device    string    `gorm:"size:256;comment:登录设备" json:"device"`
	Addr      string    `gorm:"size:64;comment:登录地址" json:"addr"`
}
