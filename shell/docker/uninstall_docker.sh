#!/bin/bash

# 安装所需工具
apt-get install -y curl awk sudo

# 执行指令，并自动选择是
apt-get autoremove docker docker-ce docker-engine docker.io containerd runc -y
dpkg -l | grep ^rc | awk '{print $2}' | sudo xargs dpkg -P -y
apt-get autoremove docker-ce-* -y
rm -rf /etc/systemd/system/docker.service.d
rm -rf /var/lib/docker

# 显示 Docker 版本
docker --version
echo "docker has been uninstalled!"