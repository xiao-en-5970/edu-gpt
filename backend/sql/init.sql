-- 用户表（更新版）
CREATE TABLE `user` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `username` VARCHAR(50) NOT NULL COMMENT '登录用户名（唯一）',
  `password_hash` VARCHAR(255) NOT NULL COMMENT '加密后的密码',
  `email` VARCHAR(100) DEFAULT NULL COMMENT '邮箱（可选）',
  `phone` VARCHAR(20) DEFAULT NULL COMMENT '手机号（可选）',
  `account_status` ENUM('active', 'locked', 'disabled') NOT NULL DEFAULT 'active' COMMENT '账号状态',
  `last_login_time` TIMESTAMP NULL DEFAULT NULL COMMENT '最后登录时间',
  
  `nickname` VARCHAR(50) NOT NULL COMMENT '可选，仅供查看',
  `avatar_path` VARCHAR(255) NOT NULL DEFAULT 'default-avatar.png' COMMENT '头像路径(相对路径)',
  `self_evaluated_level` VARCHAR(20) DEFAULT NULL COMMENT '自评技术水平',
  `system_score` INT DEFAULT NULL COMMENT '系统评估得分(0~100)',
  `personality_tags` JSON DEFAULT NULL COMMENT '性格标签',
  `play_style_tags` JSON DEFAULT NULL COMMENT '打球风格',
  `preferred_skill_level` VARCHAR(20) DEFAULT NULL COMMENT '希望对手技术水平',
  `preferred_time_slots` JSON DEFAULT NULL COMMENT '时间偏好',
  `preferred_regions` JSON DEFAULT NULL COMMENT '常活动区域',
  `max_cost` INT DEFAULT NULL COMMENT '可接受的花销（单位：元）',
  `historical_partners` JSON DEFAULT NULL COMMENT '历史搭档ID列表',
  `ratings_given` JSON DEFAULT NULL COMMENT '对别人的评价(用户ID:评分)',
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`),
  UNIQUE KEY `idx_email` (`email`),
  UNIQUE KEY `idx_phone` (`phone`),
  CHECK (`system_score` BETWEEN 0 AND 100),
  CHECK (`max_cost` >= 0)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';


