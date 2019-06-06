DROP DATABASE IF EXISTS sample;
CREATE DATABASE sample;
USE sample;

DROP TABLE IF EXISTS `users`;

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
