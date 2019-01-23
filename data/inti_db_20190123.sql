CREATE DATABASE `auto4play` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE `auto4play`;

CREATE TABLE `record` (
  `record_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `user_id` int(16) NOT NULL COMMENT '用户ID',
  `record_submission` enum('auto','manual') NOT NULL COMMENT '提交方式',
  `record_type` enum('sleep') NOT NULL COMMENT '记录类型',
  `status` enum('valid','invalid') NOT NULL COMMENT '记录状态',
  `source` varchar(255) NOT NULL COMMENT '记录来源',
  `remark` text COMMENT '记录备注信息',
  `created_at` bigint(20) NOT NULL COMMENT '记录创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '记录更新时间',
  PRIMARY KEY (`record_id`),
  UNIQUE KEY `record_unique` (`record_id`) USING BTREE COMMENT 'RECORDID 唯一索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;