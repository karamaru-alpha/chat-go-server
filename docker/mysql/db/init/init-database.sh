#!/bin/bash

mysql -u$MYSQL_USER -p$MYSQL_ROOT_PASSWORD $MYSQL_DATABASE < "/docker-entrypoint-initdb.d/001-ddl.sql"
mysql -u$MYSQL_USER -p$MYSQL_ROOT_PASSWORD $MYSQL_DATABASE < "/docker-entrypoint-initdb.d/002-dml.sql"
