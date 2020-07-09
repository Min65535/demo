#!/usr/bin/env bash

root=$1
env=$2

cd ${root}/cmd/mysql-db-migrate

if [[ ! ${env} ]];then
    env=test
fi

db_env=${env} go run db.go drop
db_env=${env} go run db.go creation
db_env=${env} go run db.go migrate