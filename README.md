# creditninja# CreditNinja

Lean & mean Golang credit‑repair automation platform.

## Quick Start

```bash
git clone <repo-url>
cd creditninja
cp .env.example .env   # edit secrets
go mod tidy
go run ./cmd
```

The app can integrate with a local AI service for PDF parsing and dispute letter generation.
Set `LOCAL_AI_URL` and `PARSER_URL` in your `.env` to the endpoints of your microservices.

## Stack

* **Fiber** – blazing fast web framework
* **HTMX + Tailwind** – dynamic HTML without JS bloat
* **PostgreSQL** – primary datastore
* **Stripe** – payments
* **gopdf** – dispute‑letter PDF generation
