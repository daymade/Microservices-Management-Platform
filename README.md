Here's the English translation of the document:

# Service Catalog Demo

* [English](README.md)
* [简体中文](README_zh-CN.md)

<a href="https://asciinema.org/a/CATqU6FRQWSHzVb6L193EoKTR" target="_blank"><img src="https://asciinema.org/a/CATqU6FRQWSHzVb6L193EoKTR.svg" /></a>

https://github.com/daymade/catalog-service-management-api/assets/4291901/7ecb4d40-935f-4f2a-9c41-2d376f62d5b8

Service Catalog Demo is a microservice API management platform where users can manage services and versions through a frontend dashboard.

This project is the backend code for the Service Catalog Demo, from which you can launch the entire platform, including the backend, frontend, and monitoring.

## Features included in the demo:
- Basic requirements: Service supports search, filtering, sorting, pagination, viewing details, and other functions
	- List returns a list of services (supports filtering, sorting, pagination):
		- Implemented at the `/api/v1/services` endpoint
		- Supports fuzzy search and filtering through `name` and `description` fields
		- Supports sorting by `name` and `created_at` fields
		- Implements pagination mechanism based on `offset` and `limit`
	- Get retrieves details of a specific service:
		- Implemented at the `/api/v1/services/{id}` endpoint
		- Includes all version information of the service
- Authentication mechanism:
	- Implements a simple authentication mechanism based on API Key
- Multiple storage engine support:
	- Supports in-memory database and PostgreSQL as storage engines
- Monitoring and tracing
	- Integrates Grafana and VictoriaMetrics for performance monitoring
		- Provides two pre-configured dashboards for monitoring Golang Metrics and HTTP API
	- Uses OpenTelemetry and Jaeger for distributed tracing
- Developer experience:
	- Uses Swagger to generate API documentation
	- Provides a Makefile to simplify development and deployment processes
	- Automated unit tests and integration tests, supports generating test coverage reports
- Containerized deployment:
	- Uses docker compose to start backend, frontend, and monitoring

Features not included in the demo:
- Role-based authorization mechanism
- CRUD operations for Service

## Directory structure:

```
.
├── Makefile # Project makefile, use make commands to quickly run, test, and build the project
├── api      # Automatically generated swagger documentation
├── assets   # Static resources like images
├── build    # CI/CD related, including docker file, grafana and victoriametrics configuration files
├── cmd      # Code main entry
├── docs     # Detailed documentation
├── internal # Most of the project code is here
├── scripts  # Scripts called by makefile, including docker-compose and database initialization scripts
└── test     # Stores test data
```

## Demo Background Statement

> In actual project development, we need to communicate back and forth with product managers, designers, and business operations personnel to determine details that were not fully defined in the first version of the product documentation PDF.
> Due to the special nature of this project, I've made some simple assumptions about usage scenarios here, just to reduce communication overhead with the interviewer.

We have the following assumptions:

- Business definition:
	- We assume each Service is a backend API project containing a series of API collections
	- Version management: Services have version management, with versioning at the Service level rather than the API level. For example, a `/v1` Service might contain 10 APIs, while a `/v2` Service might contain 12 APIs. Version numbers follow the rule of `v1`, `v2`, but can be any value conforming to semantic versioning. We know that Google Cloud's APIs are like `v2024-06-26`.
	- Multi-tenancy: Only the core Service Cards are designed, without cross-region and multi-tenant design considerations like Region, Tenant.
	- Access control: Users can see their own projects and **also** see other people's projects. Implementing project filtering at the user level is not within the scope of this phase.

- Functional requirements:
	- Search and filter: Users can search for specific Services by name and description, other fields are not supported
	- Sorting: Users can sort by name and creation time
	- Pagination: Since the data volume is small, it can support jumping to a specific page, otherwise only previous and next page navigation is needed
	- View details: Users can view Service details, including version list, API list, etc.
	- Developer experience:
		- UI: Needs to support URL normalization, allowing navigation to any intermediate page through URL, for example:
			- `services` is the list page, if filter conditions are entered, the URL should change to `services?query=name`.
			- `services/12` can directly jump to the details page of a specific Service.
		- API: Needs to support Swagger documentation for developers to easily view interface documentation.

