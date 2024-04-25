# go test的两种模式
1. 本地模式 go test (-v) : 当前目录下进行测试
2. 列表模式 go test dir1 dir2 : 测试目录列表下的文件

# go test常用的flag
1. -bench regexp: 接收一个正则表达式, 运行与正则匹配的测试用例
2. -benchtime t: 接收一个指定时间或者次数。表示进行测试的时间或者次数。1s 1m 100x
3. -count n: 接收一个次数。表示运行go test测试的次数
4. -cover: 表示启用覆盖率分析
5. -cpu 1 2 4: 表示启用多少个cpu进行go test，其中runtime.GOMAXPROCS分别是1，2，4时，个执行一次，共执行3次
6. -fuzz regexp: 接收一个正则表达式。运行匹配的模糊测试用例。只能匹配一个用例
7. -fuzztime t: 接收一个模糊测试时长。默认一直执行
8. -list regexp: 接收一个正则表达式。列出匹配的顶层测试用例
9. -parallel -n: 接收一个cpu数。指定性能测试时并行的cpu数量。
10. -run regexp: 接收一个正则表达式。运行匹配的功能测试用例。
11. -short: 表示缩短运行时间。short情况下忽略被short标记(程序员在用例上自己指定)的测试用例
12. -timeout d: 单次go test的超时时间。默认为10min，设置为0表示禁用
13. -v: 打印所有输出。在本地模式下，有无-v都会打印用例中的输出。列表模式下，添加-v会打印用例中的输出。
14. -benchmem: 打印性能测试的内存分配统计信息
15. -blockprofile block.out: 接收一个文件名。导出性能测试的阻塞数据到指定文件
16. -coverprofile cover.out: 接收一个文件名。导出性能测试的覆盖率数据到指定文件
17. -cpuprofile cpu.out: 接收一个文件名。导出性能测试cpu使用数据到指定文件
18. -memprofile mem.out: 接收一个文件名。导出性能测试内存数据到指定文件
19. -mutexprofile mutex.out: 接收一个文件名。导出性能测试互斥锁数据到指定文件
20. -outputdir dir: 接收一个文件夹名。导出性能测试文件到指定文件夹
21. -trace trace.out: 接收一个文件名。导出性能测试执行跟踪信息到指定文件