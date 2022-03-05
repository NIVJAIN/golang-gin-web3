pragma solidity >=0.4.0 <0.8.0;
// pragma solidity 0.4.0;

contract MySmartContract {
    function Hello() public pure returns (string memory) {
        return "Hello World";
    }
    function Greet(string memory str) public pure returns (string memory) {
        return str;
    }
}