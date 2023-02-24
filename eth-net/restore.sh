#bin/bash
rm -r node1/*
rm -r node2/*
geth --datadir node1 init genesis.json
geth --datadir node2 init genesis.json
cp keystore/* node1/keystore
cp keystore/* node2/keystore
