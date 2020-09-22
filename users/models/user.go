package models

import (
	"time"
)

//Users Fields
type Users struct {
	BankruptcyIndicatorFlag bool      `json:"bankruptcyIndicatorFlag" sql:",notnull"`
	CompanyName             string    `json:"companyName" sql:",notnull"`
	CreatedDate             time.Time `json:"createdDate"`
	DateOfBirth             time.Time `json:"dateOfBirth"`
	FirstName               string    `json:"firstName" validate:"required,min=3,max=50" sql:",notnull"`
	LastName                string    `json:"lastName" validate:"required,max=25" sql:",notnull"`
	LegalEntityID           int64     `json:"legalEntityId"`
	LegalEntityStage        string    `json:"legalEntityStage"`
	LegalEntityType         string    `json:"legalEntityType"`
}
