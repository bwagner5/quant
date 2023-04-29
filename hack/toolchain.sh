#!/usr/bin/env bash
set -euo pipefail

go install github.com/google/go-licenses@v1.6.0
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1
go install golang.org/x/vuln/cmd/govulncheck@v0.0.0-20230217165152-67742527d09b