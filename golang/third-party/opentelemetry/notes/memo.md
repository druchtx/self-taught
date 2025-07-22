# Opentelemetry

OpenTelemetry 专注于遥测数据(信号)的生成、采集、管理和导出,搭配包括 Jaeger 和 Prometheus 这类可观测后端使用。

## 基本概念

Opentelemtry支持的遥测数据类别:

- 日志 Log : 日志是带有时间戳的文本记录，可以是结构化(CLF/ELF)的或非结构化的，且可附带元数据
- 指标 Metric: 
- 链路 Trace: 
- 行李 Baggage: 
- 事件 Event:(正在开发或处于提案阶段)
- 性能分析数据 Profiling:(正在开发或处于提案阶段)

### 日志 Log

OpenTelemetry 没有API 或 SDK 来创建日志。日志来自现有的日志框架或基础设施组件。OpenTelemetry SDK 和自动注入功能利用多个组件， 自动将日志与链路关联。
OpenTelemetry Collector: 提供 收集、处理、转换和导出日志的工具

#### 日志核心组件

- Logs Bridge API: 为日志库作者构建日志附加器或桥接器而设
- Logger Provider: 属于Logs Bridge API，是Logger工厂
- Logger: 日志记录器 ,实例由Logger Provider创建
- Log Record: 日志记录,代表对事件的记录,包含2类字段:
  - 具备特定含义的命名顶级字段
  - 可任意定义的 Resource 和 Attributes 字段

### 指标 Metric



### 链路 Trace

### 行李 Baggage
