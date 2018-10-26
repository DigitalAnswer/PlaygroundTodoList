Sample project to learn GoLang/MySQL/GraphQL
--

cmd Docker mysql
```shell
docker run -d --name mysql-dev -e MYSQL_ROOT_PASSWORD=admin --publish 6603:3306 --mount type=bind,source="$(pwd)"/datadir,target=/var/lib/mysql mysql
```

Connection:
![Connection to mysql database](https://github.com/DigitalAnswer/PlaygroundTodoList/blob/master/docs/img/ConnectMysql.png)

## Setup DB
``` sql
CREATE DATABASE `TodoDev` DEFAULT CHARACTER SET = `utf8`;

DROP TABLE IF EXISTS User;
CREATE TABLE `User` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) NOT NULL DEFAULT '',
  `first_name` varchar(50) NOT NULL DEFAULT '',
  `last_name` varchar(50) NOT NULL DEFAULT '',
  `password_hash` varchar(100) NOT NULL DEFAULT '',
  `password_salt` varchar(50) NOT NULL DEFAULT '',
  `is_disable` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `list` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `description` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `task_status` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(11) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `task` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `order` int(11) NOT NULL,
  `task_status_id` int(11) unsigned NOT NULL DEFAULT '1',
  `list_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `list_order` (`list_id`,`order`),
  KEY `task_status_id` (`task_status_id`),
  CONSTRAINT `task_list_id` FOREIGN KEY (`list_id`) REFERENCES `list` (`id`),
  CONSTRAINT `task_status_id` FOREIGN KEY (`task_status_id`) REFERENCES `task_status` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_list` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `list_id` int(11) unsigned NOT NULL,
  `order` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_order` (`user_id`,`order`),
  KEY `list_id` (`list_id`),
  CONSTRAINT `list_id` FOREIGN KEY (`list_id`) REFERENCES `list` (`id`),
  CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SELECT `list`.`name`
FROM `user_list`
INNER JOIN `list`
ON `user_list`.`list_id`=`list`.`id`
INNER JOIN user
ON user_list.`user_id`=user.`id`

```
