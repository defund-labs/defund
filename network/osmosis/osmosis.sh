#!/bin/bash

# Install Osmosis
cd $HOME
yes | rm -r osmosis
git clone https://github.com/osmosis-labs/osmosis
cd osmosis
git checkout v10.0.0
make install