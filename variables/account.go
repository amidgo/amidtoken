package variables

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
)

func ImportAccount(addr common.Address) *accounts.Account {
	for _, a := range Keystore.Accounts() {
		if a.Address == addr {
			return &a
		}
	}
	return nil
}
