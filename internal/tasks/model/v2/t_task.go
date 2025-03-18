package model

import "fmt"

const taskTableNameFormat = "t_task_type_%s_%s"

// TTaskType0000 任务分片表
type TTask struct {
	ID              int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:任务ID - 自增id" json:"id"`              // 任务ID - 自增id
	TaskID          string `gorm:"column:task_id;not null;comment:任务唯一ID(UUID|snowflake|MD5)" json:"task_id"`          // 任务唯一ID(UUID|snowflake|MD5)
	TaskType        string `gorm:"column:task_type;not null;comment:任务类型" json:"task_type"`                            // 任务类型
	ExecutorHandler string `gorm:"column:executor_handler;not null;comment:执行器处理器" json:"executor_handler"`            // 执行器处理器
	ShardingParam   string `gorm:"column:sharding_param;comment:分片参数" json:"sharding_param"`                           // 分片参数
	ScheduleType    string `gorm:"column:schedule_type;not null;comment:调度类型 CRON/FIX_RATE" json:"schedule_type"`      // 调度类型 CRON/FIX_RATE
	ScheduleConf    string `gorm:"column:schedule_conf;not null;comment:调度配置" json:"schedule_conf"`                    // 调度配置
	TaskContext     string `gorm:"column:task_context;comment:任务上下文(JSON格式)" json:"task_context"`                      // 任务上下文(JSON格式)
	RetryCount      int32  `gorm:"column:retry_count;not null;comment:已重试次数" json:"retry_count"`                       // 已重试次数
	TriggerStatus   int32  `gorm:"column:trigger_status;not null;comment:触发状态 0-STOP 1-RUNNING" json:"trigger_status"` // 触发状态 0-STOP 1-RUNNING
	TriggerLastTime int64  `gorm:"column:trigger_last_time;comment:上次触发时间戳" json:"trigger_last_time"`                  // 上次触发时间戳
	TriggerNextTime int64  `gorm:"column:trigger_next_time;comment:下次触发时间戳" json:"trigger_next_time"`                  // 下次触发时间戳
	CreateTime      int64  `gorm:"column:create_time;not null;comment:创建时间戳" json:"create_time"`                       // 创建时间戳
	UpdateTime      int64  `gorm:"column:update_time;not null;comment:更新时间戳" json:"update_time"`                       // 更新时间戳
}

// TableName TTask's table name
func (*TTask) TableName() string {
	return "t_task"
}

func (tt *TTask) GetTableName(shardID string) string {
	return fmt.Sprintf(taskTableNameFormat, tt.TaskType, shardID)
}
