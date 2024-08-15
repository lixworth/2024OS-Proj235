#!/bin/sh
echo "1 install check env....."
mkdir /usr/metro2014/
# 检查包管理器
if command -v apt-get &>/dev/null; then
    package_manager="apt"
elif command -v yum &>/dev/null; then
    package_manager="yum"
elif command -v zypper &>/dev/null; then
    package_manager="zypper"
elif command -v pacman &>/dev/null; then
    package_manager="pacman"
elif command -v dnf &>/dev/null; then
    package_manager="dnf"
elif command -v aptitue &>/dev/null; then
    package_manager="aptitue"
elif command -v oma &>/dev/null; then
    package_manager="oma"
else
    echo "无法识别系统的包管理器。"
    exit 1
fi

# 检查指令是否存在并安装缺失的包
check_and_install() {
    if ! command -v "$1" &>/dev/null; then
        echo "缺失 $1，正在安装..."
        sudo "$package_manager" install -y "$1"
    fi
}

# 检查并安装缺失的指令
check_and_install busybox
check_and_install wget
check_and_install bash
check_and_install switchconf
check_and_install curl
check_and_install pv
check_and_install git




echo "检查和安装完成。"

# 定义变量
REPO_URL="https://github.com/libiunc/2024OS-Proj235.git"
TARGET_DIR="/usr/metro2014/shells"
FOLDER_TO_CLONE="shells"

# 创建并进入目标目录
mkdir -p $TARGET_DIR
cd $TARGET_DIR

# 初始化 Git 仓库
git init

# 添加远程仓库
git remote add origin $REPO_URL

# 启用稀疏检出
git sparse-checkout init --cone

# 设置要检出的文件夹
git sparse-checkout set $FOLDER_TO_CLONE

# 拉取内容
git pull origin main

echo "Successfully cloned $FOLDER_TO_CLONE from $REPO_URL into $TARGET_DIR enjoy~"# WTF?