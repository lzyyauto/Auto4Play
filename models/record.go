package models

type RecordStatus string

type RecordType string

type RecordSubmission string

const (

	// RecordTypeSleep represent record im sleep
	RecordTypeSleep = "sleep"
	// RecordTypeFitness represent record im Fitness
	RecordTypeFitness = "fitness"
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

//设置表名
func (Record) TableName() string {
	return "record"
}
