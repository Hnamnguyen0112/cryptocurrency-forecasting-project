# Cryptocurrency Data Collector Monorepo

This repository is a monorepo for collecting real-time cryptocurrency data from various exchanges. It includes a collector service built using Go, which listens to WebSocket streams from Coinbase and Binance to gather ticker and candlestick data and store it in a PostgreSQL database.

## Project Structure

The project is structured as follows:

```
collector
├── cmd
│   └── binance_worker
│   └── coinbase_worker
├── config
├── deployments
├── internal
│   └── binance_worker
│   └── coinbase_worker
├── pkg
│   └── database
│   └── entities
│   └── websocket
```

## Features

- **Real-Time Data Collection**: Connects to Coinbase and Binance WebSocket APIs to collect ticker and candlestick data.
- **Data Storage**: Collected data is stored in a PostgreSQL database for further processing and analysis.
- **Monorepo with Bazel**: The project is set up as a monorepo using Bazel to manage dependencies and build configurations.

## Getting Started

### Prerequisites

- **Bazel**: Ensure that Bazel is installed on your system. Follow the [Bazel installation guide](https://bazel.build/install) if necessary.
- **Go**: Install Go (version 1.20 or later) from the [official Go website](https://golang.org/dl/).
- **PostgreSQL**: Set up a PostgreSQL database to store the collected data.

### Installation

1. Clone the repository:
2. Set up the PostgreSQL database and create the necessary tables using the provided SQL scripts.

```bash
docker-compose -f collector/deployments/docker-compose.yml up -d --build
```

3. Build the project using Bazel:

```bash
bazel run //:gazelle
bazel run //:gazelle-update-repos
bazel build //...
```

Note:

```bash
go_repository(
    name = "com_github_confluentinc_confluent_kafka_go_v2",
    importpath = "github.com/confluentinc/confluent-kafka-go/v2",
    patches = ["//bazel:kafka.patch"],
    sum = "h1:icCHutJouWlQREayFwCc7lxDAhws08td+W3/gdqgZts=",
    version = "v2.3.0",
)
```

4. Run the collector service:

```bash
bazel run //collector/cmd/coinbase_worker
bazel run //collector/cmd/binance_worker
```

## Contributing

1. Fork the repository
2. Create a new branch (`git checkout -b feature/branch-name`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/branch-name`)
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
