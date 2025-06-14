from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from ..models import CreditReport, Dispute
from ..main import SessionLocal
import os
import openai

router = APIRouter(prefix="/ai", tags=["ai"])
openai.api_key = os.getenv("OPENAI_API_KEY")


def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

@router.post('/generate/{report_id}')
def generate_disputes(report_id: int, db: Session = Depends(get_db)):
    report = db.query(CreditReport).filter(CreditReport.id == report_id).first()
    if not report:
        raise HTTPException(status_code=404, detail="Report not found")
    # In reality we'd parse the PDF and feed to OpenAI
    prompt = "Generate a dispute letter for negative items."
    try:
        resp = openai.Completion.create(model="text-davinci-003", prompt=prompt, max_tokens=150)
        letter = resp.choices[0].text.strip()
    except Exception:
        letter = "Mock dispute letter content"
    dispute = Dispute(item="Item description", letter=letter, report_id=report.id)
    db.add(dispute)
    db.commit()
    db.refresh(dispute)
    return {"dispute_id": dispute.id, "letter": dispute.letter}
