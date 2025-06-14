from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from dotenv import load_dotenv
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
import os

load_dotenv()

DATABASE_URL = os.getenv("DATABASE_URL")
engine = create_engine(DATABASE_URL)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

app = FastAPI(title="CreditNinja")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

from routes import auth, upload, ai_disputes, stripe_webhooks

app.include_router(auth.router)
app.include_router(upload.router)
app.include_router(ai_disputes.router)
app.include_router(stripe_webhooks.router)

@app.get("/")
def read_root():
    return {"message": "Welcome to CreditNinja"}
