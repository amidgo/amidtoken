#bin/bash
docker run ethereum/solc:0.8.0 --abi --bin contract/amidtoken.sol -o contract --overwrite && abigen --abi contract/AmidToken.abi --bin contract/AmidToken.bin --pkg contract --out contract/AmidToken.go