# Catalog Service Management API

* [English](README.md)
* [简体中文](README_zh-CN.md)

https://github.com/daymade/catalog-service-management-api/assets/4291901/f30dd4e7-23d6-4a17-a13d-6c644343b7fd

Catalog-Demo is a microservices API management platform that allows users to manage services and versions through a frontend dashboard.

This project is the backend code for Catalog-Demo. You can use it to start the entire platform, including the backend, frontend, and monitoring.

The demo includes the following features:
- List and Get APIs for services, supporting search, filter, sort, pagination, and detailed view.
- Simple authentication mechanism based on API Key.
- Support for both in-memory and PostgreSQL storage engines.
- Test code and swagger documentation.
- Grafana monitoring.

The demo does not include:
- Role-based authorization mechanism.
- CRUD operations for services.

Known bugs:
- The width of the service detail page on the frontend is incorrect. Fixing it requires significant effort, so it is temporarily skipped as it is still usable.
- Grafana can automatically import data sources but requires manual import of the Dashboard: `build/config/grafana/dashboards/Go Metrics-1719497538877.json`.

Directory structure:
```
.
├── Makefile          # Makefile for the project, use `make` to quickly run, test, build the project
├── api               # Auto-generated swagger documentation
├── assets            # Store images and other static resources
├── build             # CI/CD related, including Dockerfile, Grafana, and VictoriaMetrics configuration files
├── cmd               # Main entry point of the code
├── docs              # Detailed documentation
├── internal          # Most of the project code resides here
├── scripts           # Scripts called by Makefile, including docker-compose and database initialization scripts
└── test
```

## Demo Background Declaration

> In actual project development, we need to communicate back and forth with product managers, designers, and business operators to confirm details that were not fully determined in the initial product documentation. Due to the special nature of this project, I have made some simplified assumptions for usage scenarios, solely to reduce communication overhead with the interviewer.

Assumptions:

- **Business Definition**:
  - Each service is assumed to be a backend API project containing a set of APIs.
  - **Version Management**: Services are managed at the service level, not the API level. For example, `/v1` of a service may contain 10 APIs, while `/v2` may contain 12 APIs. Versioning can follow any semantic versioning rules, like `v1`, `v2`, or even `v2024-06-26` as seen in Google Cloud APIs.
  - **Multi-tenant**: We design only the core Service Cards without cross-region and multi-tenant designs like Region or Tenant.
  - **Access Control**: Users can see their own projects and **also** see others' projects. Implementing user-based project filtering is not within the scope of this phase.

- **Functional Requirements**:
  - **Search**: Users can search for a specific service by name and description.
  - **Filter**: Users can filter services by name and description.
  - **Sort**: Users can sort services by name and creation time.
  - **Pagination**: Due to the small data volume, support for jumping to a specific page is not required; only previous and next page navigation is needed.
  - **View Details**: Users can view service details, including versions and API lists.
  - **Developer Experience**:
    - **UI**: The URL should be structured, allowing navigation to any intermediate page, such as:
      - `services` is the list page, and if a filter is applied, it becomes `services?query=name`.
      - Navigating to `services/contact-us` or `services/locate-us` should directly take the user to the details page of a specific service.

- **Non-functional Requirements**:
  - **API Standards**: APIs are designed to conform to [Google API Standards](https://google.aip.dev/).
  - **Data Volume**:
    - Total number of services: 10 - 10,000
    - Total number of users: fewer than 1,000
    - Each user can create a maximum of 10 services.
    - Each service can have up to 10 versions.

- **Technology Stack**:
  - **Search**: Due to the small data volume, we do not introduce a search engine and instead implement filtering directly in the database.
  - **Storage Engine**: Support for both in-memory and PostgreSQL storage engines. In-memory is used for fast demonstrations, while PostgreSQL is for production environments.
  - **Database Structure**: In internet architecture, foreign keys are generally not used. Given the small data volume, foreign keys do not impact performance and are used.
  - **Monitoring**: Use VictoriaMetrics and Grafana for performance monitoring.

## Runtime Environment

- Go 1.22 or higher
- Docker and Docker Compose (required when using PostgreSQL; not needed for in-memory database)

## Quick Start

### Run

1. Choose one of the following commands:

    ```bash
    make run-local  # Run Go code directly on the local machine
    # Or
    make run-docker  # Run backend and frontend using Docker
    # Or
    make run-all  # Run backend, frontend, and monitoring using Docker, with in-memory database for quick demonstration
    ```

2. Follow the prompts to choose a storage engine (in-memory database or PostgreSQL):

   1. If using the in-memory database, proceed to step 3.
   2. If using PostgreSQL, the script will run the database in Docker.
      1. Choose whether to rebuild the database. The script will handle table creation automatically. No need to choose for initial runs.
      2. Refer to the documentation for details: [Using PostgreSQL as a Storage Engine](docs/postgresql/Use-PostgreSQL.md)

3. The backend API will be available at `http://localhost:8080`.
   - Frontend: `http://localhost:5173`
   - Grafana: `http://localhost:3000`, user: admin, password: admin
   - VictoriaMetrics: `http://localhost:8428`

4. Test endpoints using curl or Insomnia:

    ```bash
    # Test fetching the service list
    curl -H "Authorization: Bearer dummy_token" http://localhost:8080/api/v1/services

    # Test fetching specific service details
    curl -H "Authorization: Bearer dummy_token" http://localhost:8080/api/v1/services/1
    ```

## Domain Modeling
```
+-------------------+           +-------------------+
|       User        |           |     Service       |
+-------------------+           +-------------------+
| - id: int         |1         *| - id: int         |
| - name: string    +-----------| - name: string    |
| - email: string   |           | - description: string |
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
- Include the concept of [user, service, version, API].
- Each service can be created by only one user.
- Each service has multiple versions.
- Each service contains multiple APIs, related to specific versions.

## Architecture Diagram

### [COLA](https://github.com/alibaba/COLA)-like Layered Architecture

The domain is at the core, with adapters like HTTP API or gRPC in the presentation layer.
<img width="558" alt="image" src="https://github.com/daymade/catalog-service-management-api/assets/4291901/4cc9a67b-5356-40a7-840d-6154c8b3d68c">

### Service-Related Class Dependency

The app layer depends on the interfaces in the domain layer, and these interfaces are implemented by the infra layer. The app layer is responsible for injecting infra into the domain. The dependency relationship is: app -> domain <- infra.
<img width="558" alt="image" src="https://github.com/daymade/catalog-service-management-api/assets/4291901/4e73e449-1e44-4dfa-a957-a5703b1b8ebb">

## API Documentation

http://localhost:8080/swagger/index.html

## Developer(s)

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
