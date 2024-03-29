#!/bin/bash

wget https://golang.org/dl/go1.21.5.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.bashrc
. ~/.bashrc
go version