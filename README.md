# Fiber E-commerce System API

A Go REST API built with Fiber, GORM, MySQL, JWT authentication, and the Midtrans Snap payment gateway.

## Prerequisites

- Go 1.21 or newer
- MySQL 8 or a compatible MySQL server
- A Midtrans account with sandbox credentials for local development

## Environment setup

Copy the safe template to a local `.env` file:

```powershell
Copy-Item .env.example .env
```

On macOS or Linux, use `cp .env.example .env`. Fill in the values in `.env`; the application validates mandatory configuration during startup and exits if one is missing. A physical `.env` file is optional in hosted environments where variables are supplied by the platform.

| Variable | Required | Description |
| --- | --- | --- |
| `APP_PORT` | No | HTTP port. Defaults to `8080`. |
| `APP_ENV` | No | Application environment label. Defaults to `development`. |
| `DB_HOST` | Yes | MySQL host name or IP address. |
| `DB_PORT` | Yes | MySQL TCP port, normally `3306`. |
| `DB_USER` | Yes | MySQL user name. |
| `DB_PASSWORD` | Yes | MySQL password. |
| `DB_NAME` | Yes | MySQL database name. |
| `JWT_SECRET` | Yes | Strong, randomly generated secret used to sign authentication tokens. |
| `MIDTRANS_SERVER_KEY` | Yes | Midtrans server key for the selected environment. |
| `MIDTRANS_CLIENT_KEY` | Yes | Midtrans client key for the selected environment. |
| `MIDTRANS_ENV` | No | `sandbox` (default) or `production`. |

Do not use example text as a real secret. Generate a long random `JWT_SECRET` and keep all real values outside version control.

## Local database setup

Create an empty database and a dedicated user with access only to that database. For example, from a MySQL administrator session:

```sql
CREATE DATABASE ecommerce CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'ecommerce_app'@'localhost' IDENTIFIED BY 'choose-a-strong-local-password';
GRANT ALL PRIVILEGES ON ecommerce.* TO 'ecommerce_app'@'localhost';
FLUSH PRIVILEGES;
```

Set the matching `DB_*` values in `.env`. The application expects the required tables to exist; apply the repository's schema or migrations before using the endpoints.

## Midtrans sandbox setup

1. Sign in to the Midtrans dashboard and switch to the sandbox environment.
2. Open **Settings → Access Keys**.
3. Put the sandbox server and client keys in your local `.env` as `MIDTRANS_SERVER_KEY` and `MIDTRANS_CLIENT_KEY`.
4. Keep `MIDTRANS_ENV=sandbox` for local development.

Use production keys only with `MIDTRANS_ENV=production`. Never log keys, return them from an API, or commit them. If any credential has ever been exposed in this repository or its Git history, rotate it in the Midtrans dashboard; removing it from the current files does not revoke it.

## Run the application

Download dependencies and start the API:

```bash
go mod download
go run ./cmd
```

By default, the API listens at `http://localhost:8080` and registers routes under `/api/v1`.

## Verification

Run the standard Go checks before submitting changes:

```bash
go fmt ./...
go vet ./...
go test ./...
go build ./...
```

## Security

`.env` and all environment-specific variants are ignored by Git; `.env.example` is the only exception and must contain variable names with empty or clearly fake values only. Never commit database passwords, JWT secrets, Midtrans keys, API keys, tokens, or other credentials. Use a secrets manager or deployment platform environment variables in production.

If credentials entered Git history, follow the separate [manual history cleanup plan](SECURITY_HISTORY_CLEANUP.md) only after rotating them and coordinating with every collaborator.
