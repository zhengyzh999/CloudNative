# 封装logrus
1. 定义一个interface
2. 定义三个实现interface的writer: std、file、file rotate
3. 应用程序启动时，选择初始化某一个writer