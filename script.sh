#!/bin/bash
docker cp ./data.sql go-mysql-binlog:/

docker exec go-mysql-binlog /bin/bash -c  'MYSQL_PWD="123456"  mysql --user=root  < /data.sql'