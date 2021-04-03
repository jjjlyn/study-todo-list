#!/bin/bash

docker run -d -p 6033:3306 --name=todo-db -e MYSQL_DATABASE=todo -e MYSQL_ROOT_PASSWORD=root mysql:5.7