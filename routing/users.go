package routing

import (
	"math/big"

	"github.com/amidgo/amidtoken/variables"
	"github.com/ethereum/go-ethereum/common"
)

type Sender struct {
	Address *common.Address
}
type User struct {
	Sender
	Role    string
	Seed    *big.Int
	Private *big.Int
	Public  *big.Int
}

func AllUsers() []*User {
	users := make([]*User, 0)
	var index int64
	for {
		addr, err := variables.Contract.UserAddresses(variables.DefaultCallOpts(), big.NewInt(index))
		index++
		if err != nil {
			break
		}
		role, _ := variables.Contract.Users(variables.DefaultCallOpts(), addr)
		sToken, _ := variables.Contract.SeedTokenTx(variables.DefaultCallOpts(), addr)
		prToken, _ := variables.Contract.PrivateTokenTx(variables.DefaultCallOpts(), addr)
		pToken, _ := variables.Contract.PublicTokenTx(variables.DefaultCallOpts(), addr)
		users = append(users, &User{Sender: Sender{&addr}, Role: role, Seed: sToken, Private: prToken, Public: pToken})
	}
	return users
}
