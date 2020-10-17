# Entry point for setup

mkdir ~/.provision
cd ~/.provision
curl -o scripts.zip https://codeload.github.com/pwae/mac/zip/main 
unzip scripts.zip
cd mac-main/
chmod +x run
./run
