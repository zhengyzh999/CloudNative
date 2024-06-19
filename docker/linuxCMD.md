1. cat /proc/mounts 查看主机挂载信息
2. cat /proc/xxxPID/mounts 查看xxx进程挂载信息
3. lsns -p xxxPID 查看xxx进程的命名空间信息
4. ls /proc/xxxPID/ns/ -al 查看xxx进程的命名空间信息
5. id xxx 查看xxx的uid、gid和groups信息
6. hostname 查看主机名
7. domainname 查看域名
8. mount -t cgroup 查看系统所有控制组
9. scp sourceFile targetPath linux间互传文件，可以上传和下载。source和target是相对的
10. 