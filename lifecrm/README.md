# LifeCRM MVP

A Go/PostgreSQL starter for a life-insurance CRM modeled around the core workflow of a GoHighLevel-style platform.

## Included

- Multi-organization data model
- Users and roles
- Contacts/leads
- Pipelines and opportunities
- Activities, tasks, and notes
- Carriers, applications, and policies
- Go HTTP API
- PostgreSQL Docker setup
- Render Blueprint for a Go web service and managed PostgreSQL

## Run locally

Start PostgreSQL:

```sh
docker compose up -d
```

Apply the schema:

```sh
psql "postgres://lifecrm:lifecrm@localhost:5432/lifecrm?sslmode=disable" -f migrations/001_init.sql
```

Set environment variables and download dependencies:

```sh
cp .env.example .env
set -a && . ./.env && set +a
go mod tidy
```

Start the API:

```sh
go run ./cmd/server
```

Health check: `GET http://localhost:8080/health`

Dashboard: `GET http://localhost:8080/api/v1/dashboard`

Contacts: `GET http://localhost:8080/api/v1/contacts`

The `.env.example` credentials are for local development only. Do not commit `.env` or use these credentials in a deployed service.

## Deploy to Render

Push this project folder to a private GitHub repository, then create a Render Blueprint from that repository. Render uses `render.yaml` to provision the Go web service and managed PostgreSQL database. Apply `migrations/001_init.sql` to the provisioned database before using database-backed features. Do not put real client data into this MVP.

## Next implementation milestone

Add authentication middleware and organization-scoped CRUD for contacts, then pipeline/opportunity CRUD. After that, add appointments, conversations, and the workflow engine.

## Production note

**Do not use this demo API with real insurance-client data.** Before production, add secure authentication and authorization, server-side sessions, audit logging, secrets management, encryption, consent and communication-preference records, secure document handling, backups and recovery, monitoring and error reporting, and a compliance/security review appropriate to the jurisdictions and insurance operations involved. Passwords must be hashed with Argon2id or bcrypt; organization-level authorization must be enforced on every query. Add CSRF protection where applicable, and configure TLS and secure headers.
