#!/usr/bin/env bash

root=$(pwd)

if [ $1 ];then
    # compile the project
    echo "compiling:demo"
    rm -rf tmp
    mkdir -p tmp

    cd ../
    GOOS=linux go build -o ${root}/tmp/demo
    cd db
    GOOS=linux go build -o ${root}/tmp/demo_db
else
    echo "plz choose the env!!!"
    exit
fi

if [ $? -ne 0 ];then
    exit
fi
echo "compile successfully!!!"
cd ${root}

image_name="tem.dev.dr:5000/demo:1.0"

if [ $1 == "test" ];then    
    image_name="tem.test.dr:5000/demo:1.0"
fi

if [ $1 == "preprod" ];then    
    image_name="tem.preprod.dr:5000/demo:2.0"
fi

if [ $1 == "prod" ];then    
    image_name="tem.prod.dr:5000/demo:2.0"
fi

#echo "delete the old image"
go rmi -f ${image_name}
echo -e "\n create the new docker image"
go build -f ./Dockerfile -t ${image_name} ./

if [ $2 ];then
    go push ${image_name}
fi
