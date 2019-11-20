#!/usr/bin/env bash

root=$(pwd)

if [ $1 ];then
    # 编译项目
    echo "正在编译项目:demo"
    rm -rf tmp
    mkdir -p tmp

    cd ../
    GOOS=linux go build -o ${root}/tmp/demo
    cd db
    GOOS=linux go build -o ${root}/tmp/demo_db
else
    echo "请选择运行环境！"
    exit
fi

if [ $? -ne 0 ];then
    exit
fi
echo "项目编译成功"
cd ${root}

image_name="phl.dev.dr:5000/qy_demo:1.0"

if [ $1 == "test" ];then    
    image_name="phl.test.dr:5000/qy_demo:1.0"
fi

if [ $1 == "preprod" ];then    
    image_name="phl.preprod.dr:5000/qy_demo:2.0"
fi

if [ $1 == "prod" ];then    
    image_name="phl.prod.dr:5000/qy_demo:2.0"
fi

#echo "删除旧的镜像"
docker rmi -f ${image_name}
echo -e "\n正在生成docker镜像"
docker build -f ./Dockerfile -t ${image_name} ./

if [ $2 ];then
    docker push ${image_name}
fi
