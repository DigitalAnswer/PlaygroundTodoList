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

```
