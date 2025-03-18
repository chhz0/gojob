// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTTaskSharding = "t_task_sharding"

// TTaskSharding 分片元信息表
type TTaskSharding struct {
	ShardID    int32     `gorm:"column:shard_id;primaryKey;comment:分片ID" json:"shard_id"`   // 分片ID
	TaskType   string    `gorm:"column:task_type;not null;comment:任务类型" json:"task_type"`   // 任务类型
	StartID    int64     `gorm:"column:start_id;not null;comment:任务分片起始ID" json:"start_id"` // 任务分片起始ID
	EndID      int64     `gorm:"column:end_id;not null;comment:任务分片结束ID" json:"end_id"`     // 任务分片结束ID
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName TTaskSharding's table name
func (*TTaskSharding) TableName() string {
	return TableNameTTaskSharding
}