- Non-functional requirements:
	- API specification: We design APIs that comply with the [Google API Specification](https://google.aip.dev/).
	- Data volume:
		- Total number of Services: 10 to 10000
		- Total number of users: Below 1000
		- Each user can create a limited number of Services, maximum 10 services.
		- Number of versions per Service: Maximum 10 versions.
	- Monitoring: We need to monitor the performance of backend services, including:
		- HTTP API QPS, latency, error rate
		- Golang Metrics: Memory, CPU, number of Goroutines, etc.
		- Distributed tracing: We need to trace the path of each request, including HTTP requests, database queries, etc.

- Technology selection:
	- Search: Due to the small data volume, we don't introduce a search engine. We can use PostgreSQL's built-in trigram index to optimize fuzzy search performance.
	- Storage engine: We support both in-memory database and PostgreSQL as storage engines. The in-memory database is for quick demonstrations, while PostgreSQL can be used for production environments.
	- Database structure: Internet architectures generally don't use foreign keys. In this scenario, the data volume is very small, so foreign keys won't significantly affect performance, so we used foreign keys.
	- Monitoring: Use VictoriaMetrics and Grafana to monitor service performance, OpenTelemetry and Jaeger for distributed tracing.

## Running Environment

- Go 1.22 or higher
- Docker and Docker Compose (needed when using PostgreSQL), not required when using in-memory database

## Quick Start

### Running

1. Choose one of the following commands:

    ```bash
    make run-local # Run Go code directly on the local machine, start backend API on port 8080
    # or
    make run-docker # Run backend using docker, start backend API on port 8080
    # or
    make run-all # Run backend, frontend, and monitoring using docker, access frontend on port 5173
    ```

2. Follow the prompts to select the storage engine (in-memory database or PostgreSQL)

	1. If you choose to use the in-memory database, there are no dependencies other than the Go code itself, proceed to step 3.

	2. If you choose PostgreSQL, the script will run the database in Docker.
		1. For the first run, you **don't need** to manually create tables, default N is fine.
		2. If you choose yes, the script will **clear the database**, then recreate tables, for details please refer to the document: [Using PostgreSQL as Storage Engine](docs/postgresql/Use-PostgreSQL.md)

3. The backend API will be available at http://localhost:8080.
	1. Frontend: All commands except `run-local` will start the frontend, please manually open the address: http://localhost:5173
	2. Grafana: http://localhost:3000, no login required by default, admin username is admin, password is admin
	3. VictoriaMetrics: http://localhost:8428

4. Test endpoints using curl or Insomnia:

   ```bash
   # Test getting service list
   curl -H "Authorization: Bearer dummy_token" http://localhost:8080/api/v1/services
   
   # Test getting specific service details
   curl -H "Authorization: Bearer dummy_token" http://localhost:8080/api/v1/services/1
   ```

### Docker compose container list

If you chose `run-all`, the following containers will be started:

<img width="507" alt="image" src="https://github.com/daymade/catalog-service-management-api/assets/4291901/55678654-e645-4d5c-9e52-6280b2cc4ab2">

- app: Backend application
- db: PostgreSQL database (optional)
- grafana: Monitoring dashboard
- jaeger: Distributed tracing
- otel-collector: Open Telemetry collector
- ui: Frontend application
- victoria-metrics: Time series database

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
- Include the concepts of [user, service, version, API]
- Each service can be created by only one user
- Each service has multiple versions
- Each service contains multiple APIs, related to specific versions

## Architecture Diagram

### Layered architecture similar to [COLA](https://github.com/alibaba/COLA)

With domain as the core, different protocol adapters like HTTP API or gRPC can exist at the presentation layer.

<img width="566" alt="image" src="https://github.com/daymade/catalog-service-management-api/assets/4291901/4cc9a67b-5356-40a7-840d-6154c8b3d68c">

### Class dependency relationship related to Service

The app layer depends on the domain layer's interfaces, the domain's interfaces are implemented by the infra layer, and the app is responsible for injecting infra into domain. The dependency relationship is: app -> domain <- infra.

<img width="558" alt="image" src="https://github.com/daymade/catalog-service-management-api/assets/4291901/4e73e449-1e44-4dfa-a957-a5703b1b8ebb">

## Testing

This project includes various types of tests to ensure code quality and functional correctness.

### Unit Tests

To run all unit tests:

```bash
make test
```

This will execute all unit tests and display the results.

### Test Coverage

To generate a test coverage report:

```bash
make test-coverage
```

This command will run the tests and generate a coverage report. You can view detailed coverage information in the `coverage.html` file.

### Integration Tests

To run integration tests:

```bash
make test-integration
```

Integration tests check if the interactions between different components of the system are working correctly.

### Cleaning Test Files

To clean up files generated during the testing process:

```bash
make test-clean
```

This will remove the test coverage reports and other temporary files.

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
