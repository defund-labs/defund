#!/bin/bash

# Install the specialized, ics-27 enabled gaiad binary
cd $HOME
yes | rm -r gaia
git clone https://github.com/cosmos/gaia
cd gaia
make install
sudo cp $GOPATH/bin/gaiad /usr/local/bin