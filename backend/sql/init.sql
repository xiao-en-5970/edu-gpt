-- 用户表（更新版）
CREATE TABLE `user` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `username` VARCHAR(50) NOT NULL COMMENT '登录用户名（唯一）',
  `account_status` ENUM('active', 'locked', 'disabled') NOT NULL DEFAULT 'active' COMMENT '账号状态',
  `nickname` VARCHAR(50) NOT NULL COMMENT '仅供查看',
  `avatar_path` VARCHAR(255) NOT NULL DEFAULT 'default-avatar.png' COMMENT '头像路径(相对路径)',
  `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';


