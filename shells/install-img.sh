#!/bin/bash

# 检查传入的参数是否正确
if [ $# -ne 2 ]; then
    echo "Usage: $0 <镜像目录> <写入目标>"
    exit 1
fi

# 设置镜像目录和写入目标
image_dir="$1"
write_target="$2"

# 检查镜像目录是否存在
if [ ! -d "$image_dir" ]; then
    echo "镜像目录不存在: $image_dir"
    exit 1
fi

# 遍历镜像目录中的所有镜像文件
for image_file in "$image_dir"/*.img; do
    # 提取镜像名称（去除路径和扩展名）
    image_name=$(basename "$image_file" .img)

    # 写入镜像到目标设备
    dd if="$image_file" of="$write_target" bs=4M status=progress

    # 优化镜像（例如，移除不必要的层）
    sudo resize2fs $img_target

    # 这里可以根据实际需求添加其他优化步骤

    echo "镜像 $image_name 写入完成到 $write_target"
done

echo "所有镜像导入和优化完成！"