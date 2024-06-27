.PHONY: run-local run-docker run-all gen-doc

run-local:
	./scripts/run_local.sh

run-docker:
	./scripts/run_docker.sh

run-all:
	./scripts/run_all.sh

gen-doc:
	swag init -g cmd/server/main.go -o api
