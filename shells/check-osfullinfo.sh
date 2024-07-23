#!/bin/bash

# 获取操作系统信息
echo "操作系统版本：$(cat /etc/os-release | grep PRETTY_NAME | cut -d '=' -f 2 | tr -d '\"')"
echo "内核版本：$(uname -r)"
echo "处理器信息：$(lscpu | grep 'Model name' | cut -d ':' -f 2 | tr -d ' ')"