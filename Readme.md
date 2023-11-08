# Redis Serialization Benchmark with JSON vs ProtoBuf

This project is a benchmarking tool for comparing the performance of marshalling and unmarshalling data stored in Redis using JSON and Protocol Buffers (ProtoBuf) formats. It provides insights into the efficiency and speed of data serialization and deserialization techniques when using Redis as a data store.

## Features

- **Serialization Comparison**: Measure and compare the time it takes to serialize data into JSON and ProtoBuf formats.

- **Deserialization Comparison**: Evaluate and compare the time it takes to deserialize data from JSON and ProtoBuf formats.

- **Data Storage**: Store benchmarked data in a Redis database.

- **Simple RESTful API**: Access the benchmarking results and data via a simple RESTful API.

## Prerequisites

- Go installed (v1.16 or later)
- Redis installed and running
- Go Redis client library installed (e.g., [go-redis](https://github.com/go-redis/redis))
- [ProtoBuf](https://developers.google.com/protocol-buffers) compiler (`protoc`) installed
- Go ProtoBuf plugin installed (`go get -u github.com/golang/protobuf/protoc-gen-go`)

<!-- ## Installation -->

<!-- 1. Clone this repository:

```bash
git clone https://github.com/yourusername/redis-serialization-benchmark.git
cd redis-serialization-benchmark -->
