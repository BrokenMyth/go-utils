#!/bin/sh

config_file="/etc/ssh/sshd_config"

key="$1"
rm ~/.ssh/authorized_keys
echo "$key" >> ~/.ssh/authorized_keys

# 使用 sed 替换匹配行的内容
sed -i 's/^.*PubkeyAuthentication.*/PubkeyAuthentication yes/' "$config_file"

sed -i 's/^.*AuthorizedKeysFile.*/AuthorizedKeysFile  .ssh\/authorized_keys/' "$config_file"

systemctl restart sshd