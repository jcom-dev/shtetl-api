# Shtetl API - Backend Services

This repository contains all three backend microservices for the Shtetl platform.

## Services

### [Zmanim Service](./zmanim/) - REST: 8101, gRPC: 8102

Zmanim/calendar calculation engine for rabbinic authorities. Handles astronomical calculations, Hebrew calendar logic, and calendar stream publishing.

- **REST API (8101):** Browser access, OpenAPI documentation, third-party integrations
- **gRPC API (8102):** High-performance microservice-to-microservice communication

### [Shul Service](./shul/) - REST: 8103, gRPC: 8104

Shul administration and minyan scheduling. Manages synagogue configuration, minyan scheduling DSL, coverage validation, and PDF generation.

- **REST API (8103):** Browser access, web admin UI, third-party integrations
- **gRPC API (8104):** High-performance microservice-to-microservice communication

### [Kehilla Service](./kehilla/) - REST: 8105

Public-facing community API. Provides schedule queries, subscription management, and notification delivery for congregation members.

## Architecture

- **Language:** Go 1.25.4
- **Databases:** PostgreSQL 18
- **Cache:** Redis 8.4
- **Authentication:** Clerk JWT validation
- **Logging:** zerolog (structured JSON)

## Getting Started

### Prerequisites

- Go 1.25.4
- PostgreSQL 18
- Redis 8.4

### Running All Services Locally

```bash
# Install dependencies for all services
cd zmanim && go mod download && cd ..
cd shul go mod download && cd ..
cd kehilla && go mod download && cd ..

# Run services (in separate terminals)
cd zmanim && go run cmd/zmanim/main.go    # REST: 8101, gRPC: 8102
cd shul && go run cmd/shul/main.go        # REST: 8103, gRPC: 8104
cd kehilla && go run cmd/kehilla/main.go  # REST: 8105
```

### Health Checks

```bash
# Zmanim (REST on 8101)
curl http://localhost:8101/health

# Shul (REST on 8103)
curl http://localhost:8103/health

# Kehilla (REST on 8105)
curl http://localhost:8105/health
```

## Development

Each service follows clean architecture with:

- `cmd/` - Service entry points
- `internal/handler/` - Request handling (gRPC/REST)
- `internal/service/` - Business logic
- `internal/repository/` - Data access
- `api/` - API contracts (proto/OpenAPI)
- `pkg/` - Shared packages

## Service Communication

```
┌─────────────────────┐
│  Zmanim Service     │ REST: 8101, gRPC: 8102
│  - Calculations     │ REST: Browser/OpenAPI access
│  - Calendar Streams │ gRPC: Microservice access
└──────────┬──────────┘
           │ gRPC (service-to-service)
           ↓
┌─────────────────────┐
│  Shul Service       │ REST: 8103, gRPC: 8104
│  - Shul Admin       │ REST: Browser/OpenAPI access
│  - Scheduling       │ gRPC: Microservice access
│  - PDF Generation   │
└──────────┬──────────┘
           │ Reads schedules
           ↓
┌─────────────────────┐
│  Kehilla Service    │ REST: 8105
│  - Public API       │
│  - Subscriptions    │
│  - Notifications    │
└─────────────────────┘
```

## Deployment

Services are deployed as AWS Lambda functions via AWS CDK (see `shtetl-infra` repository).

## Related Documentation

- [Architecture](../../docs/architecture.md)
- [Epics & Stories](../../docs/epics.md)
- [Database Schema](../../docs/architecture/database.md) (Story 1.8)

## License

See [LICENSE](./LICENSE) file.
