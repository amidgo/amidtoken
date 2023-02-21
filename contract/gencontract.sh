#bin/bash
solc --abi --bin amidtoken.sol -o . --overwrite && abigen --abi AmidToken.abi --bin AmidToken.bin --pkg contract --out AmidToken.go