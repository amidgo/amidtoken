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
	Role               string   `json:"role"`
	SeedTokenAmount    *big.Int `json:"seedTokenAmount"`
	PrivateTokenAmount *big.Int `json:"privateTokenAmount"`
	PublicTokenAmount  *big.Int `json:"publicTokenAmount"`
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
		users = append(users, &User{Sender: Sender{&addr}, Role: role, SeedTokenAmount: sToken, PrivateTokenAmount: prToken, PublicTokenAmount: pToken})
	}
	return users
}
