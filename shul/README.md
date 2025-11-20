# Shul Service

**Purpose:** Shul administration and minyan scheduling

## Overview

The Shul Service handles all synagogue administration, minyan scheduling, and PDF generation for the Shtetl platform. It provides REST APIs for all clients and service-to-service communication.

## Architecture

- **Technology:** Go 1.25.4 + REST (Lambda-compatible)
- **Port:** 8103 (REST)
- **Database:** PostgreSQL (shuls, minyanim, rules, primitives)

## Key Capabilities

- Multi-tenant Shul management
- Minyan scheduling DSL parsing and execution
- Tree-based rule configuration
- Coverage validation (ensures 100% schedule coverage)
- Hierarchical primitive system management
- Hebrew RTL PDF generation
- User/gabai management

## Directory Structure

```
shul/
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
go run cmd/shul/main.go
```

The service will start on port 8002 (REST).

### Health Check

```bash
curl http://localhost:8002/health
# => {"status":"ok","service":"shul"}
```

## Development

This service follows clean architecture principles:

1. **Handlers** (`internal/handler/`) - HTTP request/response handling
2. **Services** (`internal/service/`) - Business logic and domain operations
3. **Repository** (`internal/repository/`) - Data persistence layer

## Novel Patterns

This service implements several unique architectural patterns:

- **Minyan Scheduling DSL** - Non-technical domain-specific language for schedule rules
- **Hierarchical Primitive Cascade** - Global → Country → City → Shul primitive inheritance
- **Coverage Validation Engine** - Real-time 100% schedule coverage validation

See [Architecture: Novel Patterns](../../../docs/architecture.md#novel-architectural-patterns) for details.

## API Documentation

API contracts will be defined in Story 1.3 (API Contract Design).

## Related Documentation

- [Architecture](../../../docs/architecture.md)
- [Epics & Stories](../../../docs/epics.md)
- [Database Schema](../../../docs/architecture/database.md) (Story 1.8)
