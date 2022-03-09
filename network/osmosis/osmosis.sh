#!/bin/bash

# Install the special ics-27 enabled osmosis binary
cd $HOME
yes | rm -r osmosis
git clone https://github.com/schnetzlerjoe/osmosis
cd osmosis
make install
sudo cp $GOPATH/bin/osmosisd /usr/local/bin
