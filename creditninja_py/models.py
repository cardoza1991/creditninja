from sqlalchemy import Column, Integer, String, Boolean, ForeignKey, Text
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import relationship

Base = declarative_base()

class User(Base):
    __tablename__ = 'users'
    id = Column(Integer, primary_key=True, index=True)
    email = Column(String, unique=True, index=True, nullable=False)
    hashed_password = Column(String, nullable=False)
    is_active = Column(Boolean, default=True)
    is_admin = Column(Boolean, default=False)
    subscription_active = Column(Boolean, default=False)
    reports = relationship('CreditReport', back_populates='owner')

class CreditReport(Base):
    __tablename__ = 'credit_reports'
    id = Column(Integer, primary_key=True, index=True)
    filename = Column(String)
    owner_id = Column(Integer, ForeignKey('users.id'))
    owner = relationship('User', back_populates='reports')
    disputes = relationship('Dispute', back_populates='report')

class Dispute(Base):
    __tablename__ = 'disputes'
    id = Column(Integer, primary_key=True, index=True)
    item = Column(Text)
    letter = Column(Text)
    status = Column(String, default='pending')
    report_id = Column(Integer, ForeignKey('credit_reports.id'))
    report = relationship('CreditReport', back_populates='disputes')
