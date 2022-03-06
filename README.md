### Smart Contract deployment via truffle
```
ganache-cli -p 7545 -m "lemon example with key phrase stereo"
truffle compile --all
truffle migrate --network ganache
cd middlewares/blockchain
check metadata.abi file exists.
Generate a wrapper file in go with metadata.abi file using below command
abigen --abi=./middlewares/blockchain/ganache/metadata.abi --pkg=api --out=./api/ERC20Coin2.go
.env

```


###  SOLC ABIGEN GO
```

../truffle-upgradable-proxypatterns/node_modules/solc/solcjs --optimize --abi ./contracts/MySmartContract.sol -o build

../truffle-upgradable-proxypatterns/node_modules/solc/solcjs --optimize --bin ./contracts/MySmartContract.sol -o build

abigen --abi=./build/__contracts_MySmartContract_sol_MySmartContract.abi --bin=./build/__contracts_MySmartContract_sol_MySmartContract.bin --pkg=api --out=./api/MySmartContract.go


../truffle-upgradable-proxypatterns/node_modules/solc/solcjs --optimize --abi ./contracts/ERC20Coin.sol -o build
../truffle-upgradable-proxypatterns/node_modules/solc/solcjs --optimize --bin ./contracts/ERC20Coin.sol -o build
abigen --abi=./build/__contracts_MySmartContract_sol_MySmartContract.abi --bin=./build/__contracts_MySmartContract_sol_MySmartContract.bin --pkg=api --out=./api/MySmartContract.go

cat <contract.json> | abigen --abi -  --pkg <packagename> --out <output.go>
cat build/contracts/MySmartContract.json | abigen --abi - --pkg truffleapi --out=./api/TruffleMySmartContract.go
abigen --abi=./build/contracts/MySmartContract.json --pkg=multichainresolver --out=multi_chain_resolver.go
abigen --abi=./build/myAbi.abi --pkg=myabi --out=myabi.go
```

## ABIGEN
```
../truffle-upgradable-proxypatterns/node_modules/solc/solcjs --optimize --bin ./contracts/ERC20Coin.sol -o build
abigen --abi=./middlewares/blockchain/ganache/metadata.abi --pkg=api --out=./api/ERC20Coin.go
```
