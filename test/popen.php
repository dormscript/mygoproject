<?php
error_reporting(E_ALL);


$descriptorspec = array( 0 => array("pipe", "r"), 1 => array("pipe", "w"));
$handle = proc_open('/webroot/go/src/test/test', $descriptorspec, $pipes);

$fp = fopen("title.txt", "rb");

while (!feof($fp)) {
    fwrite($pipes['0'], trim(fgets($fp))."\n");
    echo fgets($pipes[1]);
}

fclose($pipes['0']);
fclose($pipes['1']);
proc_close($handle);
