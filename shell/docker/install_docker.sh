#!/bin/bash

curl -fsSL https://get.docker.com/ -o docker.sh
yes | sh docker.sh
rm docker.sh
echo "docker has been installed!"