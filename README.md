# MashinChi 
> **Find it. Drive it.**

**A Go-powered backend for a car marketplace, built around clean architecture, dynamic search, and production-oriented infrastructure.**

MashinChi is a RESTful API project for a vehicle marketplace, built with Go and organized around Clean Architecture principles. It brings together a layered application structure, relational data storage, caching, authentication, flexible data querying, and an observability stack in a Docker-based development environment.

The project is intended as a backend foundation for a platform where users can discover and manage car listings. Its technical design keeps HTTP delivery, business logic, and data access separated so that the application can evolve without coupling every part of the system to the web framework or database implementation.

## Features

- **RESTful API:** HTTP endpoints implemented with the Gin framework.
- **Authentication and authorization:** JWT-based protection for authenticated API operations.
- **Relational persistence:** PostgreSQL with GORM for database access.
- **Caching:** Redis integration for cache-backed operations.
- **Flexible querying:** Dynamic filtering, sorting, and pagination support for list-style API queries.
- **Input validation:** Request validation with `go-playground/validator`.
- **API documentation:** Swagger/OpenAPI documentation generated with Swaggo.
- **Configuration management:** YAML-based application configuration managed with Viper.
- **Structured logging:** Zap and Zerolog integrations, with Filebeat forwarding application logs to Elasticsearch.
- **Monitoring and dashboards:** Prometheus metrics with Grafana dashboards.
- **Containerized environment:** Docker Compose configuration for the API and its supporting services.

> **Scope note:** This README describes the architecture and capabilities documented in the upstream codebase. For the exact vehicle, user, and lookup operations available in your current version, refer to the registered routes and generated Swagger documentation. Features such as online payments, in-app messaging, or image hosting should only be listed after they have been implemented.

## Technology Stack

| Area | Technology | Purpose |
|---|---|---|
| Language | Go 1.22 | Backend implementation |
| HTTP framework | Gin | Routing, handlers, and middleware |
| Database | PostgreSQL | Primary relational data store |
| ORM | GORM | Database access and model mapping |
| Cache | Redis | Cache storage |
| Authentication | JWT | Authentication and authorization |
| Validation | go-playground/validator | Request validation |
| Configuration | Viper | Application configuration |
| API documentation | Swaggo / Swagger | API specification and interactive documentation |
| Logging | Zap, Zerolog | Structured application logs |
| Log shipping and search | Filebeat, Elasticsearch | Log collection and storage |
| Log visualization | Kibana | Log exploration and analysis |
| Metrics | Prometheus | Metrics collection and storage |
| Dashboards | Grafana | Monitoring dashboards |
| Infrastructure | Docker, Docker Compose | Local development environment |
| Database administration | pgAdmin | PostgreSQL administration |

## Architecture

The codebase follows a layered design inspired by Clean Architecture. Responsibilities are separated so that HTTP concerns stay in the API layer, business rules remain in the service layer, and persistence concerns are handled by the data layer.

```text
                  Clients
                     |
                     v
              Gin HTTP API
       (Routes, Handlers, Middleware, DTOs)
                     |
                     v
               Service Layer
              (Business Logic)
                     |
                     v
                Data Layer
           /                 \
          v                   v
     PostgreSQL              Redis
       (GORM)                (Cache)
```

Supporting components provide configuration, validation, API documentation, logging, and metrics around the application.

### Observability

```text
Application Logs --> Filebeat --> Elasticsearch --> Kibana

Application Metrics --> Prometheus --> Grafana
```

The Docker Compose environment also includes Node Exporter and Alertmanager for infrastructure metrics and alert handling.

## Project Structure

The repository is organized by responsibility. The main application areas are:

| Area | Responsibility |
|---|---|
| `api` | HTTP server setup, DTOs, routes, handlers, and middleware |
| `services` | Application and business logic |
| `data` | PostgreSQL and Redis access |
| `cmd` | Application entry point and startup wiring |
| `config` | Environment-specific application configuration |
| `common` | Shared utilities |
| `constants` | Application constants |
| `docs` | Swagger documentation and related assets |
| `pkg` | Reusable supporting packages |
| `test` | Test code |
| `docker` | Docker Compose and infrastructure configuration |

> Directory names above describe the logical modules documented by the project. Check the current repository tree for the exact paths in your working copy.

## Getting Started

### Prerequisites

- Go 1.22 or later
- Docker Engine
- Docker Compose plugin
- Git

