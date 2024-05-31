#!/bin/bash

# 检查参数数量
if [ $# -ne 1 ]; then
    echo "Usage: $0 <img_path>"
    exit 1
fi

# 获取传入的镜像文件路径
img_path="$1"

# 使用fdisk命令获取镜像的分区信息
partition_info=$(sudo fdisk -l "$img_path")

# 检查是否存在名为 "boot" 的分区
if echo "$partition_info" | grep -q "boot"; then
    echo "镜像中存在名为 'boot' 的分区。"
else
    echo "镜像中不存在名为 'boot' 的分区。"
fi

# 检查是否存在名为 "systemos" 的分区
if echo "$partition_info" | grep -q "systemos"; then
    echo "镜像中存在名为 'systemos' 的分区。"
else
    echo "镜像中不存在名为 'systemos' 的分区。"
fi
