#!/bin/bash

# 检查是否安装了 fw_printenv 工具
if ! command -v fw_printenv &> /dev/null
then
    echo "fw_printenv could not be found. Please install it first."
    exit 1
fi

# 列出所有 U-Boot 环境变量
fw_printenv