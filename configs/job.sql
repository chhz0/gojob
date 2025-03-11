
CREATE DATABASE `gojob`;

USE `gojob`;

CREATE TABLE jobs (
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT '任务ID - 自增id',
    unique_id VARCHAR(36) COMMENT '任务唯一ID(UUID|snowflake)',
    name VARCHAR(255) NOT NULL COMMENT '任务名称',
    type_id BIGINT UNSIGNED NOT NULL COMMENT '任务类型ID',
    status ENUM('pending', 'inprocess', 'finished', 'failed', 'paused') NOT NULL DEFAULT 'pending' COMMENT '任务状态',
    priority TINYINT NOT NULL DEFAULT 0 COMMENT '优先级（数值越小优先级越高）',
    next_run_time DATETIME COMMENT '下一次执行时间（用于调度）',
    config JSON COMMENT '任务配置（如超时时间、参数等）',
    retry_count INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '重试次数',
    last_error TEXT COMMENT '最近一次错误信息',
    queue_id INT UNSIGNED NOT NULL COMMENT '所属队列ID(外键)',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (task_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE job_types (
    type_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL COMMENT '类型名称（如：定时任务、队列任务）',
    handler VARCHAR(255) NOT NULL COMMENT '任务处理函数/类路径',
    description TEXT COMMENT '类型描述'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE job_history (
    history_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    task_id VARCHAR(36) NOT NULL COMMENT '关联任务ID',
    start_time DATETIME NOT NULL COMMENT '开始时间',
    end_time DATETIME COMMENT '结束时间',
    status ENUM('success', 'failed', 'timeout', 'ending') NOT NULL COMMENT '执行结果',
    result TEXT COMMENT '执行结果数据',
    error_message TEXT COMMENT '错误信息',
    FOREIGN KEY (task_id) REFERENCES tasks(task_id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;