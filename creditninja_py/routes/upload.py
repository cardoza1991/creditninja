from fastapi import APIRouter, UploadFile, File, HTTPException, Depends
from sqlalchemy.orm import Session
from ..models import CreditReport
from ..main import SessionLocal
import magic, os, shutil
from PyPDF2 import PdfReader

router = APIRouter(prefix="/upload", tags=["upload"])


def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

UPLOAD_DIR = 'uploads'
os.makedirs(UPLOAD_DIR, exist_ok=True)

@router.post('/')
def upload_report(file: UploadFile = File(...), db: Session = Depends(get_db)):
    mime = magic.from_buffer(file.file.read(1024), mime=True)
    file.file.seek(0)
    if mime != 'application/pdf':
        raise HTTPException(status_code=400, detail='Only PDF files allowed')
    path = os.path.join(UPLOAD_DIR, file.filename)
    with open(path, 'wb') as out:
        shutil.copyfileobj(file.file, out)
    reader = PdfReader(path)
    text = "\n".join(page.extract_text() or '' for page in reader.pages)
    report = CreditReport(filename=path)
    db.add(report)
    db.commit()
    return {"message": "Uploaded", "text_snippet": text[:200]}
