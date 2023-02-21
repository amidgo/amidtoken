#bin/bash
rm -r eth-net/node1/*
rm -r eth-net/node2/*
geth --datadir eth-net/node1 init eth-net/genesis.json
geth --datadir eth-net/node2 init eth-net/genesis.json
cp eth-net/keystore/* eth-net/node1/keystore
cp eth-net/keystore/* eth-net/node2/keystore
