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

## Stack

* **Fiber** – blazing fast web framework
* **HTMX + Tailwind** – dynamic HTML without JS bloat
* **PostgreSQL** – primary datastore
* **Stripe** – payments
* **gopdf** – dispute‑letter PDF generation
