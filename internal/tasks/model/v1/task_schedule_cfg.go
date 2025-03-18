package modelv1

import "time"

// TaskScheduleCfg 任务调度配置表
type TaskScheduleCfg struct {
	TaskType         string
	ScheduleLimit    int
	ScheduleInterval int
	MaxProcessTime   int64
	MaxRetryNum      int
	RetryInterval    int
	MaxRetryInterval int

	CreateTime *time.Time
	UpdateTime *time.Time
}

func (tsc *TaskScheduleCfg) TableName() string {
	return "t_task_schedule_cfg"
}
