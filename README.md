sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc

abigen --sol vinTrack.sol --pkg main --out token.go

