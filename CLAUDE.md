# Project guide

Todo list API in Go (`net/http`, MySQL via `database/sql`). Module: `a.com/test`.

## Treat this as a production app
- **Do not treat this as a learning or toy project.** Treat it as a complete production application used by many users and developed by many developers.
- Every recommendation should be what a production team would do: enforced tooling (linters, formatters, CI), reproducible builds, consistent conventions, error handling, context propagation, connection pooling, migrations, testing, observability, graceful shutdown.
- Don't suggest skipping a practice because "it's just a small project". If a practice is optional, say what it costs and when a team needs it.

## About the user
- Experienced backend engineer (TypeScript/Node.js, NestJS, MySQL, Redis, Elasticsearch, RabbitMQ) who is learning Go.
- Compare with Node/TypeScript equivalents when it helps.

## Working rules
- **Do not change any file except `CLAUDE.md`.** The user writes all the code themselves to learn.
- Instead of editing, print the code in the chat, explain it, and review code the user writes.
- Don't run commands that modify the project (installs, `go get`, `go mod tidy`, generators, migrations) unless the user explicitly asks for that specific command.

## Tooling
- Migrations use **goose**, pinned in `go.mod` as a Go tool. Run it with `go tool goose`, for example `go tool goose -dir migrations status`.
- MySQL runs from `docker/compose.yaml`. Connection settings come from `.env` (`MYSQL_HOST`, `MYSQL_PORT`, `MYSQL_USERNAME`, `MYSQL_PASSWORD`, `MYSQL_DATABASE`).
