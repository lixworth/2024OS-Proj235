#!/bin/bash

web_version=$(curl -s "https://ota.1280768.xyz/version.txt" | grep -oE '[0-9]{4}[0-9]{2}[0-9]{2}')

# 网页上的镜像 URL
web_image_url="https://pomhub.xyz/upgrade/系统版本/$web_version.img"

# 本地保存路径
local_path="/usr/metro2014/upgrade/"

# 下载镜像
wget "$web_image_url" -P "$local_path"

# 检查下载是否成功
if [ $? -eq 0 ]; then
    echo "升级包下载完成！"
else
    echo "下载失败，请检查网页 URL 或本地路径。"
fi