# docker容器编排神器compose应用

# compose的两个关键概念
1. 项目，一个项目包含多个服务
2. 服务，即运行应用程序的容器

# docker-compose配置文件
1. 默认配置文件名 docker-compose.yml

# compose命令行操作

# docker-compose常用命令
1. docker-compose命令一般在docker-compose.yml文件所在目录下执行，表示当前目录下创建镜像和启动容器
2. docker-compose config 检查当前目录下的docker-compose.yml配置文件是否有错误
3. docker-compose up -d 启动项目 -d代表后台启动
4. docker-compose restart 重启项目
5. docker-compose down 停止项目
6. docker-compose build --no-cache 编译项目。 --no-cache表示不使用缓存
7. docker-compose logs serviceName 查看某个服务的日志
8. docker-compose exec -it serviceName bash 以终端方式进入某个服务
9. docker-compose top 查看每个service的进程信息