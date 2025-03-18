package apiv1

import "time"

// TaskScheduleCfg 任务调度信息
type TaskScheduleCfg struct {
	TaskType          string     `json:"task_type"`
	ScheduleLimit     int        `json:"schedule_limit"`
	ScheduleInterval  int        `json:"schedule_interval"`
	MaxProcessingTime int64      `json:"max_processing_time"`
	MaxRetryNum       int        `json:"max_retry_num"`
	RetryInterval     int        `json:"retry_interval"`
	MaxRetryInterval  int        `json:"max_retry_interval"`
	CreateTime        *time.Time `json:"create_time"`
	ModifyTime        *time.Time `json:"modify_time"`
}

// TaskData 任务调度数据
type TaskData struct {
	UserId           string    `json:"user_id"`
	TaskId           string    `json:"task_id"`
	TaskType         string    `json:"task_type"`
	TaskStage        string    `json:"task_stage"`
	Status           int       `json:"status"`
	Priority         *int      `json:"priority"`
	CrtRetryNum      int       `json:"crt_retry_num"`
	MaxRetryNum      int       `json:"max_retry_num"`
	MaxRetryInterval int       `json:"max_retry_interval"`
	ScheduleLog      string    `json:"schedule_log"`
	TaskContext      string    `json:"context"`
	OrderTime        int64     `json:"order_time"`
	CreateTime       time.Time `json:"create_time"`
	ModifyTime       time.Time `json:"modify_time"`
}
