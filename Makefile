.PHONY: run-local run-docker

run-local:
	./scripts/run_local.sh

run-docker:
	./scripts/run_docker.sh

build-production:
	./scripts/build_production.sh
