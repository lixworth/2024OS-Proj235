#!/bin/sh

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
check_and_install curl

echo "检查和安装完成。"