---
theme: default
background: https://w.wallhaven.cc/full/3k/wallhaven-3k16pd.png
title: Linux的OTA升级系统
info: |
  Learn more at [Sli.dev](https://sli.dev)
class: text-center
drawings:
  persist: false
transition: slide-left
mdc: true
---

2024年全国大学生计算机系统能力大赛 操作系统设计赛(全国) OS功能挑战赛道 
# Linux的OTA升级系统

项目编号：Proj235 | 队伍名称：地铁行动2014 

---
transition: fade-out
---

# 需求分析与实现方案：AB分区切换进行系统升级
第一题 升级系统的升级功能实现 && 第二题 升级系统基础框架功能实现
* 单独将 home、opt、usr、var 等与应用配置相关的目录单独设置挂载点
* 建立AB分区，升级过程中，通过修改 `/etc/fstab` 实现切换分区
* 更改挂载点、使用dd刷写 （initrd,kernel,rootfs）镜像
* 建立默认启动应用，检测重启次数，超出限制回滚回另一系统并标记

<hr class="my-4">

* sys-update cli应用
  * sys-update switch_ab/update/check/reset-pwd [arg]
* webclient web交互界面
  * 升级包上传、校验、执行升级
  * 日志查询
  * 升级任务管理、升级队列状态查询

<style>
h1 {
  background-color: #2B90B6;
  background-image: linear-gradient(45deg, #4EC5D4 10%, #146b8c 20%);
  background-size: 100%;
  -webkit-background-clip: text;
  -moz-background-clip: text;
  -webkit-text-fill-color: transparent;
  -moz-text-fill-color: transparent;
}
</style>

---
transition: slide-up
level: 2
---

# 技术实现

升级程序总共分为 ota-updater 升级程序 与 ota-manage 管理平台俩部分， 其中 shells、client-ui 为 ota-updater 提供操作逻辑与操作界面，ota-manage 为独立的前后端分离架构WEB平台，提供管理界面与API。

### ota-updater 升级程序客户端
基于 Golang + Gin 构建，使用 [BadgerDB](https://github.com/dgraph-io/badger) 存储数据，通过 exec 执行 shells 中的命令实现系统级操作。此外除本体守护进程外，也开发了对应管理工具cli应用，用于管理守护进程、执行重置操作。

### client-ui 升级程序客户端界面
基于 Nuxt(Vue3) + NuxtUI 构建升级操作界面，与 ota-updater 交互实现系统升级、升级设置、日志查看等操作。

* 由 ota-updater 启动 HTTP服务器，由根目录输出 client-ui 前端构建的静态文件
* client-ui 通过 fetch 向后端 API 发起请求
* `/login` 登陆页面 基于浏览器 SessionStorage 和 Token 的鉴权-机制
* `/` 主页面 负责查看当前系统/版本信息、版本检查更新、手动更新、更新设置等操作

---
layout: image-right
image: https://cover.sli.dev
---

# 升级脚本运行过程

* check_img.sh 校验升级镜像文件
* write_image_by_dd.sh 用 dd 写入镜像
* ab_switch.sh 通过修改 `/etc/fstab` 实现AB分区的切换
* check-osfullinfo.sh 检测系统信息
* check-install-sysenv.sh 检测系统依赖
* check-sysupdate.sh 检测系统更新
* sysupdate.sh 系统更新

---
---

# 安装要求

<div grid="~ cols-2 gap-4">
<div>

* 适用于分区类型为MBR&GPT
* 磁盘分区要求: 
  * 一个启动分区(通常为efi)
  * 两个uuid不相同且分区大小相同的分区(文件格式不限,建议f2fs或ext4,使用dd复制的分区须更改为"LABEL=&pathname")
  * 一个数据分区(挂载常用及定制的目录)如:
    * opt ~ 
    * www ~ 是经典的bt面板所在地
    * home ~ 用户文件所在地,是切换系统的同时保留设置的最基本文件夹
    * var ~ 至少tftp serv默认位置在这里
    * snap ~ ubuntu常见

</div>
<div>

<img border="rounded" src="https://github.com/lixworth/2024OS-Proj235/raw/master/document/screenshot/image.png" alt="">

</div>
</div>

---
class: px-20
---

# 测试情况

| 操作系统       | 测试架构     | 启动方式    |  测试结果 |
|---------------|------------|------------|----------|
| Ubuntu/Debian | x86        | grub | &#10004; |
| openKylin     | x86        | grub | &#10004; |
| ArchLinux     | x86        | grub | &#10004; |
| openEuler     | x86        | grub | &#10004; |
| Deepin        | risc-v     | grub/uboot | &#10004; |
| AOSC          | x86        | grub | &#10004; |

---
class: px-20
---

# 使用与演示 

* 访问 http://127.0.0.1:9301/
  * 默认用户密码 passwordroot
  * 使用cli工具重置密码
* 修改升级服务器为 ota-manage 的API地址 例如: https://cscc.kokomi.ltd/api
* 检查更新/手动上传更新

<div grid="~ cols-2 gap-2" m="t-2">
<img border="rounded" src="https://github.com/lixworth/2024OS-Proj235/raw/master/document/screenshot/%E6%88%AA%E5%B1%8F2024-07-15%2022.03.59.png" alt="">
<img border="rounded" src="https://github.com/lixworth/2024OS-Proj235/raw/master/document/screenshot/%E6%88%AA%E5%B1%8F2024-07-28%2013.26.37.png" alt="">
</div>
---
layout: center
class: text-center
---

# 感谢观看

<PoweredBySlidev mt-10 />
