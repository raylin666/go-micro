CREATE TABLE `link_relation` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ident` bigint unsigned NOT NULL COMMENT '唯一标识ID',
  `value` varchar(30) NOT NULL COMMENT '唯一标识值',
  `long_url` varchar(500) NOT NULL COMMENT '长链接地址',
  `state` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态 0:关闭 1:开启',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_ident` (`ident`) USING BTREE,
  UNIQUE KEY `uk_value` (`value`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='短链接关系表';