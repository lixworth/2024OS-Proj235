#!/bin/bash

# 检查是否传入了参数
if [ -z "$1" ]; then
  echo "Usage: $0 <rootfs_location>"
  exit 1
fi

# 获取传入的参数
ROOTFS_LOCATION=$1

# 设置 U-Boot 环境变量
fw_setenv part rootfs $ROOTFS_LOCATION

# 检查设置是否成功
if [ $? -eq 0 ]; then
  echo "Successfully set part rootfs to $ROOTFS_LOCATION"
else
  echo "Failed to set part rootfs"
fi

# how 2 use?
#./set_rootfs.sh /dev/sdX