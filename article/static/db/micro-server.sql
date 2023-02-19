CREATE TABLE `api_article` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(32) NOT NULL DEFAULT '' COMMENT '文章标题',
    `author` varchar(20) NOT NULL COMMENT '文章作者',
    `summary` varchar(140) NOT NULL DEFAULT '' COMMENT '文章摘要',
    `cover` varchar(160) NOT NULL DEFAULT '' COMMENT '文章封面图片',
    `sort` smallint unsigned NOT NULL DEFAULT '0' COMMENT '排序',
    `recommend_flag` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '文章推荐标识 0:未推荐，1:已推荐',
    `commented_flag` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '文章是否允许评论 1:允许，0:不允许',
    `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:已关闭 1:已开启',
    `view_count` int unsigned NOT NULL DEFAULT '0' COMMENT '文章浏览量',
    `comment_count` int unsigned NOT NULL DEFAULT '0' COMMENT '文章评论数',
    `collection_count` int unsigned NOT NULL DEFAULT '0' COMMENT '文章收藏量',
    `zan_count` int unsigned NOT NULL DEFAULT '0' COMMENT '文章点赞数',
    `share_count` int unsigned NOT NULL DEFAULT '0' COMMENT '文章分享数',
    `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '发布者编号',
    `last_commented_at` int unsigned NOT NULL DEFAULT '0' COMMENT '最新评论时间',
    `created_at` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` int unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_user` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文章表';

CREATE TABLE `api_article_extend` (
   `id` bigint unsigned NOT NULL AUTO_INCREMENT,
   `article_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '文章ID',
   `source` varchar(32) NOT NULL DEFAULT '' COMMENT '文章来源',
   `source_url` varchar(160) NOT NULL DEFAULT '' COMMENT '文章来源链接',
   `content` longtext NOT NULL COMMENT '文章正文',
   `keyword` varchar(255) NOT NULL DEFAULT '' COMMENT '文章关键词',
   `attachment_path` text COMMENT '文章附件路径',
   PRIMARY KEY (`id`),
   UNIQUE KEY `uqe_acticle` (`article_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文章扩展表';

CREATE TABLE `api_article_category` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `pid` int unsigned NOT NULL DEFAULT '0' COMMENT '上级分类',
    `name` varchar(12) NOT NULL COMMENT '分类名称',
    `color` varchar(25) NOT NULL DEFAULT '' COMMENT '分类颜色',
    `sort` smallint unsigned NOT NULL DEFAULT '0' COMMENT '排序',
    `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:关闭 1:开启',
    `created_at` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` int unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uqe_name` (`name`) USING BTREE,
    KEY `idx_pid` (`pid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文章分类表';

CREATE TABLE `api_article_category_relation` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `article_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '文章ID',
    `category_id` int unsigned NOT NULL DEFAULT '0' COMMENT '分类ID',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_article_category` (`article_id`, `category_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文章分类关系表';