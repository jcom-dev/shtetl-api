# Shtetl API - Backend Services

This repository contains all three backend microservices for the Shtetl platform.

## Services

### [Zmanim Service](./zmanim/) - REST: 8101

Zmanim/calendar calculation engine for rabbinic authorities. Handles astronomical calculations, Hebrew calendar logic, and calendar stream publishing.

- **REST API (8101):** All access (browser, OpenAPI documentation, service-to-service, third-party integrations)

### [Shul Service](./shul/) - REST: 8103

Shul administration and minyan scheduling. Manages synagogue configuration, minyan scheduling DSL, coverage validation, and PDF generation.

- **REST API (8103):** All access (browser, web admin UI, service-to-service, third-party integrations)

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
cd zmanim && go run cmd/zmanim/main.go    # REST: 8101
cd shul && go run cmd/shul/main.go        # REST: 8103
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

- `cmd/lambda/` - Lambda deployment entry point
- `cmd/server/` - HTTP server entry point (local dev)
- `internal/handler/` - REST request handling
- `internal/service/` - Business logic
- `internal/repository/` - Data access
- `api/` - API contracts (OpenAPI 3.1)
- `pkg/` - Shared packages

## Service Communication

```
┌─────────────────────┐
│  Zmanim Service     │ REST: 8101
│  - Calculations     │ All access via REST API
│  - Calendar Streams │
└──────────┬──────────┘
           │ REST API (service-to-service)
           ↓
┌─────────────────────┐
│  Shul Service       │ REST: 8103
│  - Shul Admin       │ All access via REST API
│  - Scheduling       │
│  - PDF Generation   │
└──────────┬──────────┘
           │ Reads schedules (database)
           ↓
┌─────────────────────┐
│  Kehilla Service    │ REST: 8105
│  - Public API       │ All access via REST API
│  - Subscriptions    │
│  - Notifications    │
└─────────────────────┘
```

## Deployment

Services are deployed as AWS Lambda functions via CDKTF (Terraform + TypeScript, see `shtetl-infra` repository).

## Related Documentation

- [Architecture](../../docs/architecture.md)
- [Epics & Stories](../../docs/epics.md)
- [Database Schema](../../docs/architecture/database.md) (Story 1.8)

## License

See [LICENSE](./LICENSE) file.
