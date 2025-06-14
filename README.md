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

## Python FastAPI Version

A minimal FastAPI implementation is located in `creditninja_py/`. Use it for experimentation with AI-powered dispute generation.

### Quick Start

```bash
cd creditninja_py
cp env.example .env
pip install -r requirements.txt
uvicorn main:app --reload
```

This version includes basic auth, PDF upload, AI dispute generation and Stripe webhook stubs. It is for educational use only and does not constitute legal advice.
