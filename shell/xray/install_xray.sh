#!/bin/bash

port=9000
id="1eb6e917-774b-4a84-aff6-b058577c60a5"
docker pull teddysun/xray
mkdir -p /etc/xray
rm /etc/xray/config.json
cat > /etc/xray/config.json <<EOF
{
  "inbounds": [{
    "port": $port,
    "protocol": "vmess",
    "settings": {
      "clients": [
        {
          "id": "$id"
        }
      ]
    }
  }],
  "outbounds": [{
    "protocol": "freedom",
    "settings": {}
  }]
}
EOF
echo "write config success."
docker run -d -p 9000:9000 --name xray --restart=always -v /etc/xray:/etc/xray teddysun/xray