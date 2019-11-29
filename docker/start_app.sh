#!/usr/bin/env bash

demo_db --action=create --env=${ENV} --docker_env=${DOCKER_ENV}
demo_db --action=migrate --env=${ENV} --docker_env=${DOCKER_ENV}
demo --env=${ENV} --docker_env=${DOCKER_ENV} --run_env=${RUN_ENV}