# Shtetl API - Backend Services

This repository contains all three backend microservices for the Shtetl platform.

## Services

### [Zmanim Service](./zmanim/) - Port 8001 (gRPC)

Zmanim/calendar calculation engine for rabbinic authorities. Handles astronomical calculations, Hebrew calendar logic, and calendar stream publishing.

### [Shul Service](./shul/) - Port 8002 (REST)

Shul administration and minyan scheduling. Manages synagogue configuration, minyan scheduling DSL, coverage validation, and PDF generation.

### [Kehilla Service](./kehilla/) - Port 8003 (REST)

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
cd zmanim && go run cmd/zmanim/main.go    # Port 8001
cd shul && go run cmd/shul/main.go        # Port 8002
cd kehilla && go run cmd/kehilla/main.go  # Port 8003
```

### Health Checks

```bash
# Zmanim (gRPC - health endpoint added in Story 1.11)
# Shul
curl http://localhost:8002/health

# Keh illa
curl http://localhost:8003/health
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
│  Zmanim Service     │ Port 8001 (gRPC)
│  - Calculations     │
│  - Calendar Streams │
└──────────┬──────────┘
           │ gRPC
           ↓
┌─────────────────────┐
│  Shul Service       │ Port 8002 (REST + gRPC)
│  - Shul Admin       │
│  - Scheduling       │
│  - PDF Generation   │
└──────────┬──────────┘
           │ Reads schedules
           ↓
┌─────────────────────┐
│  Kehilla Service    │ Port 8003 (REST)
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
