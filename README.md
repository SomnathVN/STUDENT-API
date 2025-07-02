# STUDENT-API

A simple Go-based API for managing student data.

## Features

- Minimal setup and configuration
- Loads configuration from a YAML file
- HTTP server with customizable address

## Getting Started

### Prerequisites

- Go 1.24+ installed on your system

### Installation

Clone the repository:

```sh
git clone https://github.com/SomnathVN/students-api.git
cd students-api
```

### Configuration

Create a configuration file (e.g., `config/local.yaml`). Example:

```yaml
env: "dev"
storage_path: "storage/storage.db"
http_server:
  address: "localhost:8082"
```

### Running the API

You can run the API server with:

```sh
go run cmd/students-api/main.go -config config/local.yaml
```

Or build and run the binary:

```sh
go build -o students-api ./cmd/students-api
./students-api -config config/local.yaml
```

### Environment Variable Alternative

You can also set the config path using the `CONFIG_PATH` environment variable:

```sh
set CONFIG_PATH=config/local.yaml
go run cmd/students-api/main.go
```

### Endpoints

Currently, the only available endpoint is:

- `GET /` — returns a welcome message.

> **Note:** Any other route will return a 404 Not Found.

### Dependencies

Key dependencies (see `go.mod` for full list):

- [github.com/ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv) for config loading
- [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3) for YAML parsing

---

_No sensitive or private information is included in this README._
