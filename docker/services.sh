#!/usr/bin/env bash

cd docker && docker-compose exec services /bin/sh -c "cd cmd/$1 && APP_MODE=local go run main.go"
