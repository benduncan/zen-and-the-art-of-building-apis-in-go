<?php

header("Content-Type: application/json");

file_put_contents("/tmp/iot-sensor-php.log", sprintf("%d => %s\n", time(), $_GET['logdata']), FILE_APPEND);

echo json_encode(array('Status' => "OK"));