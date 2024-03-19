package models

type MessageModel struct {
	Model
	SenderID       uint      `gorm:"primaryKey" json:"sender_id"`
	SenderUser     UserModel `gorm:"foreignKey:SenderID" json:"-"`
	SenderNickName string    `gorm:"size:36" json:"sender_nick_name"`
	SenderAvatar   string    `json:"sender_avatar"`

	ReceiverID       uint      `gorm:"primaryKey" json:"receiver_id"`
	ReceiverUser     UserModel `gorm:"foreignKey:ReceiverID" json:"-"`
	ReceiverNickName string    `gorm:"size:36" json:"receiver_nick_name"`
	ReceiverAvatar   string    `json:"receiver_avatar"`
	IsRead           bool      `gorm:"default:false" json:"is_read"`
	Content          string    `gorm:"size:256;not null;comment:内容" json:"content"`
}
