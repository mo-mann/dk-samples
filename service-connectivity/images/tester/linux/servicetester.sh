#!/bin/sh
while :
do
    echo "*********************************************"
    echo "[`date`]: Checking the URL $SERVICE_URL"
    2>/dev/stderr wget -O /dev/stdout $SERVICE_URL
    echo "[`date`]: Waiting $SERVICE_DELAY"
    sleep $SERVICE_DELAY
done