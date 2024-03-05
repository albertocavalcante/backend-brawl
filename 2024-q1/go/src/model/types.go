package model

type TransactionType string

const (
	Credit TransactionType = "c"
	Debit  TransactionType = "d"
)

func AllowedTransactionTypes() []TransactionType {
	return []TransactionType{Credit, Debit}
}

type AccountBalanceInCents int64

type AccountTransactionLimit int64

type Customer struct {
	ID      int                     `json:"id"`
	Limit   AccountTransactionLimit `json:"limite"`
	Balance AccountBalanceInCents   `json:"saldo"`
}

// NewCustomer creates a new customer, settings its ID, limit and balance
func NewCustomer(id int, limit int64) *Customer {
	var initialBalance int64 = 0

	return &Customer{
		ID:      id,
		Limit:   AccountTransactionLimit(limit),
		Balance: AccountBalanceInCents(initialBalance),
	}
}

type Transaction struct {
	ID          int
	CustomerID  int
	Type        TransactionType `json:"tipo"`
	Description string          `json:"descricao"`
}

type AccountBalance struct {
	CustomerID     int
	BalanceInCents int64
}
