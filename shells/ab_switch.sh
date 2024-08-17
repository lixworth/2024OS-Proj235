!/bin/bash

# 检查参数数量
if [ $# -ne 2 ]; then
    echo "Usage: $0 <a|b> <mount_point>"
    exit 1
fi

# 获取传入的参数
option="$1"
mount_point="$2"

# 检查是否为有效的选项
if [ "$option" != "a" ] && [ "$option" != "b" ]; then
    echo "Invalid option. Please use 'a' or 'b'."
    exit 1
fi

# 修改 /etc/fstab 文件
if [ "$option" == "a" ]; then
    sed -i "s|^UUID=.* / .*|UUID=your_uuid_here $mount_point ext4 defaults 0 1|" /etc/fstab
elif [ "$option" == "b" ]; then
    sed -i "s|^UUID=.* / .*|UUID=your_other_uuid_here $mount_point ext4 defaults 0 1|" /etc/fstab
fi

echo "Updated /etc/fstab with the new mount point."
