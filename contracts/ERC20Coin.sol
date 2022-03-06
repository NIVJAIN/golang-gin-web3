pragma solidity >=0.4.0 <0.9.0;
// pragma solidity 0.4.0;
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";


contract ERC20Coin is ERC20 {
    string public setMessage = "Hello Solidity";
    constructor() ERC20("Scaffold ETH Token", "SET") {
        _mint(msg.sender, 55555 * 10 ** 18); //1000000000000000000000n
        // _mint(msg.sender, 5555 * 1 ** 1);
    }

   uint256 public count = 0;
    function Hello() public pure returns (string memory) {
        return "Hello World";
    }
    function Greet(string memory str) public pure returns (string memory) {
        return str;
    }
    function setText(string memory _setText) public {
        setMessage = _setText;
    }
    function getText() public view returns(string memory) {
        return(setMessage);
    }
       // Function that increments our counter
    function increment() public {
        count += 1;
    }

    // Not necessary getter to get the count value
    function getCount() public view returns (uint256) {
        return count;
    }
     // Function that increments our counter
    function decrement() public {
        count--;
    }
}