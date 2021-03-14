#!/bin/bash

mysql -u root -ppassword chat_go_server < "/docker-entrypoint-initdb.d/001-ddl.sql"
mysql -u root -ppassword chat_go_server < "/docker-entrypoint-initdb.d/002-dml.sql"
