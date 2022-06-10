#!/usr/bin/env bash
CHAIN_ID=test

if [ -f ~/.documentservice/config/genesis.json ]; then
    echo "run node"
    documentserviced start --home ~/.documentservice
else
	echo "Config empty"
fi
