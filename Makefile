.PHONY: run-local run-docker build-production gen-doc

run-local:
	./scripts/run_local.sh

run-docker:
	./scripts/run_docker.sh

build-production:
	./scripts/build_production.sh

gen-doc:
	swag init -g cmd/server/main.go -o api
