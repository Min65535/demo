#!/usr/bin/env bash

set -e

# shellcheck disable=SC2009
gamePid=$(ps -ef | grep 'game.yaml' | grep -v 'grep' | awk '{print $2}')
if [ "${gamePid}" -gt 0 ]; then kill -9 "${gamePid}"; fi

nohup ./game -config="./game.yaml" >game_nohup.out 2>&1 &

exit 0
