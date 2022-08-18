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

    docker rmi -f tplgo:$tag
    docker build -t tplgo:$tag .
    docker image prune -f
    # docker push registry/tplgo:$tag

    exit 0
fi

echo "👉 测试环境 ($tag)"

docker rmi -f tplgo_beta:$tag
docker build -t tplgo_beta:$tag .
docker image prune -f
# docker push registry/tplgo_beta:$tag
