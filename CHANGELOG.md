## 0.1.1 (2024-06-29)

### Fix

- **ci**: remove unnecessary folder copy when build image for backend

## 0.1.0 (2024-06-29)

### Feat

- **test**: add test coverage to makefile
- **monitoring**: enable monitoring the backend api's performance using open telemetry
- **monitoring**: add grafana and victoriametrics
- **storage**: add more random data into memory storage
- **user**: add user handler
- **api**: support cors from the frontend writing by vue
- **swagger**: support swagger docmentation
- **api**: implements list and detail api for services
- bootstrap service management api
- **init-project**: using go standard layout

### Fix

- **version**: fix the sort of versions does not following the semantic version standard
- **doc**: fix wrong line-break in readme
- **monitoring**: slow down the grab interval of metrics
- ignore ide config files

### Refactor

- **api**: split out model and viewmodel for layers of api and domain
- **architecture**: refactor whole project using cola architecture
