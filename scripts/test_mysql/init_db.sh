#!/bin/bash

set -e
set -x

/usr/sbin/mysqld --initialize-insecure

/usr/sbin/mysqld --user=root &
mysql_pid=$!

for i in {30..0}; do
	if echo 'SELECT 1' | mysql &> /dev/null; then
		break
	fi
	echo "waiting"
	sleep 1
done

mysql < /tmp/init.sql
mysql --database=gorani_test < /tmp/schema.sql

echo 'SHUTDOWN' | mysql
wait $mysql_pid

tar czvf default_mysql.tar.gz /var/lib/mysql
