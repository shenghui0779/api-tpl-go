#!/bin/bash
docker rm -f api
docker rmi -f img_api

docker build -t img_api .
docker image prune -f

docker run -d --name=api --restart=always --privileged -p 10086:8000 -v /data/api:/data img_api
