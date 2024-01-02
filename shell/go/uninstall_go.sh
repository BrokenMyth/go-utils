#!/bin/bash

rm -rf /usr/local/go
sed -i '/export PATH=\$PATH:\/usr\/local\/go\/bin/d' ~/.bashrc
sed -i '/export PATH="$PATH:$(go env GOPATH)\/bin"/d' ~/.bashrc
source ~/.bashrc