# 基于go-kit工具集理解微服务开发
1. auth鉴权 :go-kit提供basic,casbin,jwt三种鉴权
2. circuitbreaker熔断器/断路器
3. endpoint
4. log日志打印
5. metrics服务指标采集/系统监控
6. ratelimit限流
7. sd服务注册和发现
8. tracing链路追踪: opencensus,opentracing(两者合并为opentelemetry),zipkin
9. transport传输协议

# 微服务需要哪些功能
1. 鉴权
2. 服务熔断
3. 日志打印
4. 服务指标采集/系统监控
5. 限流
6. 服务注册和发现(云原生不需要，有容器提供注册和发现)
7. 链路追踪
8. 传输协议


# 基于中间件编程
# 服务的限流与熔断
# 服务链路追踪及透传(指跨服务的操作，需要透传)