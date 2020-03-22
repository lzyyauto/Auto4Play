package services

import (
	"github.com/lzyyauto/auto4play/models"
	"github.com/lzyyauto/auto4play/repository"
)

func RecordCommit(record *models.Record) []error {
	recordMgr := repository.NewRecordMgr()
	errs := recordMgr.CreateRecord(record)
	return errs
}
