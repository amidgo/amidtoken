package variables

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func TransactOpts(from common.Address, value *big.Int) (*bind.TransactOpts, error) {
	acc := ImportAccount(from)
	chainId, _ := Client.ChainID(context.Background())
	auth, err := bind.NewKeyStoreTransactorWithChainID(Keystore, *acc, chainId)
	auth.From = from
	auth.Value = value
	gasPrice, _ := Client.SuggestGasPrice(context.Background())
	auth.GasPrice = gasPrice
	return auth, err
}

func DefaultTransactOpts() *bind.TransactOpts {
	tx, err := TransactOpts(Owner, big.NewInt(0))
	if err != nil {
		fmt.Println("transaction")
		log.Fatal(err, tx)
	}
	return tx
}
