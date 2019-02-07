<?php

$time = time(true);

header("Content-Type: application/json");
echo json_encode(array('Status' => "OK", 'Payload' => $time));

?>