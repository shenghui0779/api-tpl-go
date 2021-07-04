#!/bin/bash
docker rm -f tplgo
docker rmi -f tplgo
docker build -t tplgo .
docker run -d --name=tplgo --privileged -p 10086:10086 -v /data/config/tplgo:/data/config -v /data/logs/tplgo:/data/logs tplgo