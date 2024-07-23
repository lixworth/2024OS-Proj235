#!/bin/bash

# 检查参数数量
if [ $# -ne 1 ]; then
    echo "用法: $0 <绝对路径文件>"
    exit 1
fi

# 计算 SHA-256 哈希值
file="$1"
hash=$(sha256sum "$file" | cut -d ' ' -f 1)

echo "文件 $file 的 SHA-256 哈希值为: $hash"
