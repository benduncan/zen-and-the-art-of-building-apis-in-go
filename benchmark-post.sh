#!/bin/sh

echo "Go app docker, log post"
ab -k -c 40 -n 1000 "http://localhost:8888/log?logdata=temp:29.5,pressure:1060,humidity:90,co2:350,site:1" | grep "Requests per second"

echo "Go app metal, log post"
ab -k -c 40 -n 1000 "http://localhost:8080/log?logdata=temp:29.5,pressure:1060,humidity:90,co2:350,site:1" | grep "Requests per second"

echo "PHP app, log post"
ab -k -c 40 -n 1000 "http://localhost:9999/log.php?logdata=temp:29.5,pressure:1060,humidity:90,co2:350,site:1" | grep "Requests per second"

