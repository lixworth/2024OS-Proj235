#!/bin/bash

# 获取网页上的版本信息
web_version=$(curl -s "https://pomhub.xyz/update.txt" | grep -oE '[0-9]{4}[0-9]{2}[0-9]{2}')

# 获取本地操作系统版本（os-release 信息 + PATCH version）
local_version=$(grep -E '^VERSION_ID=' /etc/os-release | cut -d '=' -f 2)

# 比较版本
if [ "$web_version" \> "$local_version" ]; then
    echo "有更新可用！网页版本：$web_version，本地版本：$local_version"
else
    echo "本地操作系统已是最新版本。"
fi