#!/bin/bash
docker rm -f tplgo
docker rmi -f img_tplgo

docker build -t img_tplgo .
docker image prune -f

docker run -d --name=tplgo --restart=always --privileged -p 10086:8000 -v /data/tplgo:/data img_tplgo
