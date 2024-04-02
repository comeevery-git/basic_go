
-- users
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL COMMENT '사용자 ID',
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '사용자 이름',
  `email` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '사용자 이메일',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
-- INSERT INTO `test`.`users` (`id`, `name`, `email`) VALUES (<{id: }>, <{name: }>, <{email: }>);
INSERT INTO `test`.`users` (`id`, `name`, `email`) VALUES ('1', 'lydia', 'comeevery@gmail.com');

