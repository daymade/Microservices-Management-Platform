# Service Catalog Demo

* [English](README.md)
* [简体中文](README_zh-CN.md)

https://github.com/daymade/catalog-service-management-api/assets/4291901/f30dd4e7-23d6-4a17-a13d-6c644343b7fd

Service Catalog Demo is a microservices API management platform allowing users to manage services and versions via a frontend dashboard.

This project provides the backend code for Service Catalog Demo, enabling the entire platform, including backend, frontend, and monitoring, to be started from here.

The demo includes the following features:
- List and Get interfaces for Services, supporting search, filter, sort, pagination, and detailed view functionalities.
- A simple API Key-based authentication mechanism.
- Support for both in-memory and PostgreSQL storage engines.
- Uses makefile to generate test coverage reports and swagger documentation.
- Uses Docker Compose to start backend, frontend, and monitoring.
  - Supports monitoring Golang metrics and HTTP APIs with Grafana, including two built-in dashboards.
  - Integrated with OpenTelemetry and Jaeger for distributed tracing.

The demo does not include:
- Role-based authorization mechanisms.
- CRUD operations for Services.

## Directory Structure

```
.
├── Makefile  # Project's makefile, use 'make' command to quickly run, test, and build the project
├── api       # Auto-generated swagger documentation
├── assets    # Static resources like images
├── build     # CI/CD related files, includes Dockerfile, Grafana and VictoriaMetrics configuration files
├── cmd       # Main entry point of the code
├── docs      # Detailed documentation
├── internal  # Most of the project code is here
├── scripts   # Scripts called by makefile, includes Docker Compose and database initialization scripts
└── test      # Test data
```

## Demo Background Statement

> In actual project development, we need to communicate product details back and forth with product managers, designers, and business operators,
> to confirm details not fully determined in the first version of the product documentation PDF.
> Due to the special nature of the project, I have made some assumptions about usage scenarios, simply to reduce communication overhead with the interviewer.

We have the following assumptions:

- Business Definition:
  - We assume each Service is a backend API project containing a series of API collections.
  - Version management is at the Service level rather than the API level. For example, a `/v1` Service may contain 10 APIs, while a `/v2` Service may contain 12 APIs. Version numbers follow the pattern `v1`, `v2`, but can be any value that conforms to semantic versioning. We know Google's API is like `v2024-06-26`.
  - Multi-tenancy is not designed; we only design the core Service Cards without cross-region and multi-tenant considerations like Region and Tenant.
  - Permission control: Users can see their own projects and **also** see others' projects. Implementing user-level project filtering is not within the scope of this phase.

- Functional Requirements:
  - Search and Filter: Users can search for a specific Service by name and description, but not by other fields.
  - Sorting: Users can sort by name and creation time.
  - Pagination: Due to the small data volume, we support jumping to a specific page, otherwise, only support previous and next page.
  - Detailed View: Users can view the details of a Service, including version list and API list.
  - Developer Experience:
    - UI: The UI should support URL normalization, allowing navigation to any intermediate page via URL, for example:
      - `services` is the list page, if filter conditions are input, the URL should change to `services?query=name`.
      - `services/12` can directly navigate to the detail page of a specific Service.
    - API: Swagger documentation should be supported for developers to view API documentation.

