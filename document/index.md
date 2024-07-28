# Linux的OTA升级系统 文档

## 基础信息
大连理工大学城市学院 T202413198993225 地铁行动2014 proj235 Linux的OTA升级系统

Github仓库：https://github.com/lixworth/2024OS-Proj235 <br>
Gitlab镜像仓库：https://gitlab.eduxiji.net/T202413198993225/project2210132-239709

## 需求方案分析
> https://github.com/oscomp/proj235-linux-upgrade-system

第一题：升级系统的升级功能实现
* 单独将 home、opt、usr、var 等与应用配置相关的目录单独设置挂载点
* 建立AB分区，升级过程中，通过修改 `/etc/fstab` 实现切换分区
* 更改挂载点、使用dd刷写 （initrd,kernel,rootfs）镜像
* 建立默认启动应用，检测重启次数，超出限制回滚回另一系统并标记

第二题：升级系统基础框架功能实现 以题目一为基础，升级系统需要追加下列功能
* sys-update cli应用
  * sys-update switch_ab [arg]
  * sys-update update [arg.img]
  * sys-update check 
  * sys-update reset-pwd
* webclient web交互界面
  * 升级包上传、校验、执行升级
  * 日志查询
  * 升级任务管理、升级队列状态查询

（可选）第三题：升级模块扩展功能实现 以题目二为基础，升级系统需要追加下列功能：
* 掉电引导失败切换另一分区、电池/空间升级前检测并加以限制
* 生成对称密钥验证、md5检验升级包等
* 版本与终端管理平台，存放升级包，客户端监测api查询最新版本

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

### ota-manage 
基于 Hyperf + ArcoDesignPro 构建的前后端版本与终端管理平台。

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

![截屏2024-07-15%2022.03.59](./screenshot/截屏2024-07-15%2022.03.59.png)
![截屏2024-07-28%2013.26.37](./screenshot/截屏2024-07-28%2013.26.37.png)
