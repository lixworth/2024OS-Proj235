#!/bin/bash

# 检查参数数量
if [ $# -ne 1 ]; then
    echo "Usage: $0 <img_path>"
    exit 1
fi

# 获取传入的镜像文件路径
img_path="$1"

# 使用dd命令将镜像写入/dev/sda
sudo dd if="$img_path" of=/dev/sda bs=4M status=progress

echo "镜像已成功写入 /dev/sda。"
