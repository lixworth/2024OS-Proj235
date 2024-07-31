#!/bin/bash

# 获取所有分区信息
partitions=$(lsblk -o NAME,SIZE,FSTYPE,MOUNTPOINT -n -l)

# 输出表头
echo "分区名称 | 分区大小 | 分区格式 | 挂载点"
echo "------------------------------------"

# 遍历分区信息并输出
while read -r line; do
    name=$(echo "$line" | awk '{print $1}')
    size=$(echo "$line" | awk '{print $2}')
    format=$(echo "$line" | awk '{print $3}')
    mount_point=$(echo "$line" | awk '{print $4}')

    echo "$name | $size | $format | $mount_point"
done <<< "$partitions"