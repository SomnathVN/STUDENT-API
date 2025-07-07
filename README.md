# STUDENT-API

A simple Go-based API for managing student data.

## Features

- Minimal setup and configuration
- Loads configuration from a YAML file
- HTTP server with customizable address
- **Supports both SQLite and Firestore as storage backends**
- **API key authentication middleware**
- **Rate limiting middleware (60 requests/minute per IP)**
- **Request logging middleware (logs remote IP, method, and path)**

## Getting Started

### Prerequisites

- Go 1.24+ installed on your system
- (For Firestore) Google Cloud project and credentials
- (Optional) [ngrok](https://ngrok.com/) for public API exposure

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
storage_path: "storage/storage.db" # Used for SQLite
project_id: "your-gcp-project-id" # Used for Firestore
api_key: "your-secret-api-key" # Used for API key authentication
http_server:
  address: "localhost:8082"
```

- `storage_path`: Path to SQLite database file (if using SQLite backend)
- `project_id`: Google Cloud project ID (required for Firestore backend)
- `api_key`: API key required in the `X-API-Key` header for all requests
- `http_server.address`: Address to bind the HTTP server

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

### Selecting Storage Backend

- By default, the code uses Firestore as the storage backend. To use SQLite, uncomment the relevant lines in `cmd/students-api/main.go` and comment out Firestore lines.
- Ensure your config file has the correct fields for your chosen backend.

### Security Middleware

- **API Key Auth:** All requests must include a valid `X-API-Key` header matching the value in your config.
- **Rate Limiting:** Each IP is limited to 60 requests per minute. Exceeding this returns HTTP 429.
- **Logging:** All requests are logged with remote IP, method, and path.

### Exposing API with ngrok

To make your API publicly accessible:

1. [Download ngrok](https://ngrok.com/download) and set up your auth token.
2. Start your API server locally.
3. In a new terminal, run:
   ```sh
   ngrok http 8082
   ```
4. Use the HTTPS forwarding URL provided by ngrok for external access.

### Android/Retrofit Integration

- Use the HTTPS ngrok URL as your API base URL in your Android app.
- Add the `X-API-Key` header to all requests.
- Example Retrofit setup:

```java
OkHttpClient client = new OkHttpClient.Builder()
    .addInterceptor(chain -> chain.proceed(
        chain.request().newBuilder()
            .header("X-API-Key", "your-secret-api-key")
            .build()
    ))
    .build();

Retrofit retrofit = new Retrofit.Builder()
    .baseUrl("https://your-ngrok-url/")
    .client(client)
    .addConverterFactory(GsonConverterFactory.create())
    .build();
```

### Endpoints

- `POST /api/students` - Create a new student
- `GET /api/students/{id}` — Get a student by ID
- `GET /api/students` — List all students
- `PUT /api/students/{id}` - Update a student by ID
- `DELETE /api/students/{id}` - Delete a student by ID

- `POST /` - creates student in students table.
- `GET /` — returns student by id and list of student.
- `UPDATE /` - updates student in students table at given id.
- `DELETE /` - deletes student in students table at given id. 
=======
>>>>>>> b88813d13ecf4d7dff5583ab2ebab861910b7f4f

> **Note:** Any other route will return a 404 Not Found.

### Dependencies

Key dependencies (see `go.mod` for full list):

- [github.com/ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv) for config loading
- [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3) for YAML parsing
- [cloud.google.com/go/firestore](https://pkg.go.dev/cloud.google.com/go/firestore) for Firestore backend
- [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) for SQLite backend

---
