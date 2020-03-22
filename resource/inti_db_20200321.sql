CREATE DATABASE `auto4play` CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `auto4play`;

-- ----------------------------
-- Table structure for record
-- ----------------------------
DROP TABLE IF EXISTS `record`;
CREATE TABLE `record` (
  `record_id` int NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `user_id` int NOT NULL COMMENT '用户ID',
  `record_submission` enum('auto','manual') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'manual' COMMENT '提交方式',
  `record_type` enum('sleep','fitness') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'sleep' COMMENT '记录类型',
  `status` enum('valid','invalid') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'valid' COMMENT '记录状态',
  `source` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '记录来源',
  `remark` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '记录备注信息',
  `created_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
  PRIMARY KEY (`record_id`),
  UNIQUE KEY `record_unique` (`record_id`) USING BTREE COMMENT 'RECORDID 唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;