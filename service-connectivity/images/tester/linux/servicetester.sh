#!/bin/sh
while :
do
    echo "*********************************************"
    echo "[`date`]: Checking the URL $SERVICE_URL"
    2>/dev/stderr wget -O /dev/stdout $SERVICE_URL
    echo "[`date`]: Waiting 5 seconds"
    sleep 5s
done