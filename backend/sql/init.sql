-- 用户表（更新版）
CREATE TABLE `user` (

  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `username` VARCHAR(50) NOT NULL COMMENT '登录用户名（唯一）',
  `password` VARCHAR(256) NOT NULL COMMENT '双向加密后的密码',
  `account_status` ENUM('active', 'locked', 'disabled') NOT NULL DEFAULT 'active' COMMENT '账号状态',
  `nickname` VARCHAR(50) NOT NULL COMMENT '仅供查看',
  `avatar_path` VARCHAR(255) NOT NULL DEFAULT 'default-avatar.png' COMMENT '头像路径(相对路径)',
  `tags` VARCHAR(255) NOT NULL DEFAULT '[]' COMMENT '用户tag',
  `signature` VARCHAR(255) NOT NULL DEFAULT '这个人啥也没说' COMMENT '个性签名',
  
  `username_en` VARCHAR(100) NOT NULL COMMENT '英文姓名',
  `username_zh` VARCHAR(100) NOT NULL COMMENT '中文姓名',
  `sex` ENUM('男', '女', '其他') NOT NULL COMMENT '性别',
  `cultivate_type` VARCHAR(50) NOT NULL COMMENT '培养类型',
  `department` VARCHAR(100) NOT NULL COMMENT '院系',
  `grade` VARCHAR(20) NOT NULL COMMENT '年级',
  `level` VARCHAR(50) NOT NULL COMMENT '学历层次',
  `student_type` VARCHAR(50) NOT NULL COMMENT '学生类型',
  `major` VARCHAR(100) NOT NULL COMMENT '专业',
  `class` VARCHAR(50) NOT NULL COMMENT '班级',
  `campus` VARCHAR(50) NOT NULL COMMENT '校区',
  `status` VARCHAR(50) NOT NULL COMMENT '学籍状态',
  `length` VARCHAR(20) NOT NULL COMMENT '学制',
  `enrollment_date` VARCHAR(50) NOT NULL COMMENT '入学日期',
  `graduate_date` VARCHAR(50) NOT NULL COMMENT '毕业日期',

  

  `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';


