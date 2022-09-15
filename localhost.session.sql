create table sys_user_problem (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sys_user_id` bigint(20) unsigned DEFAULT NULL COMMENT '管理ID',
  `problem` varchar(255) not null comment '问题',
  `answer` varchar(255) not null comment '答案',
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;