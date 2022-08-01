# go-oss-demo
学习《分布式对象存储——原理、架构及Go语言实现》的项目

## 第一章完成

第一章只实现了很简单的一个上传下载功能，代码也很少。

目前只是一个单机版本的实现，下一章会将接口与数据存储解耦合以支持后续的分布式实现。

[第一章笔记](https://note.yahui.tech/notes/object-storage-golang-book/start.html)

## 第二章完成

提供了 `play.sh` 脚本

脚本主要做了两件事：
- 将 apiServer 和 dataServer 编译成可执行文件（默认是 linux 的，其它 os 自行修改环境变量属性）
- 运行 docker compose

docker compose 会首先尝试 build apiServer 和 dataServer 的镜像，然后运行1个 apiServer 实例和5个 dataServer 实例，数量可以在 `paly.sh` 中修改。

除此之外，我们需要预先在 rabbit mq 中创建好各自的 exchange，这是在代码中硬编码了的，所以不能自定义。分别是 `apiServers` 和 `dataServers`，Type 设置为 fanout，其它属性默认即可。

第二章基本仍是按照书上的设计来完成的，不过进行了以下扩展：
- 采用 docker 部署
- 添加了获取本地 ip 的功能，不用再去环境变量里设置 ip 地址，这样子可以方便部署多台实例

[第二章笔记](https://note.yahui.tech/notes/object-storage-golang-book/ch02.html)