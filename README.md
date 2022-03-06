### This Repo has both truffle and golang
```
1. truffle for smart contract deployment.
2. golang with gin fasthttp framework for rest API's
```

### Smart Contract deployment via truffle
```
1. ganache-cli -p 7545 -m "lemon example with key phrase stereo"
2. truffle compile --all
3. truffle migrate --network ganache
4. cd middlewares/blockchain
5. check metadata.abi file exists.
6. Generate a wrapper file in go with metadata.abi file using below command
7. abigen --abi=./middlewares/blockchain/ganache/metadata.abi --pkg=api --out=./api/ERC20Coin.go
8. Final stage update your .env with your own private key and the ganache-cli port and chainid
    ETH_CLIENT_URL="http://127.0.0.1:7545"
    PRIVATE_KEY="0000000000000000000000000000000000000000000000000000
    SMART_CONTRACT_ADDRESS="0xF036e5fE5a1310D1f08196288eDBf3EC1Ff638A8"
    CHAIN_ID=1337
```

### How to Run the application
```
1. sh zzz_run.sh 
  or 
2. go build && ./golang-gin-web3
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