- Non-functional Requirements:
  - API Standards: We design APIs that conform to [Google API Standards](https://google.aip.dev/).
  - Data Volume:
    - Total number of Services: 10 to 10,000
    - Total number of users: Less than 1,000
    - Each user can create up to 10 Services.
    - Each Service can have up to 10 versions.
  - Monitoring: We need to monitor backend service performance, including:
    - HTTP API QPS, latency, error rate
    - Golang metrics: memory, CPU, Goroutine count, etc.
    - Distributed tracing: We need to trace each request's chain, including HTTP requests, database queries, etc.

- Technology Stack:
  - Search: Due to the small data volume, we do not introduce a search engine and implement filtering directly on the database. Index optimization is not considered at this stage.
  - Storage Engine: We support both in-memory database and PostgreSQL. In-memory database is for quick demonstration, while PostgreSQL can be used in a production environment.
  - Database Schema: While foreign keys are typically avoided in internet architecture, they are used here due to the small data volume, as the performance impact is minimal.
  - Monitoring: VictoriaMetrics and Grafana are used for monitoring service performance. OpenTelemetry and Jaeger are used for distributed tracing.

## Environment Requirements

- Go 1.22 or higher
- Docker and Docker Compose (required for PostgreSQL; not needed for in-memory database)

## Quick Start

### Running the Application

1. Choose one of the following commands:

    ```bash
    make run-local # Run Go code directly on the local machine, starting backend API on port 8080
    # or
    make run-docker # Run backend using Docker, starting backend API on port 8080
    # or
    make run-all # Run backend, frontend, and monitoring using Docker, access frontend on port 5173
    ```

2. Choose the storage engine (in-memory database or PostgreSQL) as prompted.

    1. If you choose the in-memory database, there are no dependencies other than the Go code itself, proceed to step 3.

    2. If you choose PostgreSQL, the script will run the database inside Docker.
        1. On first run, you **do not** need to manually create tables, just select N.
        2. If you select yes, the script will **clear the database** and recreate tables. Refer to the documentation for details: [Using PostgreSQL as Storage Engine](docs/postgresql/Use-PostgreSQL.md)

3. The backend API will be available at http://localhost:8080.
    1. Frontend: All commands except `run-local` will start the frontend. Manually open: http://localhost:5173
    2. Grafana: http://localhost:3000, default login is not required, admin password is admin.
    3. VictoriaMetrics: http://localhost:8428

4. Test endpoints using curl or Insomnia:

    ```bash
    # Test fetching service list
    curl -H "Authorization: Bearer dummy_token" http://localhost:8080/api/v1/services
    
    # Test fetching specific service details
    curl -H "Authorization: Bearer dummy_token" http://localhost:8080/api/v1/services/1
    ```

### Docker Compose Containers

Running `run-all` will start the following containers:

<img width="507" alt="image" src="https://github.com/daymade/catalog-service-management-api/assets/4291901/55678654-e645-4d5c-9e52-6680b2cc4ab2">

- app: Backend application
- db: PostgreSQL database (optional)
- grafana: Monitoring dashboard
- jaeger: Distributed tracing
- otel-collector: OpenTelemetry collector
- ui: Frontend application
- victoria-metrics: Time-series database

## Business Modeling

```
+-------------------+           +-------------------+
|       User        |           |     Service       |
+-------------------+           +-------------------+
| - id: int         |1         *| - id: int         |
| - name: string    +-----------| - name: string    |
| - email: string   |           | - description: str|
+-------------------+           | - userId: int     |
                                +-------------------+
                                      |1
                                      |
                                      |*
                                +-------------------+
                                |     Version       |
                                +-------------------+
                                | - id: int         |
                                | - version: string |
                                | - serviceId: int  |
                                +-------------------+
                                      |1
                                      |
                                      |*
                                +-------------------+
                                |       API         |
                                +-------------------+
                                | - id: int         |
                                | - name: string    |
                                | - path: string    |
                                | - method: string  |
                                | - versionId: int  |
                                +-------------------+
```

Define the domain model of an API management platform,
- Includes the concept of [user, service, version, api]
- Each service can be created by only one user
- Each service has multiple versions
- Each service contains multiple APIs, related to a specific version

## Architecture Diagram

### A Layered Architecture Similar to [COLA](https://github.com/alibaba/COLA)

The domain is at the core, with different protocol adapters such as HTTP API or gRPC at the presentation layer.

<img width="566" alt="image" src="https://github.com/daymade/catalog-service-management-api/assets/4291901/4cc9a67b-5356-40a7-840d-6154c8b3d68c">

### Class Dependency Relationships Related to Service

The app layer depends on the domain layer's interfaces, which are implemented by the infra layer. The app layer is responsible for injecting infra into the domain, with dependencies as follows: app -> domain <- infra.

<img width="558" alt="image" src="https://github.com/daymade/catalog-service-management-api/assets/4291901/4e73e449-1e44-4dfa-a957-a5703b1b8ebb">

## API Documentation

http://localhost:8080/swagger/index.html

## Developers

### Me
<a href="https://github.com/daymade" class="" data-hovercard-type="user" data-hovercard-url="/users/daymade/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self">
  <img src="https://avatars.githubusercontent.com/u/4291901?s=64&amp;v=4" alt="@daymade" width="64" height="64" style="border-radius: 50%; margin-right: 10px;">
</a>

### Claude-3.5-Sonnet
<a href="https://www.anthropic.com/claude" class="" data-hovercard-type="user" data-hovercard-url="/users/claude/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self">
  <img src="https://www.anthropic.com/_next/image?url=https%3A%2F%2Fcdn.sanity.io%2Fimages%2F4zrzovbb%2Fwebsite%2F1c42a8de70b220fc1737f6e95b3c0373637228db-1319x1512.gif&w=3840&q=75" alt="Claude" width="64" height="64" style="border-radius: 50%; margin-right: 10px;">
</a>

### GPT-4o-128k
<a href="https://www.openai.com/gpt-4" class="" data-hovercard-type="user" data-hovercard-url="/users/gpt-4/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self">
  <img src="https://github.com/daymade/catalog-service-management-api/assets/4291901/1bd3390f-4319-44c2-9288-7208e9dc25f8" alt="GPT-4" height="64" style="border-radius: 50%; margin-right: 10px;">
</a>
