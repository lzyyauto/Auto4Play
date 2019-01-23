package model

import (
	"fmt"
	"time"

	"wps.cn/plus-order/errors"
)

type RecordStatus string

type RecordType string

type RecordSubmission string

const (

	// RecordTypeSleep represent record im sleep
	RecordTypeSleep = "sleep"
)

const (
	// RecordSubmissionAuto represent record auto commit
	RecordSubmissionAuto = "auto"

	// RecordSubmissionManual represent record manual commit
	RecordSubmissionManual = "manual"
)

const (
	// RecordStatusValid represent record is valid
	RecordStatusValid = "valid"

	// RecordStatusInvalid represent record is invalid
	RecordStatusInvalid = "invalid"
)

type Record struct {
	RecordID         int64            `json:"record_id" db:"record_id" gpo:"record_id"`
	UserID           int64            `json:"user_id" db:"user_id" gpo:"user_id"`
	RecordSubmission RecordSubmission `json:"record_submission" db:"record_submission" gpo:"record_submission"`
	RecordType       RecordType       `json:"record_type" db:"record_type" gpo:"record_type"`
	Status           RecordStatus     `json:"status" db:"status" gpo:"status"`
	Source           string           `json:"source" db:"source" gpo:"source"`
	Remark           string           `json:"remark" db:"remark" gpo:"remark"`
	CreatedAt        int64            `json:"created_at" db:"created_at" gpo:"created_at"`
	UpdatedAt        int64            `json:"updated_at" db:"updated_at" gpo:"updated_at"`
}

const (
	recordTable  = `record`
	recordFields = `record_id, user_id, record_submission,record_type, status, source,
		remark, created_at, updated_at`
)

type RecordMgr struct {
}

// NewRecordMgr create record manager instance
func NewRecordMgr() *RecordMgr {
	return &RecordMgr{}
}

func (m *RecordMgr) CreateRecord(record *Record) (int64, error) {
	createSQL := fmt.Sprintf("insert into `%s` ( %s ) values ( %s )",
		recordTable, recordFields, preArgs(recordFields))

	now := time.Now().Unix()

	// set default parameters
	// record.RecordID = 1
	// record.RecordSubmission = RecordSubmissionManual
	// record.RecordType = RecordTypeSleep
	// record.Status = RecordStatusValid
	record.CreatedAt = now
	record.UpdatedAt = now
	// record.Source = `from test`

	if _, err := db.Exec(createSQL, preData(recordFields, record)...); err != nil {
		return 0, errors.ErrCreateOrder.Bind(err)
	}

	return record.RecordID, nil
}
