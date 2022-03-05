
../truffle-upgradable-proxypatterns/node_modules/solc/solcjs --optimize --abi ./contracts/MySmartContract.sol -o build

../truffle-upgradable-proxypatterns/node_modules/solc/solcjs --optimize --bin ./contracts/MySmartContract.sol -o build

abigen --abi=./build/__contracts_MySmartContract_sol_MySmartContract.abi --bin=./build/__contracts_MySmartContract_sol_MySmartContract.bin --pkg=api --out=./api/MySmartContract.go
