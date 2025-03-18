package apiv1

// CreateTaskReq 请求消息sdsadadsa
type CreateTaskReq struct {
	TaskData TaskData `json:"task_data"`
}

// GetTaskListReq 请求消息
type GetTaskListReq struct {
	TaskType string `json:"task_type" form:"task_type" binding:"required"`
	Status   int    `json:"status" form:"status" binding:"required"`
	Limit    int    `json:"limit" form:"limit" binding:"required"`
}

// HoldTasksReq 请求消息
type HoldTasksReq struct {
	TaskType string `json:"task_type" form:"task_type"`
	Limit    int    `json:"limit" form:"limit"`
}

// GetTaskCountsByTypeReq 请求消息
type GetTaskCountsByTypeReq struct {
	TaskType string `json:"task_type" form:"task_type"`
}

// GetTaskReq 请求消息
type GetTaskReq struct {
	TaskId string `json:"task_id" form:"task_id"`
}

// GetTaskCountByStatusReq 请求消息
type GetTaskCountByStatusReq struct {
	TaskType string `json:"task_type" form:"task_type"`
	Status   int    `json:"status" form:"status"`
}

// GetTaskScheduleCfgListReq 获取任务配置信息 请求体（空）
type GetTaskScheduleCfgListReq struct {
}

// RegisterTaskReq 任务注册接口 请求体
type RegisterTaskReq struct {
	TaskType string `json:"task_type" form:"task_type"`
}

// SetTaskReq 请求消息
type SetTaskReq struct {
	TaskData `json:"task_data"`
}
