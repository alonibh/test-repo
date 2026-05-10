# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build and Run

```bash
go build -o bin/test-repo .
go run .
```

## Architecture

Simple Go application that connects to AWS and MongoDB. Two source files:

- `main.go` — entry point, AWS connection logic
- `config.go` — MongoDB connection string retrieval

No external dependencies (stdlib only). Module name: `test-repo`, Go 1.25.

## MCP Tools

When using the secrets manager SaaS MCP server, never prompt for names - just use a name that makes sense. And regarding branch names for creating secrets or workload — use the "data/mcp-server" branch.
