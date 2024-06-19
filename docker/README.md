# docker是如何解决应用程序的环境依赖问题的
1. 为什么选择docker？
    docker解决了应用程序对环境的依赖问题
2. docker怎么解决应用程序对环境的依赖
    docker镜像，将应用程序与应用程序所依赖的文件一起打包。
    基于docker运行容器

# docker镜像基本原理
1. 分层: 将多个镜像共用的资源抽取成一份
    缩小容器或者镜像所占用的磁盘空间
    缩小镜像传输的数据流
    容器启动更快速
    只读层(共享层)文件修改问题，某个镜像若是修改，可能造成其他镜像引用出错
2. 写时复制
    如果某镜像需要对共享层文件进行修改，会将要修改的文件copy到各自的执行空间内，进行修改，以后使用该修改后的文件。不影响其他镜像对共享层的引用
3. 联合挂载机制
    将多个镜像层内的文件联合挂载到一个挂载点，使容器拥有完整的根目录系统

# dockerfile定义镜像构建步骤
1. 两种镜像构建方式
    命令行docker commit。将一个容器提交为镜像
    docker build + dockerfile 构建镜像
2. 两类镜像
    软件镜像，镜像运行之后可以提供服务：mysql、nginx、redis、xxxServer
    纯文件系统镜像，为软件镜像提供环境上的依赖

# docker build构建流程
1. 将上下文中的文件发送到docker守护线程daemon
2. docker build向docker server发送请求，将上下文信息发送到server
3. docker server接收到请求后，执行构建流程
    创建临时目录，解压上下文
    读取dockerfile，遍历dockerfile中的指令
    为每条指令创建一个临时容器，在容器中执行指令，再提交为镜像
    将所有指令构建出的镜像层合并，形成最后的构建结果
4. 构建过程中基于dockerfile，重复构建未改动dockerfile的项目会使用cache

# docker所涉及的内核知识
1. linux namespace: 相当于四面墙，提供独立的空间
2. control groups: 相当于房顶，房子容器过度使用资源
3. rootfs: 相当于房子的地板，文件系统

# 容器: 解决应用部署
1. 解决了应用所依赖的环境问题
2. 解决了应用程序资源使用问题
3. 有利于快速扩展和弹性伸缩
4. docker可以运行多个应用程序，只能指定一个程序入口。原则上只部署一个应用程序

# docker容器如何限制CPU和内存的使用
1. docker容器与linux命名空间
   1. cgroup: cgroup根目录
   2. pid: 进程pid
   3. user: 隔离用户和用户组
   4. uts: 隔离hostname(主机名)和domainname(域名)
   5. ipc: 隔离进程间通讯资源
   6. mnt: 挂载点
   7. net: 网络栈
2. pid为1的为源主进程
3. 当进程没有指定某个命名空间时，沿用父进程的命名空间
4. unshare --fork -m -u -i -n -p -U -C sleep 100 fork一个进程
5. docker 启动命令添加参数取消命名空间隔离 docker run --containerName -d --userns=host --pid=host xxx
6. docker容器与控制组
7. 使用命名空间和控制组限制容器能力

# 容器和镜像的配置文件
1. LowerDir: 被引用的所有镜像层，只读
2. UpperDir: 容器启动之后创建的读写层
3. MergedDir: 容器启动之后，将LowerDir和UpperDir所有条目，联合挂载之后的完整根目录
4. WorkDir:

# docker注册中心是如何分发镜像的
1. docker镜像基本组成原理
2. 12个镜像管理及镜像仓库命令
3. 搭建自己的私有注册中心

# docker常见概念
1. registry注册中心
2. repository仓库。一个软件就是一个仓库，通常某个仓库包含该软件的所有迭代版本
3. manifest镜像元数据文件
4. image镜像。一个镜像存储了一些相关信息
5. layer镜像层
6. dockerfile