### 1. Clone the repository

```bash
git clone <YOUR-MASHINCHI-REPOSITORY-URL>
cd mashinchi
```

Replace the clone URL with the URL of your repository.

### 2. Start the complete environment with Docker Compose

From the repository root, run:

```bash
docker compose -f docker/docker-compose.yml up -d --build
```

This starts the API containers and the supporting services defined in the Compose file.

### 3. Access the services

The Compose configuration exposes the following local endpoints:

| Service | URL / Port |
|---|---|
| MashinChi API instance 1 | [http://localhost:9002](http://localhost:9002) |
| MashinChi API instance 2 | [http://localhost:9003](http://localhost:9003) |
| Elasticsearch | [http://localhost:9200](http://localhost:9200) |
| Kibana | [http://localhost:5601](http://localhost:5601) |
| Prometheus | [http://localhost:9090](http://localhost:9090) |
| Grafana | [http://localhost:3000](http://localhost:3000) |
| Alertmanager | [http://localhost:9093](http://localhost:9093) |
| Node Exporter | [http://localhost:9100](http://localhost:9100) |
| pgAdmin | [http://localhost:8090](http://localhost:8090) |
| PostgreSQL | `localhost:5432` |
| Redis | `localhost:6379` |

The API is also configured to run locally on port `5005` when started outside Docker, according to the development configuration documented by the upstream project.

### 4. Run the API locally (optional)

To run the API process on your machine while using Docker for its dependencies, start the supporting services first:

```bash
docker compose -f docker/docker-compose.yml up -d \
  setup elasticsearch kibana filebeat postgres pgadmin redis \
  prometheus node-exporter alertmanager grafana
```

Then start the Go application:

```bash
cd src
go mod download
go run ./cmd/main.go
```

The local API is expected at:

```text
http://localhost:5005
```

If your working copy has a different entry-point path or configuration, use the paths and settings in that copy.

### 5. Stop the environment

```bash
docker compose -f docker/docker-compose.yml down
```

To also remove persistent Docker volumes and the data stored in them, use the following command only when you intentionally want to reset the local environment:

```bash
docker compose -f docker/docker-compose.yml down -v
```

## API

The API uses versioned routes under `/api/v1`. The upstream example documents username-based login at:

```http
POST /api/v1/users/login-by-username
Content-Type: application/json
```

Example request:

```json
{
  "username": "YOUR_USERNAME",
  "password": "YOUR_PASSWORD"
}
```

A successful login returns an authentication response containing a JWT. Use the token according to the authentication requirements of protected endpoints.

### Dynamic filtering, sorting, and pagination

The project includes a flexible query structure for list endpoints. An example filter request is:

```json
{
  "filter": {
    "Name": {
      "filterType": "text",
      "from": "sedan",
      "type": "contains"
    }
  },
  "pageNumber": 1,
  "pageSize": 10,
  "sort": [
    {
      "colId": "name",
      "sort": "desc"
    }
  ]
}
```

The accepted filter fields and operators depend on the endpoint. Use the Swagger specification for the exact request and response schemas supported by your current version.

## Configuration

Application configuration is managed with Viper and is organized into environment-specific YAML configuration files. The deployment setup also uses `APP_ENV` to select the production environment.

Before running the project, review the configuration files under `src/config` and make sure that the database, Redis, logging, and other service settings match your local environment.

Do not commit real passwords, API keys, JWT secrets, or other credentials. Use local development values for local work and inject production secrets through a secure configuration mechanism.

## Development Notes

- Keep HTTP-specific logic in the API layer and business rules in services.
- Keep database and cache operations behind the data layer instead of accessing them directly from handlers.
- Validate incoming requests before passing data to business logic.
- Update the Swagger annotations whenever an endpoint or DTO changes.
- Add tests for new service and data-layer behavior.
- Keep observability useful: log meaningful context, avoid sensitive data, and expose metrics for important application behavior.

## Security Notice

The upstream Docker Compose configuration contains development-oriented default credentials and settings. Review and replace these values before exposing services outside your local machine or deploying the project. In particular, secure PostgreSQL, pgAdmin, Grafana, Elasticsearch, Kibana, Redis, and all application secrets.

This repository is a development foundation, not a production deployment blueprint. Production use requires a security review, secret management, access restrictions, TLS, backup and recovery procedures, and appropriate resource limits.

---

<p align="center">
  Built with Go for a better car marketplace experience.
</p>
