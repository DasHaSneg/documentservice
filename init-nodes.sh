#!/bin/bash

NODES_NUM=3
CHAIN_ID="test"
BINARY="/home/${USER}/go/bin/documentserviced"

rm -R nodecluster

for (( i=0; i<$NODES_NUM; i++))
do
    $BINARY init "node${i}" --chain-id ${CHAIN_ID} --home "./nodecluster/node${i}"
    $BINARY keys add "validator${i}" --keyring-backend test --home "./nodecluster/node${i}"  --keyring-dir "./nodecluster"
done

for (( i=0; i<$NODES_NUM; i++))
do
    for (( j=0; j<$NODES_NUM; j++))
    do
        VADDRESS=$($BINARY keys show validator${j} -a --keyring-backend test --keyring-dir "./nodecluster")
        $BINARY add-genesis-account $VADDRESS  100000000000stake --home "./nodecluster/node${i}"
    done
done

mkdir ./nodecluster/gentx

for (( i=0; i<$NODES_NUM; i++))
do
    $BINARY gentx "validator${i}" 100000000stake --chain-id $CHAIN_ID --keyring-backend test --home "./nodecluster/node${i}" --keyring-dir "./nodecluster" --output-document "./nodecluster/gentx/val${i}.json"
done

$BINARY collect-gentxs --home "./nodecluster/node0" --gentx-dir "./nodecluster/gentx"

for (( i=1; i<$NODES_NUM; i++))
do 
    let NUM=$i+2
    sed -i "0,/192.168.0.105/s//192.168.10.$NUM/" ./nodecluster/node0/config/config.toml  
    rm "./nodecluster/node${i}/config/genesis.json"
    cp "./nodecluster/node0/config/genesis.json" "./nodecluster/node${i}/config"
done



