#!/bin/bash

cd $HOME
git clone https://github.com/defund-labs/ts-relayer
cd ts-relayer
npm install
npm run build --location=global
npm link