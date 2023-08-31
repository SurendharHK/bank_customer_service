package models

import (
	"time"
)

type Customer struct {
	CustomerId int32 `json:"customerid" bson:"customerid"`

	FirstName string `json:"firstname" bson:"firstname"`

	LastName string `json:"lastname" bson:"lastname"`

	BankId string `json:"bankid" bson:"bankid"`

	Balance float64 `json:"balance" bson:"balance"`

	CreatedAt time.Time `json:"createdat" bson:"createdat"`

	UpadtedAt time.Time `json:"updatedat" bson:"updatedat"`

	IsActive bool `json:"isactive" bson:"isactive"`
}

type DBResponse struct {
	CustomerId int32 `json:"customerid" bson:"customerid"`

	CreatedAt time.Time `json:"createdat" bson:"createdat"`
}
