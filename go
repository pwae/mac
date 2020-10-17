#!/bin/bash
# Entry point for setup

# Install command line utils
touch /tmp/.com.apple.dt.CommandLineTools.installondemand.in-progress;
PROD=$(softwareupdate -l |
  grep -B 1 -E 'Command Line Tools' |
  awk -F'*' '/^ *\\*/ {print \$2}' |
  sed -e 's/^ *Label: //' -e 's/^ *//' |
  sort -V |
  tr -d '\n')
softwareupdate -i "$PROD";

mkdir ~/.provision
cd ~/.provision
git clone https://github.com/pwae/mac git
cd git
chmod +x scripts/run
./scripts/run

