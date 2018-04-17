Sample project to learn GoLang/MySQL/GraphQL
--

cmd Docker mysql
```shell
docker run -d --name mysql-dev -e MYSQL_ROOT_PASSWORD=admin --publish 6603:3306 --mount type=bind,source="$(pwd)"/datadir,target=/var/lib/mysql mysql
```

Connection:
![Connection to mysql database](https://github.com/DigitalAnswer/PlaygroundTodoList/blob/master/docs/img/ConnectMysql.png)