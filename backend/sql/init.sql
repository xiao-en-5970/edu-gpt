-- 用户表（更新版）
CREATE TABLE `user` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '用户编号',
    `username` VARCHAR(50) NOT NULL COMMENT '登录用户名（唯一）',
    `password` VARCHAR(256) NOT NULL COMMENT '双向加密后的密码',
    `active` ENUM(
        'active',
        'locked',
        'disabled'
    ) NOT NULL DEFAULT 'active' COMMENT '状态',
    `nickname` VARCHAR(50) NOT NULL COMMENT '仅供查看',
    `avatar_path` VARCHAR(255) NOT NULL DEFAULT './static/avatars/default-avatar.png' COMMENT '头像路径(相对路径)',
    `backimage_path` VARCHAR(255) NOT NULL DEFAULT './static/backgrounds/default-image.png' COMMENT '背景图路径(相对路径)',
    `tags` VARCHAR(255) NOT NULL DEFAULT '[]' COMMENT '用户tag',
    `follow` BIGINT NOT NULL COMMENT '用户关注数量',
    `fans` BIGINT NOT NULL COMMENT '用户粉丝数量',
    `allpost_like` BIGINT NOT NULL COMMENT '用户点赞数量',
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
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '用户表';

CREATE TABLE `post` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '帖子ID',
    `poster_id` BIGINT COMMENT '发帖人id',
    `title` VARCHAR(200) NOT NULL COMMENT '标题',
    `content` TEXT NOT NULL COMMENT '内容',
    `community_id` BIGINT NOT NULL DEFAULT 1 COMMENT '社区id',
    `view_count` INT NOT NULL DEFAULT 0 COMMENT '浏览数',
    `like_count` INT NOT NULL DEFAULT 0 COMMENT '点赞数',
    `collect_count` INT DEFAULT 0 COMMENT '收藏数',
    `comment_count` INT DEFAULT 0 COMMENT '(被）评论数',
    `active` ENUM(
        'active',
        'locked',
        'disabled'
    ) NOT NULL DEFAULT 'active' COMMENT '状态',
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX `idx_title_prefix` (`title` (10))
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '帖子表';

CREATE TABLE `post_image` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '帖子图片ID',
    `post_id` BIGINT NOT NULL COMMENT '发帖人id',
    `number` INT NOT NULL COMMENT '第几张图片',
    `images_path` VARCHAR(255) DEFAULT 0 COMMENT '图片路径',
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_post` (`post_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '帖子图片表';

CREATE TABLE `post_likes` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `post_id` BIGINT NOT NULL COMMENT '帖子ID',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1-点赞 0-取消',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_post_user` (`post_id`, `user_id`),
    KEY `idx_user` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '用户点赞记录表';

CREATE TABLE `comment` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `post_id` BIGINT NOT NULL COMMENT '评论的帖子ID',
    `user_id` BIGINT NOT NULL COMMENT '发评论的用户ID',
    `content` TEXT NOT NULL COMMENT '内容',
    `like_count` INT NOT NULL DEFAULT 0 COMMENT '点赞数',
    `child_count` INT NOT NULL DEFAULT 0 COMMENT '子评论数量',
    `active` ENUM(
        'active',
        'locked',
        'disabled'
    ) NOT NULL DEFAULT 'active' COMMENT '状态',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_post` (`post_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '评论表';

CREATE TABLE `sub_comment` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `post_id` BIGINT NOT NULL COMMENT '评论的帖子ID',
    `user_id` BIGINT NOT NULL COMMENT '发评论的用户ID',
    `parent_comment_id` BIGINT NOT NULL DEFAULT 0 COMMENT '上层评论id',
    `reply_user_id` BIGINT NOT NULL DEFAULT 0 COMMENT '回复的子评论用户id,默认为0表示回复顶层评论',
    `content` TEXT COMMENT '内容',
    `like_count` INT DEFAULT 0 COMMENT '子评论点赞数',
    `active` ENUM(
        'active',
        'locked',
        'disabled'
    ) NOT NULL DEFAULT 'active' COMMENT '状态',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_parent` (`parent_comment_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '子评论表';

CREATE TABLE `comment_likes` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `comment_id` BIGINT NOT NULL COMMENT '评论ID',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1-点赞 0-取消',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_commment_user` (`comment_id`, `user_id`),
    KEY `idx_user` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '评论点赞记录表';

CREATE TABLE `subcomment_likes` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `subcomment_id` BIGINT NOT NULL COMMENT '子评论ID',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1-点赞 0-取消',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_subcomment_user` (`subcomment_id`, `user_id`),
    KEY `idx_user` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '子评论点赞记录表';

CREATE TABLE `user_follow` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `follow` BIGINT NOT NULL COMMENT '关注的用户ID',
    `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1-关注 0-取消关注',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_follow` (`user_id`, `follow`),
    KEY `idx_user` (`user_id`),
    KEY `idx_flo` (`follow`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '用户关注表';

CREATE TABLE `course` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `course_name` VARCHAR(200) NOT NULL COMMENT '课程名字',
    `course_code` VARCHAR(32) NOT NULL COMMENT '课程代码',
    `course_type` VARCHAR(32) NOT NULL COMMENT '课程类型',
    `credits` INT NOT NULL DEFAULT 0 COMMENT '学分',
    `open_depart` VARCHAR(32) NOT NULL COMMENT '开课学院' ,
    `description` VARCHAR(256) NOT NULL DEFAULT "尚无描述" COMMENT '课程描述',
    `view_count` INT NOT NULL DEFAULT 0 COMMENT '浏览数',
    `like_count` INT NOT NULL DEFAULT 0 COMMENT '点赞数',
    `comment_count` INT NOT NULL DEFAULT 0 COMMENT '(被）评论数',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `idx_cource` (`course_code` (8))
)