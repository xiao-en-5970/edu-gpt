TRUNCATE TABLE `comment`;

ALTER TABLE user
ADD COLUMN `follow` BIGINT NOT NULL COMMENT '用户关注数量',
ADD COLUMN `fans` BIGINT NOT NULL COMMENT '用户粉丝数量',
ADD COLUMN `allpost_like` BIGINT NOT NULL COMMENT '用户点赞数量';


ALTER TABLE `post` 
ADD COLUMN `community_id` BIGINT NOT NULL DEFAULT 1 COMMENT '社区id';

ALTER TABLE post
ADD COLUMN comment_table_id BIGINT NOT NULL COMMENT '评论区ID';

ALTER TABLE comment
ADD COLUMN comment_table_id BIGINT NOT NULL COMMENT '评论区ID';

SELECT DISTINCT `course_type` from `course` ORDER BY  `course_type`