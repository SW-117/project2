Create table user (
   `id` bigint(20) NOT NULL  AUTO_INCREMENT COMMENT "id",
    `name` varchar(255) NOT NULL DEFAULT '' COMMENT  "用户名" unique,
    `pass` varchar(1024) NOT NULL DEFAULT '' COMMENT "密码",
    `sex` tinyint(4) NOT NULL DEFAULT '0' COMMENT  "性别",
    `create_time` bigint(20) NOT NULL DEFAULT '0' COMMENT  "创建时间",
    `update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT  "更新时间",
   PRIMARY KEY ( `id` )         
)  ENGINE=InnoDB DEFAULT CHARSET=utf8;