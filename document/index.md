# Linux的OTA升级系统 文档

## 技术实现
### shells 升级脚本


### ota-updater 升级程序本体
基于 Golang + Gin 构建，使用 BadgerDB 存储数据，通过 exec 执行 shells 中的命令实现系统级操作，除本体守护进程外，也开发了对应管理工具cli应用。


### client-ui 升级程序界面
基于 Nuxt(Vue3) + NuxtUI 构建升级操作界面，与 ota-updater 交互实现系统升级、升级设置、日志查看等操作。

## 安装编译

* 构建 client-ui
  * cd clinet-ui/
  * pnpm install
  * pnpm run generate
* 构建 ota-updater (交叉编译请自行修改构建参数)
  * cd ota-updater
  * go build cmd/daemon/daemon.go
  * go build cmd/cli/sys-upgrade.go
* 构建 ota-manage （基础Docker容器，手动构建需配置.env文件）
  * docker compose up --build -d

## 操作使用
* 访问 http://ip:9301/ 
  * 默认用户密码 passwordroot
* 修改升级服务器为 ota-manage 的API地址 例如: https://cscc.kokomi.ltd/api
* 检查更新/手动上传更新
