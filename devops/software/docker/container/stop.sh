#!/usr/bin/env bash

set -e

SVC=$1

if [ "$SVC" == "" ]; then
  echo "empty service"
  exit 0
fi
echo "service: ""${SVC}"

echo "docker start to stop: ""${SVC}"
docker stop "${SVC}" && docker rm "${SVC}"
echo "${SVC}"" is shutdown"

exit 0
