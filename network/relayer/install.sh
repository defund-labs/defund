#!/bin/bash

cd $HOME
cargo install ibc-relayer-cli --bin hermes --locked
sudo cp $HOME/.cargo/bin/hermes /usr/local/bin