#!/bin/sh

docker stop php-api
docker stop go-api
docker rm php-api
docker rm go-api

