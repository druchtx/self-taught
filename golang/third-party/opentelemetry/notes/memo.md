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

metric是在运行时捕获的服务的测量值。捕获测量值的时刻称为metric事件，它不仅包括测量值本身，还包括捕获它的时间和相关的元数据。

#### 指标重要概念

- Meter Provider: 是 Meter 的工厂,一般只在程序启动时初始化一次
- Meter: 是 Metric Provider 的实例，用于创建 metrics instruments,在运行时捕获服务测量量
- Meteric Exporter: 将指标数据导出到后端
- Metrics Instruments: 用于捕获服务测量量,由以下定义组成:
  - 名称
  - 类型 : 
    - Counter
    - Asynchronous Counter
    - UpDownCounter
    - Asynchronous UpDownCounter
    - Gauge
    - Asynchronous Gauge
    - Histogram
  - 单位(可选)
  - 描述(可选)
  
- Aggregation 聚合: 将测量值组合成在一个时间窗口内发生的指标事件的精确或估计统计数据,OpenTelemetry 项目旨在提供可视化工具和遥测后端支持的默认聚合
- View 视图: 可自定义 SDK 输出的指标,自定义要处理或忽略的度量仪器,可以自定义聚合和希望在指标上报告的属性。

### 链路 Trace

链路为我们提供了向应用发出请求时会发生什么的总览图,了解程序的完整的路径

#### 链路重要概念

- Trace Provider: 是 Trace 的工厂，一般只在程序启动时初始化一次
- Tracer: 用于创建Span,包含相关给定操作
- Tracer Exporter: 将链路数据导出到后端
- Context Propagation 上下文传播: 用于跨进程/服务传递链路信息,将多个span关联组合成链路
- Span: 表示工作或操作单元,是链路的构建块,主要由以下信息构成
  - 名称
  - 父SpanID
  - 开始结束时间戳
  - Span上下文: traceid,spanid,traceflags,tracestate
  - 属性: 包含元数据的键值对
  - Span事件: 持续时间内有意义的单个时间点
  - Span连接: 将一个Span与一个或多个Span相关联，从而暗示因果关系,常用于异步,队列等无法知道何时开始的情况,或者多个相关服务产生的后续操作?
  - Span状态:
    - Unset: 默认值，表示它跟踪的操作已成功完成，没有错误
    - OK: 手动标记为OK时使用，一般不用
    - Error: 意味着它跟踪的操作中发生了一些错误
  
  Span类型：
  
  - Client
  - Server
  - Internal
  - Producer
  - Consumer

### 行李 Baggage

是一个键值对存储，用于span之间传递上下文信息。
允许你在服务和进程之间传递数据， 从而可以将其添加到这些服务中的链路、指标或日志中。
Baggage 最适合用于将请求开始时通常可用的信息，传递到后续处理流程中。 这些信息可以包括账户标识、用户 ID、产品 ID 和来源 IP 等。

## 插桩

要使系统具备可观测性，就必须进行插桩（Instrumentation）： 也就是说，系统组件中的代码必须发出信号，如链路、指标和日志。

- 基于代码的插桩(Code-Based)
- 零代码插桩(Zero-Code)

### 零代码插桩 

golang使用eBPF方案进行零代码插桩

**相关资源**
- [Repository](https://github.com/open-telemetry/opentelemetry-go-instrumentation)
- [Getting Started with OpenTelemetry Go Automatic Instrumentation](https://github.com/open-telemetry/opentelemetry-go-instrumentation/blob/main/docs/getting-started.md)
- [How it works](https://github.com/open-telemetry/opentelemetry-go-instrumentation/blob/main/docs/how-it-works.md)
- [Tutorial](https://github.com/open-telemetry/opentelemetry-go-instrumentation/tree/main/docs/tutorial)


