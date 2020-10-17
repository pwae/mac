#!/bin/bash
# Entry point for setup

xcode-select --install

mkdir ~/.provision
cd ~/.provision
git clone https://github.com/pwae/mac git
cd git
chmod +x scripts/run
./scripts/run
