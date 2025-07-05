package inmem

import "task/store"

var AccountsIndex = map[uint64]store.Account{
	1: {
		ID:        1,
		Asset:     "Bitcoin",
		Balance:   0.001,
		IsRecount: true,
	},
	2: {
		ID:        2,
		Asset:     "Ethereum",
		Balance:   0.00001,
		IsRecount: true,
	},
}
