package model

import (
	"time"
)

type Transaction struct {
	UUID          string //primary key, non editable
	Type          string //expense/income
	Amount        int64
	CategoryID    int //foreign key
	Description   string
	TransactionAt time.Time //the actual that transaction happend
	CreatedAt     time.Time //the time when user created the transaction in DB
	UpdatedAt     time.Time
}

type Category struct {
	ID          int    //primary key, non editable
	Name        string //non editable
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
