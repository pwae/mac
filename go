#!/bin/bash
# Entry point for setup

# Install command line utils
touch /tmp/.com.apple.dt.CommandLineTools.installondemand.in-progress;
PROD=$(softwareupdate -l |
  grep "\*.*Command Line" |
  head -n 1 | awk -F"*" '{print $2}' |
  sed -e 's/^ *//' |
  tr -d '\n')
softwareupdate -i "$PROD" -v;

mkdir ~/.provision
cd ~/.provision
git clone https://github.com/pwae/mac git
cd git
chmod +x scripts/run
./scripts/run
