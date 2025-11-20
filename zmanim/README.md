# Zmanim Service

**Purpose:** Zmanim/calendar calculation engine for rabbinic authorities

## Overview

The Zmanim Service provides astronomical and Hebrew calendar calculations for the Shtetl platform. It exposes a REST API for publishing and consuming calendar streams with multiple halachic calculation methodologies.

## Architecture

- **Technology:** Go 1.25.4 + REST (Lambda-compatible)
- **Port:** 8101 (REST)
- **Database:** PostgreSQL (calculation formulas, streams, versions)

## Key Capabilities

- Zmanim DSL parsing and execution
- Astronomical calculations (alot, netz, shkiah, etc.)
- Hebrew calendar calculations (holidays, fast days, Rosh Chodesh)
- Calendar stream publishing with version control
- Audit trails for calculation changes (7-year retention)

## Directory Structure

```
zmanim/
├── cmd/
│   ├── lambda/          # Lambda deployment entry point
│   └── server/          # HTTP server (local dev)
├── internal/            # Private application logic
│   ├── handler/         # REST handlers
│   ├── service/         # Business logic
│   └── repository/      # Data access layer
├── api/                 # OpenAPI 3.1 spec (Story 1.6)
├── pkg/                 # Public reusable packages
├── go.mod               # Go module definition
└── README.md            # This file
```

## Getting Started

### Prerequisites

- Go 1.25.4
- PostgreSQL 18

### Running Locally

```bash
# Install dependencies
go mod download

# Run the service
go run cmd/zmanim/main.go
```

The service will start on port 8101 (REST).

### Health Check

Once Story 1.11 (Monitoring & Observability) is complete, the health endpoint will be available.

## Development

This service follows clean architecture principles:

1. **Handlers** (`internal/handler/`) - REST request/response handling
2. **Services** (`internal/service/`) - Business logic and domain operations
3. **Repository** (`internal/repository/`) - Data persistence layer

## API Documentation

API contracts will be defined in Story 1.3 (API Contract Design).

## Related Documentation

- [Architecture](../../../docs/architecture.md)
- [Epics & Stories](../../../docs/epics.md)
- [Database Schema](../../../docs/architecture/database.md) (Story 1.8)
