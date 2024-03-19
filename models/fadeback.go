package models

type FadeBackModel struct {
	Model
	Email        string `gorm:"size:64;comment:邮箱" json:"email"`
	Content      string `gorm:"size:256;not null;comment:反馈内容" json:"content"`
	ApplyContent string `gorm:"size:256;not null;comment:回复内容" json:"apply_content"`
	IsApply      bool   `json:"is_apply"`
}
