#!/bin/sh

echo "Go app docker, time sync"
ab -k -c 30 -n 5000  http://localhost:8888/time | grep "Requests per second"

echo "Go app metal, time sync"
ab -k -c 30 -n 5000  http://localhost:8080/time | grep "Requests per second"

echo "PHP app, time sync"
ab -k -c 30 -n 5000  http://localhost:9999/time.php | grep "Requests per second"

