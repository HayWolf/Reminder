package store

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID         uint           `gorm:"column:id;primaryKey;autoIncrement"`
	UserId     uint           `gorm:"column:user_id;not null"`
	Nickname   string         `gorm:"column:nickname;default:"`
	Event      string         `gorm:"column:event;type:string"`
	Unit       string         `gorm:"colum:unit;type:string;not null;default:none"`
	Interval   uint           `gorm:"colum:interval;not null;default:0"`
	FirstTime  time.Time      `gorm:"colum:first_time"`
	LastTime   time.Time      `gorm:"colum:last_time;comment:最近一次派单时间"`
	CreateTime time.Time      `gorm:"column:create_time;autoCreateTime:nano"`
	ModifyTime time.Time      `gorm:"column:modify_time;autoUpdateTime:nano"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Task) TableName() string {
	return "t_task"
}
