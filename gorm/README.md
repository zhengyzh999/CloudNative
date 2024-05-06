# gorm 四个优势与四个弊端如何选择
# 选择orm的理由
1. 规范一致，代码工整
2. 减少一定的工作量
3. 对于一些通用系统，部署更方便
4. 解耦数据库与数据访问层，方便更换数据库引擎

# 不建议使用orm的理由
1. 数据访问层不会因为使用orm而比原生sql显著减小
2. 大量使用反射，导致程序性能不佳
3. 一个没有sql基础的开发人员，大概率不能通过orm构建正确的sql语句
4. orm提供了大量表关系接口，主要是表连接操作，数据量大的情况下，会导致查询性能显著下降

# 数据库连接串格式
1. "username:password@tcp(ip:port)/databaseName?charset=utf8mb4&parseTime=True&loc=Local"

# 入库操作
1. 指定表或者model
2. 正向选择或反向选择哪些字段入库
3. 单条记录还是多条记录，批量可以指定每次批处理的条数

# 查询操作
1. 指定表或者model
2. 正向或者反向选择查询字段
3. where子句构建(and,or,in,not)
4. order by子句
5. group having子句
6. limit offset子句
7. 查询条目，一条或多条
8. 结果填充，使用对象、集合、切片等接收查询结果

# 更新操作
1. 指定表或者model
2. 正向选择或反向选择更新字段
3. where子句

# 删除操作
1. 指定表或者model
2. where子句

# 事务
1. 普通事务
2. 嵌套事务
3. 手动事务
4. savepoint与rollback