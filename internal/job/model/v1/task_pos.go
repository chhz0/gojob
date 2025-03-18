package modelv1

import "time"

// TaskPos 任务位置表
type TaskPos struct {
	Id               uint64
	TaskType         string
	ScheduleBeginPos int
	ScheduleEndPos   int

	CreateTime *time.Time
	UpdateTime *time.Time
}

func (tp *TaskPos) TableName() string {
	return "t_task_pos"
}
