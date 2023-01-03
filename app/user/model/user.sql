CREATE TABLE `user` (
                        `id` int NOT NULL AUTO_INCREMENT COMMENT '用户id',
                        `username` varchar(64) NOT NULL COMMENT '用户名',
                        `password` varchar(128) NOT NULL COMMENT '用户密码',
                        `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表'