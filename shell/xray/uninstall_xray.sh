#!/bin/bash
docker rm $(docker ps -a | awk '/xray/ {print $1}') -f
docker rmi $(docker images | awk '/xray/ {print $3}')
echo "xray has been uninstalled!"