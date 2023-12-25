#!/bin/bash

rm -rf /usr/local/go
sed -i '/export PATH=\$PATH:\/usr\/local\/go\/bin/d' ~/.zshrc
. ~/.zshrc