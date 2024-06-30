#!/bin/bash
set -e

go test -v ./cmd/... ./internal/...
