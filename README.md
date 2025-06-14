# CreditNinja

Lean & mean Golang credit‑repair automation platform.

## Quick Start

```bash
git clone <repo-url>
cd creditninja
cp .env.example .env   # edit secrets
# start postgres (builds an image from Dockerfile.postgres)
docker build -f Dockerfile.postgres -t creditninja-db .
docker run --name creditdb -p 5432:5432 -d creditninja-db
go mod tidy
go run ./cmd
```

The app can integrate with a local AI service for PDF parsing and dispute letter generation.
Set `LOCAL_AI_URL`, `PARSER_URL`, and `APP_URL` in your `.env` to the endpoints of your microservices and the public app URL used in verification emails.

## Stack

* **Fiber** – blazing fast web framework
* **HTMX + Tailwind** – dynamic HTML without JS bloat
* **PostgreSQL** – primary datastore
* **Stripe** – payments
* **gopdf** – dispute‑letter PDF generation
