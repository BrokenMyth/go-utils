#!/bin/bash

#x86_64下载二进制文件
wget https://johnvansickle.com/ffmpeg/builds/ffmpeg-git-amd64-static.tar.xz

#解压文件
tar xvf ffmpeg-git-*-static.tar.xz && rm -rf ffmpeg-git-*-static.tar.xz

#将ffmpeg和ffprobe可执行文件移至/usr/bin方便系统直接调用
mv ffmpeg-git-*/ffmpeg  ffmpeg-git-*/ffprobe /usr/bin/

#查看版本
ffmpeg -version