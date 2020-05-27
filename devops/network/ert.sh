#!/usr/bin/env bash

inNames=`ls -lh|awk '{print $9}'`
echo ${inNames}

for i in ${inNames}
do
  echo "${i}_hahahah"
done
