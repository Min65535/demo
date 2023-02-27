#!/usr/bin/env bash

set -e

## env
ENV=$1
IMAGE=$2
SVC=$3
case $ENV in
"dev") ;;
"test") ;;
"prod") ;;
*)
  echo "run dev|test|prod"
  exit 0
  ;;
esac
if [ "$IMAGE" == "" ]; then
  echo "empty image"
  exit 0
fi
echo "IMAGE: ""${IMAGE}"
if [ "$SVC" == "" ]; then
  echo "empty service"
  exit 0
fi
echo "service: ""${SVC}"

## docker run
echo "docker pull the IMAGE: ""${IMAGE}"
docker pull "${IMAGE}"
echo "docker successfully get the IMAGE: ""${IMAGE}"
echo "docker try to start service: ""${SVC}"
docker run -d --restart=always --name="${SVC}" -v /data/min/logs:/data/min/logs \
  --net=host "${IMAGE}" /min/"${SVC}"/bin/server \
  -config /min/"${SVC}"/configs/config."${ENV}".yaml
echo "docker start service: ""${SVC}"" successfully"

exit 0
