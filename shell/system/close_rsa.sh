#!/bin/sh

config_file="/etc/ssh/sshd_config"

# 使用 sed 替换匹配行的内容
sed -i 's/^.*PubkeyAuthentication.*/PubkeyAuthentication no/' "$config_file"
sed -i '/AuthorizedKeysFile/s/^/#/' "$config_file"

systemctl restart sshd