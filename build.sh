#!/bin/bash
docker rm -f tplgo
docker rmi -f tplgo
docker build -t tplgo .
docker image prune -f
docker run -d --name=tplgo --privileged -p 10086:8000 -v /data/config/tplgo:/data/config -v /data/logs/tplgo:/data/logs tplgo