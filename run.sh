#!/bin/bash
env="beta"
tag="latest"

usage="-v 👉 版本Tag (默认: latest)\n-p 👉 生产环境\n-h Help"

while getopts "v:ph" arg
do
    case $arg in
        v) tag="$OPTARG" ;;
        p) env="prod" ;;
        h) echo $usage; exit 0 ;;
        ?) break ;;
    esac
done

if [ $env = "prod" ]
then
    echo "👉 生产环境 ($tag)"

    docker rm -f tplgo
    docker rmi -f iiinsomnia/tplgo:$tag
    docker run -d --name=tplgo --restart=always --privileged -p 10086:8000 \
    -v /data/tplgo/config:/data/config \
    -v /data/tplgo/logs:/data/logs \
    iiinsomnia/tplgo:$tag

    exit 0
fi

echo "👉 测试环境 ($tag)"

docker rm -f beta_tplgo
docker rmi -f iiinsomnia/tplgo_beta:$tag
docker run -d --name=beta_tplgo --restart=always --privileged -p 20086:8000 \
-v /data/beta/tplgo/config:/data/config \
-v /data/beta/tplgo/logs:/data/logs \
iiinsomnia/tplgo_beta:$tag
