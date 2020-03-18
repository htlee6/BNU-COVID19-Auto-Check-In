package recordusers

import (
	"SignInHelper/pkg/db"
)

type RecordService struct {
	dbClient db.Client
}

func New() *RecordService {
	return &RecordService{dbClient: db.NewClient()}
}
