package models

type MessageModel struct {
	Model
	SenderID       uint      `gorm:"primaryKey;comment:发送者ID" json:"sender_id"`
	SenderUser     UserModel `gorm:"foreignKey:SenderID;comment:发送者" json:"-"`
	SenderNickName string    `gorm:"size:36;comment:发送者昵称" json:"sender_nick_name"`
	SenderAvatar   string    `json:"sender_avatar"`

	ReceiverID       uint      `gorm:"primaryKey;comment:接收者ID" json:"receiver_id"`
	ReceiverUser     UserModel `gorm:"foreignKey:ReceiverID;comment:接收者" json:"-"`
	ReceiverNickName string    `gorm:"size:36;comment:接收者昵称" json:"receiver_nick_name"`
	ReceiverAvatar   string    `json:"receiver_avatar"`
	IsRead           bool      `gorm:"default:false;comment:已读" json:"is_read"`
	Content          string    `gorm:"size:256;not null;comment:内容;comment:内容" json:"content"`
}
