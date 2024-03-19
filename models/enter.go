package models

import "time"

type Model struct {
	ID       uint      `json:"id" gorm:"column:id;type:int(10);primarykey;unsigned;auto_increment;not null;comment:主键"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at;type:datetime;not null;comment:创建时间"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at;type:datetime;not null;comment:更新时间"`
}
