1. docker info 查看docker信息
2. docker ps 列出所有正在运行中的容器
3. docker ps -a 列出所有的容器
4. docker images 列出所有的镜像
5. docker rmi imageId 根据镜像id删除镜像
6. docker search xxx 搜索可用的xxx镜像
7. docker pull xxx:version 拉取指定版本的xxx镜像
8. docker build -t xxx:version . 以当前执行命令所在文件夹作为上下文，构建docker镜像 -t 指定target
9. docker exec xxx bash/sh 使用命令行方式进入xxx容器(会进入类似linux系统的目录结构)，对运行状态下的容器有效
10. docker top containerID 查看容器运行进程信息
11. docker history xxx:version
12. docker inspect containerId/imageId(image:version) 查看容器/镜像的配置信息
13. docker image prune 删除所有悬空(既没有被镜像引用，也没有被容器引用的)镜像
14. docker run -d imageName:version 以镜像启动容器，-d代表后台运行
15. docker start/stop -d containerId 启动/关闭容器，-d代表后台运行
16. docker build --add-host hostname:ip 不会对生成的镜像产生影响，但docker run --add-host会