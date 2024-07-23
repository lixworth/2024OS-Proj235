#!/bin/bash

# 检查参数数量
if [ $# -ne 2 ]; then
    echo "Usage: $0 <image_path> <output_device>"
    exit 1
fi

# 获取传入的参数
image_path="$1"
output_device="$2"

# 使用dd命令写入镜像
echo "Writing image from $image_path to $output_device..."
dd if="$image_path" of="$output_device" bs=4M status=progress

# 优化分区（你可以在这里添加你的优化步骤）
sudo resize2fs $output_device

echo "Image writing completed!"
exit 0
