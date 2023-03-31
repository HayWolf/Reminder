package store

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID          uint           `gorm:"column:id;primaryKey;autoIncrement"`
	TaskId      uint           `gorm:"column:task_id;not null"`
	ToUserId    uint           `gorm:"column:to_user_id;not null"`
	TriggerTime time.Time      `gorm:"column:trigger_time;not null;index"`
	Content     string         `gorm:"column:content;type:string;"`
	Status      int            `gorm:"column:status;not null;default:0;index;comment:0待发送,1发送中,2已发送,3发送失败"`
	Retry       uint           `gorm:"column:retry;not null;default:0""`
	CreateTime  time.Time      `gorm:"column:create_time;autoCreateTime:nano"`
	ModifyTime  time.Time      `gorm:"column:modify_time;autoUpdateTime:nano"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Message) TableName() string {
	return "t_message"
}
