#!/bin/bash

cd $HOME
cargo install --version 0.9.0 ibc-relayer-cli --bin hermes --locked
sudo cp $HOME/.cargo/bin/hermes /usr/local/bin