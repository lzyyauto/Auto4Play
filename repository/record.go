package repository

import (
	db "../databases"

	"../models"
)

// const (
// 	recordTable  = `record`
// 	recordFields = `record_id, user_id, record_submission,record_type, status, source,
// 		remark, created_at, updated_at`
// )

type RecordMgr struct {
}

// NewRecordMgr create record manager instance
func NewRecordMgr() *RecordMgr {
	return &RecordMgr{}
}

func (m *RecordMgr) CreateRecord(record *models.Record) []error {
	return db.DB.Create(record).GetErrors()
}
