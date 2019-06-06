DROP DATABASE IF EXISTS sample;
CREATE DATABASE sample;
USE sample;

DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `tasks`;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,          /* ID */
  `first_name` varchar(255) DEFAULT NULL,        /* 名 */
  `last_name` varchar(255) DEFAULT NULL,         /* 姓 */
  `email` varchar(255) NOT NULL,                 /* メールアドレス */
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8
;

CREATE TABLE `tasks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,          /* ID */
  `user_id` int(11) NOT NULL,                    /* USER ID */
  `title` varchar(255) DEFAULT NULL,             /* 名 */
  `content` varchar(255) DEFAULT NULL,           /* 姓 */
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  CONSTRAINT `fkey_tasks_users_user_id` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8
;
