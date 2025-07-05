// Package web implements HTTP handlers for account-related operations.
package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"task/store"
)

// AccountHandler handles requests for account operations.
//
// It provides functionality to retrieve and display account information.
type AccountHandler struct {
	// TODO: Add AccountsStore provides access to account storage operations
}

// Serve handles requests for account operations by retrieving and printing
// account information in JSON format. It uses the AccountsStore to fetch account details
// based on the provided account ID.
//
// The method performs the following operations:
//  1. Retrieves account information using the AccountsStore
//  2. Marshals the account data into JSON
//  3. Writes the JSON response to the client
//
// Parameters:
//   - ctx: The context for the request, which may include deadlines, cancellation signals,
//     and request-scoped values
//   - accountID: The unique identifier of the account to retrieve from the store
//
// Returns:
//   - error: Returns nil on successful operation. Returns an error if:
//   - Account retrieval fails
//   - JSON marshaling fails
func (h *AccountHandler) Serve(ctx context.Context, accountID uint64) error {
	// In the real-world application we will call ctx.UserValue(key)
	// to retrieve URL parameter :account_id from the *fasthttp.RequestCtx

	accounts, err := h.AccountsStore.FindOne(ctx, store.Account{
		ID: accountID,
	})
	if err != nil {
		return fmt.Errorf("failed to find an account: %w", err)
	}

	bits, err := json.Marshal(accounts)
	if err != nil {
		return fmt.Errorf("failed to marshal an account: %w", err)
	}

	// Writing response body to client
	log.Println(string(bits))
	return nil
}
