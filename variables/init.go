package variables

import (
	"log"

	"github.com/amidgo/amidtoken/contract"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const ABI = `[{"inputs":[{"internalType":"address","name":"owner_","type":"address"},{"internalType":"address","name":"privateProvider_","type":"address"},{"internalType":"address","name":"publicProvider_","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"approve","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"buy","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"cost","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"decimals","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getPhase","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getTime","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"sender","type":"address"},{"internalType":"bool","name":"isAccept","type":"bool"}],"name":"handleRequest","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"phaseTokenLimit","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"privateProvider","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"publicProvider","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"requestAddresses","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"requests","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"_name","type":"string"}],"name":"sendRequest","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"startTime","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"timeDiff","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"timeTravel","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"totalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transfer","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"whiteList","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`

const (
	Node1         = "http://0.0.0.0:1111"
	Node2         = "http://0.0.0.0:2222"
	Node1Keystore = "./eth-net/node1/keystore"
	Node2Keystore = "./eth-net/node2/keystore"
)

var (
	Owner           = common.HexToAddress("0x1A6be8c4C1DE78ae1Ef935D8B1344628f53BFad0")
	OwnerPWD        = "owner"
	PrivateProvider = common.HexToAddress("0xcb25dD2785f4647eEE4e5a1A8C32166aa408e652")
	PrivateProvPWD  = "private"
	PublicProvider  = common.HexToAddress("0x3471DABD95d37CF8fcd0cf8971E4018BEC489b86")
	PublicProvPWD   = "public"
	Investor1       = common.HexToAddress("0xb935DA91dbf4F7637C0EfbC682F359441eA38EB6")
	Investor1PWD    = "investor1"
	Investor2       = common.HexToAddress("0x20E47B326bb755963E6998a525E08eEFacA15718")
	Investor2PWD    = "investor2"
	BestFriend      = common.HexToAddress("0x9CD2200ec60b6BF23047c0fB44E8498ecf6097A7")
	BestFriendPWD   = "friend"
	Miner           = common.HexToAddress("0xf51878d18c731033B51e5367de79ae9443815246")
	MinerPWD        = "miner"
)

var (
	Client   *ethclient.Client
	Keystore *keystore.KeyStore
	Contract *contract.Contract
)

func Init() {

	c, err := ethclient.Dial(Node2)
	if err != nil {
		log.Fatal(err)
	}
	Client = c
	k := keystore.NewKeyStore(Node2Keystore, keystore.StandardScryptN, keystore.StandardScryptP)
	Keystore = k
	Keystore.Unlock(*ImportAccount(Owner), OwnerPWD)
	_, _, cn, err := contract.DeployContract(DefaultTransactOpts(), Client, Owner, PrivateProvider, PublicProvider, Investor1, Investor2, BestFriend)
	if err != nil {
		log.Fatal(err)
	}
	Contract = cn
}
