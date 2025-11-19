# Kehilla Service

**Purpose:** Public-facing community API for kehilla members (congregants)

> **Note:** "Kehilla" means "community" in Hebrew - this service serves regular congregation members.

## Overview

The Kehilla Service provides public RESTful APIs for community members to access minyan schedules, manage subscriptions, and receive notifications. It includes Redis caching for high-performance schedule queries.

## Architecture

- **Technology:** Go 1.25.4 + REST
- **Port:** 8003 (REST)
- **Database:** PostgreSQL (subscriptions, notifications, user preferences)
- **Cache:** Redis (schedule queries, session data)

## Key Capabilities

- Schedule queries (today's times, weekly view)
- Subscription management
- SMS/push notification coordination
- Alert delivery
- Automation webhook triggers

## Directory Structure

```
kehilla/
├── cmd/kehilla/         # Service entry point
│   └── main.go          # HTTP server
├── internal/            # Private application logic
│   ├── handler/         # REST handlers
│   ├── service/         # Business logic
│   └── repository/      # Data access layer
├── api/                 # OpenAPI specs (Story 1.3)
├── pkg/                 # Public reusable packages
├── go.mod               # Go module definition
└── README.md            # This file
```

## Getting Started

### Prerequisites

- Go 1.25.4
- PostgreSQL 18
- Redis 8.4

### Running Locally

```bash
# Install dependencies
go mod download

# Run the service
go run cmd/kehilla/main.go
```

The service will start on port 8003 (REST).

### Health Check

```bash
curl http://localhost:8003/health
# => {"status":"ok","service":"kehilla"}
```

## Development

This service follows clean architecture principles:

1. **Handlers** (`internal/handler/`) - HTTP request/response handling
2. **Services** (`internal/service/`) - Business logic and notification orchestration
3. **Repository** (`internal/repository/`) - Data persistence and caching layer

## Consumers

- Mobile app (shtetl-mobile)
- Public web interface (shtetl-web)
- Automation systems (webhooks)

## API Documentation

API contracts will be defined in Story 1.3 (API Contract Design).

## Related Documentation

- [Architecture](../../../docs/architecture.md)
- [Epics & Stories](../../../docs/epics.md)
- [Database Schema](../../../docs/architecture/database.md) (Story 1.8)
