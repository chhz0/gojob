package modelv1

import (
	"fmt"
	"time"
)

type Task struct {
	Id               int    `gorm:"primary_key;autoIncrement:15"`
	UserId           string `gorm:"type:varchar(256);column:user_id;not null;default:'';index:idx_user_id"`
	TaskId           string `gorm:"type:varchar(256);column:task_id;not null;default:'';uniqueIndex:idx_task_id"`
	TaskType         string `gorm:"type:varchar(128);column:task_type;not null;default:''"`
	TaskStage        string `gorm:"type:varchar(128);column:task_stage;not null;default:''"`
	Status           int    `gorm:"type:tinyint(3);column:status;not null;default:1;index:idx_status_order_time;index:idx_status"`
	Priority         int    `gorm:"type:int(11);column:priority;not null;default:0;comment:'优先级'"`
	CrtRetryNum      int    `gorm:"type:int(11);column:crt_retry_num;not null;default:0;comment:'重试次数'"`
	MaxRetryNum      int    `gorm:"type:int(11);column:max_retry_num;not null;default:0;comment:'最大重试次数'"`
	MaxRetryInterval int    `gorm:"type:int(11);column:max_retry_interval;not null;default:0;comment:'最大重试间隔'"`
	ScheduleLog      string `gorm:"type:varchar(4096);column:schedule_log;not null;default:'';comment:'调度信息表'"`
	TaskContext      string `gorm:"type:varchar(8192);column:task_context;not null;default:'';comment:'任务上下文，用户自定义'"`
	OrderTime        int64  `gorm:"type:int(20);column:order_time;not null;default:0;comment:'调度时间';index:idx_status_order_time'"`

	CreateTime time.Time `gorm:"type:datetime;column:create_time;not null;autoCreateTime:milli"`
	UpdateTime time.Time `gorm:"type:datetime;column:update_time;not null;autoUpdateTime:milli"`
}

// 以任务类型和任务位置生成表名
func getTableName(taskType, pos string) string {
	return fmt.Sprintf("t_%s_task_%s", taskType, pos)
}

func (t *Task) TableName() string {
	return "task"
}
