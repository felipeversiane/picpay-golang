#!/usr/bin/env bash
# Use this script to wait for a service to become available

TIMEOUT=15
while getopts t: option
do
  case "${option}"
  in
    t) TIMEOUT=${OPTARG};;
  esac
done

shift $((OPTIND -1))
HOST=$1
PORT=$2

if [ -z "$HOST" ] || [ -z "$PORT" ]; then
  echo "Usage: $0 host port [-t timeout]"
  exit 1
fi

for i in `seq $TIMEOUT` ; do
  nc -z $HOST $PORT > /dev/null 2>&1
  result=$?
  if [ $result -eq 0 ] ; then
    exit 0
  fi
  sleep 1
done

echo "Timeout: Unable to connect to $HOST:$PORT"
exit 1
