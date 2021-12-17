#!/usr/bin/env bash

root=$1

env=$2

# shellcheck disable=SC2164
cd "${root}"/dfm-test/cmd/mysql-db-migrate

if [ -z "${env}" ]; then
  # shellcheck disable=SC2209
  env=test
fi

#db_env=${env} go run main.go drop
db_env=${env} go run main.go create
db_env=${env} go run main.go migrate
