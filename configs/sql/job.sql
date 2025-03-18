CREATE DATABASE IF NOT EXISTS `gojob`;
USE `gojob`;

-- 核心任务表
-- task_sharding 分片主表
DROP TABLE IF EXISTS `t_task_sharding`;
CREATE TABLE `t_task_sharding` (
    `shard_id` int(11) NOT NULL COMMENT '分片ID',
    `task_type` varchar(64) NOT NULL COMMENT '任务类型',
    `start_id` bigint(20) NOT NULL COMMENT '任务分片起始ID',
    `end_id` bigint(20) NOT NULL COMMENT '任务分片结束ID',
    `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`shard_id`),
    UNIQUE KEY `idx_task_type` (`task_type`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '分片元信息表';

-- task_type 任务类型策略配置表
DROP TABLE IF EXISTS `t_task_type`;
CREATE TABLE `t_task_type` (
  `type_id` varchar(64) NOT NULL COMMENT '任务类型ID(如email/sms)',
  `schedule_policy` varchar(512) NOT NULL COMMENT '调度策略(JSON格式, 含CRON/FIX_RATE等)',
  `retry_policy` varchar(512) NOT NULL COMMENT '重试策略(JSON格式,含max_retries/backoff等)',
  `priority` int NOT NULL DEFAULT '5' COMMENT '默认优先级(1-10)',
  `shard_count` int NOT NULL DEFAULT '4' COMMENT '该类型默认分片数',
  `enable_dynamic_config` tinyint(1) DEFAULT '0' COMMENT '是否允许任务覆盖配置',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`type_id`)
) ENGINE=InnoDB COMMENT='任务类型策略配置表';

-- task_type_0000 动态分片子表
DROP TABLE IF EXISTS `t_task_type_0000`;
CREATE TABLE `t_task_type_0000` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务ID - 自增id',
    `task_id` varchar(32) NOT NULL COMMENT '任务唯一ID(UUID|snowflake|MD5)',
    `task_type` varchar(64) NOT NULL COMMENT '任务类型',
    `executor_handler` varchar(255) NOT NULL COMMENT '执行器处理器',
    `sharding_param` varchar(100) DEFAULT NULL COMMENT '分片参数',
    `schedule_type` varchar(50) NOT NULL COMMENT '调度类型 CRON/FIX_RATE',
    `schedule_conf` varchar(128) NOT NULL COMMENT '调度配置',
    `task_context` text COMMENT '任务上下文(JSON格式)',
    `retry_count` int(11) NOT NULL DEFAULT '0' COMMENT '已重试次数',
    `trigger_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '触发状态 0-STOP 1-RUNNING',
    `trigger_last_time` bigint(20) DEFAULT NULL COMMENT '上次触发时间戳',
    `trigger_next_time` bigint(20) DEFAULT NULL COMMENT '下次触发时间戳',
    `create_time` bigint(20) NOT NULL COMMENT '创建时间戳',
    `update_time` bigint(20) NOT NULL COMMENT '更新时间戳',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_task_id` (`task_id`),
    KEY `idx_trigger_next_time` (`trigger_next_time`),
    KEY `idx_task_type_status` (`task_type`, `trigger_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='任务分片表';
--

-- -- task_executor 执行器注册表
-- DROP TABLE IF EXISTS `t_task_executor`;
-- CREATE TABLE `t_task_executor` (
--   `executor_id` varchar(64) NOT NULL COMMENT '执行器实例ID',
--   `app_name` varchar(255) NOT NULL COMMENT '应用名称',
--   `title` varchar(255) NOT NULL COMMENT '执行器标题',
--   `address_type` tinyint(4) NOT NULL COMMENT '注册方式 0-自动注册 1-手动录入',
--   `address_list` text COMMENT '执行器地址列表（逗号分隔）',
--   `update_time` bigint(20) NOT NULL COMMENT '最后心跳时间',
--   `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0-离线 1-在线',
--   PRIMARY KEY (`executor_id`),
--   KEY `idx_app_name` (`app_name`(20))
-- ) ENGINE=InnoDB COMMENT='执行器注册表';

-- -- t_log_sharding_route 日志分片路由表
-- DROP TABLE IF EXISTS `t_log_sharding_route`;
-- CREATE TABLE `t_log_sharding_route` (
--   `log_date` date NOT NULL COMMENT '日志日期',
--   `shard_table` varchar(50) NOT NULL COMMENT '分片表名',
--   PRIMARY KEY (`log_date`)
-- ) ENGINE=InnoDB COMMENT='日志分片路由表';

-- -- t_task_log_2025re01 动态日志分片表示例
-- DROP TABLE IF EXISTS `t_task_log_202501`;
-- CREATE TABLE `t_task_log_202501` (
--   `id` bigint(20) NOT NULL AUTO_INCREMENT,
--   `task_id` char(32) NOT NULL COMMENT '关联任务ID',
--   `trigger_time` bigint(20) NOT NULL COMMENT '调度时间戳',
--   `trigger_code` int(11) NOT NULL COMMENT '调度结果 200-成功 500-失败',
--   `trigger_msg` text COMMENT '调度日志',
--   `handle_time` bigint(20) DEFAULT NULL COMMENT '执行时间戳',
--   `handle_code` int(11) DEFAULT NULL COMMENT '执行结果',
--   `handle_msg` text COMMENT '执行日志',
--   `alarm_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '告警状态 0-未告警 1-已告警',
--   PRIMARY KEY (`id`),
--   KEY `idx_task_id` (`task_id`),
--   KEY `idx_trigger_time` (`trigger_time`)
-- ) ENGINE=InnoDB COMMENT='任务日志分片表';

-- -- task_job_info 任务依赖关系表
-- DROP TABLE IF EXISTS `t_task_job_info`;
-- CREATE TABLE `t_task_dependency` (
--   `task_id` char(32) NOT NULL,
--   `pre_task_id` char(32) NOT NULL COMMENT '前置任务ID',
--   `condition_type` varchar(50) NOT NULL COMMENT '依赖条件 SUCCESS/FAIL/COMPLETE',
--   PRIMARY KEY (`task_id`,`pre_task_id`)
-- ) COMMENT='任务依赖关系表';

-- -- task_pipeline 任务流水线配置表
-- DROP TABLE IF EXISTS `t_task_pipeline`;
-- CREATE TABLE `t_task_pipeline` (
--   `pipeline_id` char(32) NOT NULL,
--   `pipeline_name` varchar(255) NOT NULL,
--   `task_flow` text NOT NULL COMMENT '任务流程图（JSON格式）',
--   PRIMARY KEY (`pipeline_id`)
-- ) COMMENT='任务流水线配置表';