#!/bin/bash
# Entry point for setup

# Check if command line utils already installed
xcode_installed=$(xcode-select -p 1>/dev/null;echo -n $?)
if test $xcode_installed -eq 2; then
  # Install command line utils
  touch /tmp/.com.apple.dt.CommandLineTools.installondemand.in-progress;
  PROD=$(softwareupdate -l |
    grep -B 1 -E 'Command Line Tools' |
    awk -F'*' '/^ *\\*/ {print $2}' |
    sed -e 's/^ *Label: //' -e 's/^ *//' |
    sort -V |
    tail -n1 |
    tr -d '\n')
  softwareupdate -i "$PROD";
  rm /tmp/.com.apple.dt.CommandLineTools.installondemand.in-progress
fi

mkdir ~/.provision
cd ~/.provision
git clone https://github.com/pwae/mac git
cd git
chmod +x scripts/run
./scripts/run

