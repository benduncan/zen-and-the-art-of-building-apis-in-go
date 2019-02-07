#!/bin/sh

docker run -p 9999:80 -d --name php-api php-api-example

docker run -p 8888:8080 -d --name go-api go-api-example
