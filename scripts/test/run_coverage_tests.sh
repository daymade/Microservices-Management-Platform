#!/bin/bash
set -e

go test -v -coverprofile=coverage.out ./cmd/... ./internal/...
go tool cover -html=coverage.out -o coverage.html
go tool cover -func=coverage.out
echo "Coverage report generated: coverage.html"
open coverage.html
