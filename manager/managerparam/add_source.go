package managerparam

import (
	"time"

	"github.com/ormushq/ormus/manager/entity"
)

type AddSourceRequest struct {
	Name        string
	Description string
	ProjectID   string
}

type AddSourceResponse struct {
	ID               string
	WriteKeyMetaData entity.WriteKeyMetaData
	Name             string
	Description      string
	ProjectID        string
	OwnerID          string
	Status           entity.Status
	CreateAt         time.Time
	UpdateAt         time.Time
	DeleteAt         *time.Time
}
