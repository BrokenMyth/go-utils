#!/bin/bash

rm -rf /usr/local/go
sed -i '/export PATH=\$PATH:\/usr\/local\/go\/bin/d' ~/.zshrc
sed -i '/export PATH=\$PATH:\$GOPATH\/bin/d' ~/.zshrc
. ~/.zshrc