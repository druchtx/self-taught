# Tour of Beam (GO) notes

## Concepts

### Runner
> Apache Beam provides a portable API layer for building data-parallel processing pipelines that may be executed across a diversity of execution engines, or runners

**Direct Runner**
> Runner that executes pipelines on your local machine.
> For local development and testing.

[Using the Direct Runner](https://beam.apache.org/documentation/runners/direct/)

**Google Cloud Dataflow Runner**

[Using the Google Cloud Dataflow Runner](https://beam.apache.org/documentation/runners/dataflow/)

**Apache Flink Runner**

[Using the Apache Flink Runner](https://beam.apache.org/documentation/runners/flink/)

**Apache Spark Runner**

[Using the Apache Spark Runner](https://beam.apache.org/documentation/runners/spark/)

**Apache Samza Runner**

[Using the Apache Samza Runner](https://beam.apache.org/documentation/runners/samza/)

### Pipeline

**Pipeline**
> A pipeline encapsulates your entire data processing task, from start to finish.

**PCollection**
> A PCollection represents a distributed data set that your Beam pipeline operates on

**PTransform**
> A PTransform represents a data processing operation, or a step, in your pipeline.

**Scope**


**I/O transform**
>  Beam comes with a number of “IOs” - library PTransforms that read or write data to various external storage **systems**.