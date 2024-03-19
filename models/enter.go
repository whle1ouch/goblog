package models

import "time"

type Model struct {
	ID       uint      `json:"id" gorm:"column:id;type:int(10);primarykey;unsigned;auto_increment;not null;comment:主键"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
