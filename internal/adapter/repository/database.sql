
-- users
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL COMMENT '사용자 ID',
  `user_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '사용자 이름',
  `user_email` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '사용자 이메일',
  `memo` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '사용자 메모',
  `use_yn` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '사용자 정보 사용여부',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
-- INSERT INTO `test`.`users` (`id`, `name`, `email`) VALUES (<{id: }>, <{name: }>, <{email: }>);
INSERT INTO `test`.`users` (`id`, `user_name`, `user_email`, `memo`, `use_yn`) VALUES ('1', 'Lydia', 'comeevery@gmail.com', `testUser`, `Y`);

