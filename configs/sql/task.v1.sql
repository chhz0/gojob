# t_task
DROP TABLE IF EXISTS `t_lark_task_1`;
CREATE TABLE `t_lark_task_1` (
                                 `id` bigint NOT NULL AUTO_INCREMENT,
                                 `user_id` varchar(256),
                                 `task_id` varchar(256),
                                 `task_type` varchar(128),
                                 `task_stage` varchar(128),
                                 `status` tinyint(3) unsigned NOT NULL DEFAULT '0',
                                 `crt_retry_num` int(11) COMMENT '重试次数',
                                 `max_retry_num` int(11) COMMENT '最大重试次数',
                                 `max_retry_interval` int(11) COMMENT '最大重试间隔',
                                 `schedule_log` varchar(4096) COMMENT '调度信息记录',
                                 `priority` int(11) COMMENT '优先级',
                                 `task_context` varchar(8192) COMMENT '任务上下文， 用户自定义',
                                 `order_time` int(20) COMMENT '调度时间， 越小优先级越高',
                                 `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
                                 `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `idx_task_id` (`task_id`),
                                 KEY `idx_user_id` (`user_id`),
                                 KEY `idx_status` (`status`),
                                 KEY `idx_tasktype_status_modify_time` (`status`, `order_time`)
) ENGINE = InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

# 任务配置表

# t_***_cfg
DROP TABLE IF EXISTS `t_task_schedule_cfg`;
CREATE TABLE `t_task_schedule_cfg` (
                                  `task_type` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '',
                                  `schedule_limit` int(11) DEFAULT '0' COMMENT '任务调度限制',
                                  `schedule_interval` int(11) DEFAULT '10' COMMENT '任务调度间隔',
                                  `max_processing_time` int(11) DEFAULT '0' COMMENT '任务最大处理时间',
                                  `max_retry_num` int(11) DEFAULT '0' COMMENT '任务最大重试次数',
                                  `retry_interval` int(11) DEFAULT '0' COMMENT '重试间隔',
                                  `max_retry_interval` int(11) DEFAULT '0' COMMENT '任务最大重试间隔',
                                  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
                                  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  PRIMARY KEY (`task_type`)
)ENGINE = InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


# 任务位置表

# t_***_pos
DROP TABLE IF EXISTS `t_task_pos`;
CREATE TABLE `t_task_pos` (
                                  `id` bigint NOT NULL AUTO_INCREMENT,
                                  `task_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                                  `schedule_begin_pos` int(11) NOT NULL DEFAULT '0' COMMENT '调度开始于几号表',
                                  `schedule_end_pos` int(11) NOT NULL DEFAULT '0' COMMENT '调度结束于几号表',
                                  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
                                  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  PRIMARY KEY (`id`),
                                  UNIQUE KEY `idx_task_type` (`task_type`)
)ENGINE = InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;