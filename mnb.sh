#!/usr/bin/env bash

##  ./start.sh activity-service test run

##  ./start.sh exp-level-service test run
##  ./start.sh props-service test run
##  ./start.sh reward-service test run

##  ./open.sh exp-level-service test run
##  ./open.sh props-service test run
##  ./open.sh reward-service test run

### 查找,例如此脚本叫mn.sh
#root=`pwd`
#names=`ls -lh|grep -v 'mn.sh'|grep -v 'ert.sh'|awk 'NR>1{print $9}'`
#echo ${names}
#for i in ${names}
#do
#  cp ${root}/ert.sh ${root}/${i}/
#  cd ${root}/${i}/
#  ./ert.sh
#  cd ${root}
#done

### 删除,例如此脚本叫ert.sh
#names=`ls -lh|grep 3月|grep -v 'ert.sh'|awk 'NR>1{print $9}'`
#echo ${names}
#for i in ${names}
#do
#  rm -rf ${i}
#done

## shellcheck disable=SC2006
#names=`docker images -a|awk 'NR>1{print $3}'`
#echo "${names}"
#
#for i in ${names}
#do
#  docker rmi -f "${i}"
#done

## shellcheck disable=SC2006
#names=`docker ps -a|awk 'NR>1{print $1}'`
#echo "${names}"
#for i in ${names}
#do
#  echo "service: ""${i}"
#
#  echo "docker start to stop: ""${i}"
#  docker stop "${i}" && docker rm "${i}"
#  echo "${i}"" is shutdown"
#done
