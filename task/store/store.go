// Package store provides interfaces and types for working with account data storage.
package store

import "context"

// Account represents a financial account with its properties and balances.
type Account struct {
	// ID is the unique identifier of the account.
	ID uint64 `json:"id"`

	// Asset is the type of cryptocurrency held in the account.
	Asset string `json:"asset"`

	// Balance is the current amount of cryptocurrency in the account.
	Balance float64 `json:"balance"`

	// BalanceUSD is the USD equivalent of the cryptocurrency balance.
	BalanceUSD float64 `json:"balance_usd"`

	// IsRecount indicates whether the account requires a recount operation
	// for its USD balances or transactions.
	IsRecount bool `json:"is_recount"`
}

// Store defines a generic interface for data storage operations.
type Store[T any] interface {
	// FindOne retrieves a single entity matching the provided criteria.
	//
	// It returns the found entity or an error if the operation fails.
	FindOne(context.Context, T) (T, error)
}
