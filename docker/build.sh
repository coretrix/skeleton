#!/usr/bin/env bash

rm -f .env

echo COMPOSE_PROJECT_NAME=skeleton >> .env
echo LOCAL_IP=0.0.0.0 >> .env
echo MYSQL_PORT=8004 >> .env
echo REDIS_PORT=8001 >> .env
echo MAILCATCHER_PORT=8025 >> .env
echo MAILCATCHER_WEB_PORT=8080 >> .env

docker-compose up -d --build