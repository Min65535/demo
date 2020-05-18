#!/bin/bash

if [ $# != 2 ] ; then
echo "USAGE: $0 <server_name> <commit_hash>"
echo " e.g.: $0 wallet 41902691cc3869a63d7d1c1c1a57574d1c043213"
exit 1;
fi

paramone=$1
paramtwo=$2

echo $paramone
echo $paramtwo



echo "---approvalmq start---"
paramone=approvalmq
paramtwo=eb056e5dff0fb94c5ec59b6734a2a2aee62e7ddd
echo $paramone
echo $paramtwo
depimg=xl.image:5000/$paramone_$paramtwo
prdimg=xl.image.prod:5000/$paramone_$paramtwo
docker tag $depimg $prdimg
docker push $prdimg
echo "---approvalmq finish---"

echo "---approveadm start---"
paramone=approveadm
paramtwo=eb056e5dff0fb94c5ec59b6734a2a2aee62e7ddd
echo $paramone
echo $paramtwo
depimg=xl.image:5000/$paramone_$paramtwo
prdimg=xl.image.prod:5000/$paramone_$paramtwo
docker tag $depimg $prdimg
docker push $prdimg
echo "---approveadm finish---"

echo "---dpbuzmid start---"
paramone=dpbuzmid
paramtwo=eb056e5dff0fb94c5ec59b6734a2a2aee62e7ddd
echo $paramone
echo $paramtwo
depimg=xl.image:5000/$paramone_$paramtwo
prdimg=xl.image.prod:5000/$paramone_$paramtwo
docker tag $depimg $prdimg
docker push $prdimg
echo "---dpbuzmid finish---"

echo "---dptimer start---"
paramone=dptimer
paramtwo=eb056e5dff0fb94c5ec59b6734a2a2aee62e7ddd
echo $paramone
echo $paramtwo
depimg=xl.image:5000/$paramone_$paramtwo
prdimg=xl.image.prod:5000/$paramone_$paramtwo
docker tag $depimg $prdimg
docker push $prdimg
echo "---dptimer finish---"

echo "---ivbuzmid start---"
paramone=ivbuzmid
paramtwo=eb056e5dff0fb94c5ec59b6734a2a2aee62e7ddd
echo $paramone
echo $paramtwo
depimg=xl.image:5000/$paramone_$paramtwo
prdimg=xl.image.prod:5000/$paramone_$paramtwo
docker tag $depimg $prdimg
docker push $prdimg
echo "---ivbuzmid finish---"

echo "---rebates start---"
paramone=rebates
paramtwo=eb056e5dff0fb94c5ec59b6734a2a2aee62e7ddd
echo $paramone
echo $paramtwo
depimg=xl.image:5000/$paramone_$paramtwo
prdimg=xl.image.prod:5000/$paramone_$paramtwo
docker tag $depimg $prdimg
docker push $prdimg
echo "---rebates finish---"


echo "---rebatestimer start---"
paramone=rebatestimer
paramtwo=eb056e5dff0fb94c5ec59b6734a2a2aee62e7ddd
echo $paramone
echo $paramtwo
depimg=xl.image:5000/$paramone_$paramtwo
prdimg=xl.image.prod:5000/$paramone_$paramtwo
docker tag $depimg $prdimg
docker push $prdimg
echo "---rebatestimer finish---"

echo "---tsadmin start---"
paramone=tsadmin
paramtwo=eb056e5dff0fb94c5ec59b6734a2a2aee62e7ddd
echo $paramone
echo $paramtwo
depimg=xl.image:5000/$paramone_$paramtwo
prdimg=xl.image.prod:5000/$paramone_$paramtwo
docker tag $depimg $prdimg
docker push $prdimg
echo "---tsadmin finish---"

echo "---tsbuzmid start---"
paramone=tsbuzmid
paramtwo=eb056e5dff0fb94c5ec59b6734a2a2aee62e7ddd
echo $paramone
echo $paramtwo
depimg=xl.image:5000/$paramone_$paramtwo
prdimg=xl.image.prod:5000/$paramone_$paramtwo
docker tag $depimg $prdimg
docker push $prdimg
echo "---tsbuzmid finish---"

echo "---ucbuzmid start---"
paramone=ucbuzmid
paramtwo=eb056e5dff0fb94c5ec59b6734a2a2aee62e7ddd
echo $paramone
echo $paramtwo
depimg=xl.image:5000/$paramone_$paramtwo
prdimg=xl.image.prod:5000/$paramone_$paramtwo
docker tag $depimg $prdimg
docker push $prdimg
echo "---ucbuzmid finish---"

echo "---configcenter start---"
paramone=ucbuzmid
paramtwo=eb056e5dff0fb94c5ec59b6734a2a2aee62e7ddd
echo $paramone
echo $paramtwo
depimg=xl.image:5000/$paramone_$paramtwo
prdimg=xl.image.prod:5000/$paramone_$paramtwo
docker tag $depimg $prdimg
docker push $prdimg
echo "---configcenter finish---"

echo "---lottery start---"
paramone=lottery
paramtwo=c31f3a89468a8977c416d96a278437d86fd39174
echo $paramone
echo $paramtwo
depimg=xl.image:5000/$paramone_$paramtwo
prdimg=xl.image.prod:5000/$paramone_$paramtwo
docker tag $depimg $prdimg
docker push $prdimg
echo "---lottery finish---"

echo "---gateway start---"
paramone=gateway
paramtwo=2e3ff808864f5a7b8c0b9edc0fae365dab935df9
echo $paramone
echo $paramtwo
depimg=xl.image:5000/$paramone_$paramtwo
prdimg=xl.image.prod:5000/$paramone_$paramtwo
docker tag $depimg $prdimg
docker push $prdimg
echo "---gateway finish---"