refresh:
	sudo rm -rf ./eth-net/node1/*
	sudo rm -rf ./eth-net/node2/*
	docker run --rm -v ${PWD}/eth-net:/sources ethereum/client-go:alltools-v1.11.2 geth --datadir /sources/node1 init /sources/genesis.json
	docker run --rm -v ${PWD}/eth-net:/sources ethereum/client-go:alltools-v1.11.2 geth --datadir /sources/node2 init /sources/genesis.json
	sudo cp ./eth-net/keystore/* ./eth-net/node1/keystore
	sudo cp ./eth-net/keystore/* ./eth-net/node2/keystore

amidtoken:
	make refresh
	docker run --rm -v ${PWD}/contract:/sources ethereum/solc:0.8.0 --abi --bin /sources/amidtoken.sol -o /sources --overwrite 
	docker run --rm -v ${PWD}/contract:/sources ethereum/client-go:alltools-v1.11.2 abigen --abi /sources/AmidToken.abi --bin /sources/AmidToken.bin --pkg contract --out /sources/AmidToken.go
	sudo chmod 777 contract/*
	docker-compose up -d 
stop:
	docker-compose down
	docker image rm amidtoken_api