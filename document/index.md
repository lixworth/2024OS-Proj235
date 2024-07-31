# Linux的OTA升级系统 文档

## 基础信息
大连理工大学城市学院 T202413198993225 地铁行动2014 proj235 Linux的OTA升级系统

Github仓库：https://github.com/lixworth/2024OS-Proj235 <br>
Gitlab镜像仓库：https://gitlab.eduxiji.net/T202413198993225/project2210132-239709

## 需求方案分析
> https://github.com/oscomp/proj235-linux-upgrade-system

第一题：升级系统的升级功能实现

* 能够支持linux主要系统文件(initrd,kernel,rootfs)等linux重要文件的版本升级
  * 改挂载点、使用dd刷写、修改挂载点权限问题
* 能够保证升级前后用户应用/配置的一致性
  * 单独将 etc、home、opt、usr、var 等目录设置为单独挂载点
* 能够保证升级异常时,检测到异常,并可以回退到系统可用状态
  * AB分区，引导时累计错误次数，错误重启次数异常切换原分区

第二题：升级系统基础框架功能实现 以题目一为基础，升级系统需要追加下列功能
* 实现升级系统管理后台, 主要功能包括:
  * 实现升级包管理,支持用户上传升级包,支持升级包校验
  * 实现升级任务管理,支持用户提交升级任务,支持升级任务状态查询
  * 实现升级日志管理,支持用户升级日志查询
  * 实现用户管理,支持用户登录、注册、权限管理
* 实现升级系统客户端功能,主要功能包括:
  * 支持升级包校验、下载、安装功能，上报升级状态等
  * 支持静默后台强制升级和自动检测升级功能

（可选）第三题：升级模块扩展功能实现 以题目二为基础，升级系统需要追加下列功能：
* 升级功能要有更健壮的异常保护机制,要能兼容电池不足,空间不足,升级中掉电,rootfs/initrd/kernel完全坏掉等情况
  * 掉电引导失败切换另一分区、电池/空间升级前检测并加以限制
* 升级功能要有功能能保证用户应用/配置的兼容性
* 升级功能要能支持uboot和grub两种启动方式
* 升级功能要能支持差分升级,使升级包尽可能小
* 升级系统管理后台要有终端管理,能查询终端的版本及硬件信息
* 升级系统客户端要有安全机制,防止升级后台、升级包等被篡改
  * 生成对称密钥验证、升级包md5检验等

## 技术实现

### shells 升级脚本

* check_img.sh 校验升级镜像文件
* write_image_by_dd.sh 用 dd 写入镜像
* ab_switch.sh 通过修改 `/etc/fstab` 实现AB分区的切换
* check-osfullinfo.sh 检测系统信息
* check-install-sysenv.sh 检测系统依赖
* check-sysupdate.sh 检测系统更新
* sysupdate.sh 系统更新

### ota-updater 升级程序客户端
基于 Golang + Gin 构建，使用 [BadgerDB](https://github.com/dgraph-io/badger) 存储数据，通过 exec 执行 shells 中的命令实现系统级操作。此外除本体守护进程外，也开发了对应管理工具cli应用，用于管理守护进程、执行重置操作。

### client-ui 升级程序客户端界面
基于 Nuxt(Vue3) + NuxtUI 构建升级操作界面，与 ota-updater 交互实现系统升级、升级设置、日志查看等操作。

* 由 ota-updater 启动 HTTP服务器，由根目录输出 client-ui 前端构建的静态文件
* client-ui 通过 fetch 向后端 API 发起请求
* `/login` 登陆页面 基于浏览器 SessionStorage 和 Token 的鉴权-机制
* `/` 主页面 负责查看当前系统/版本信息、版本检查更新、手动更新、更新设置等操作


### ota-manage 暂未完善

## 安装编译
* 构建 client-ui
  * cd clinet-ui/
  * pnpm install
  * pnpm run generate
* 构建 ota-updater (交叉编译请自行修改构建参数)
  * cd ota-updater/
  * go build cmd/daemon/daemon.go
  * go build cmd/cli/sys-upgrade.go
* 构建 ota-manage （基于Docker容器，若手动构建需配置.env文件）
  * docker compose up --build -d

## 操作使用
* 访问 http://127.0.0.1:9301/ 
  * 默认用户密码 passwordroot
* 修改升级服务器为 ota-manage 的API地址 例如: https://cscc.kokomi.ltd/api
* 检查更新/手动上传更新

## **安装要求**
* 适用于分区类型为MBR&GPT

![alt text](image.png)

* 磁盘分区要求: 
  * 一个启动分区(通常为efi)
  * 两个uuid不相同且分区大小相同的分区(文件格式不限,建议f2fs或ext4,使用dd复制的分区须更改为"LABEL=&pathname")
  * 一个数据分区(挂载常用及定制的目录)如:
    * opt ~ 
    * www ~ 是经典的bt面板所在地
    * home ~ 用户文件所在地,是切换系统的同时保留设置的最基本文件夹
    * var ~ 至少tftp serv默认位置在这里
    * snap ~ ubuntu常见


## 以下是webui截图
