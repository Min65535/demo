#!/usr/bin/env bash

demo_db --action=create --env=${QY_ENV} --docker_env=${QY_DOCKER_ENV}
demo_db --action=migrate --env=${QY_ENV} --docker_env=${QY_DOCKER_ENV}
demo --env=${QY_ENV} --docker_env=${QY_DOCKER_ENV} --qy_run_env=${QY_RUN_ENV}