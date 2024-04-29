# OpenTelemetry是什么
1. 针对微服务下，分布式链路追踪技术的标准
2. 主要包含数据格式、系统指标等
3. java的zipkin，golang的jaeger等分布式链路追踪
4. server、agent、dbStore、go-client
5. opentracing统一了zipkin、jaeger等技术的指标标准
6. OpenTelemetry融合了两个标准，成为了新的标准。opentracing+OpenCensus = OpenTelemetry
# OpenTelemetry能做什么
1. 对每种语言都提供了相应的lib库
2. 提供了中立的数据采集器
3. 可以生成、采集、发送、处理和导出遥测数据
4. 可以通过配置将数据并行发送到多个目的地
5. 提供了Opentracing和OpenCensus的垫片，可以让使用了这两种标准的项目快速转换到OpenTelemetry
# 什么是链路追踪
1. traces: 链路。请求的调用链。同一链路的所有span具有统一的trace_id
2. span: 跨度。一个完整的逻辑，是一个span。简单理解为一个方法就是一个span。一个trace包含多个span，
   每个span又可以保存属性、时间、事件等。
3. 方便在微服务中快速定位错误发生的服务、方法。
# metrics: 指标/度量
1. counter: 一个随时间累加的值，比如年龄、里程等。可以计算某个时间内的累计值以及累计值与时间的比率。
2. measure: 随时间聚合的值。根据某种分类，计算时间段内各个分类的聚合值
3. observer: 捕获特定时间点的当前值，可以立即观测到的值。比如时间、cpu瞬时使用率等。
# 基于OTel标准的Jaeger和zipkin接入
# 基于OTel标准接入prometheus
1. prometheus、metrics都是监控
2. OpenCensus统一了监控指标的标准
# baggage统一数据存储与传播



docker run -d --name jaeger \
-e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
-e COLLECTOR_OTLP_ENABLE=true \
-p 5775:5775/udp \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 16686:16686 \
-p 14250:14250 \
-p 14268:14268 \
-p 14269:14269 \
-p 9411:9411 \
-p 4317:4317 \
-p 4318:4318 \
jaegertracing/all-in-one:latest

sudo yum-config-manager \
--add-repo \

sudo yum-config-manager \
--add-repo \
https://download.docker.com/linux/centos/docker-ce.repo
