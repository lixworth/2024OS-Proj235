#!/bin/bash

# 定义变量
GITHUB_URL="https://github.com/libiunc/2024OS-Proj235/tree/shell/upd.txt"
LOCAL_FILE="/usr/metro2014/upd.txt"
UPDATE_SCRIPT_URL="https://github.com/libiunc/2024OS-Proj235/tree/shell/upd.sh"
UPDATE_SCRIPT="/tmp/upd.sh"

# 获取 GitHub 文件中的 num 值
GITHUB_NUM=$(curl -s $GITHUB_URL | grep 'num' | awk -F '=' '{print $2}')

# 获取本地文件中的 num 值
LOCAL_NUM=$(grep 'num' $LOCAL_FILE | awk -F '=' '{print $2}')

# 比较 num 值并执行热更新
if [ "$GITHUB_NUM" -gt "$LOCAL_NUM" ]; then
  echo "GitHub num ($GITHUB_NUM) is greater than local num ($LOCAL_NUM). Executing hot update..."
  
  # 下载并执行更新脚本
  curl -s $UPDATE_SCRIPT_URL -o $UPDATE_SCRIPT
  chmod +x $UPDATE_SCRIPT
  bash $UPDATE_SCRIPT
else
  echo "Local num ($LOCAL_NUM) is up-to-date."
fi