// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;


enum Phase {seedPhase, privatePhase, publicPhase}


struct Balance {
    uint privateTokens;
    uint publicTokens;
}

contract AmidToken {


    mapping (address => string) public users;

    address[] public userAddresses;

    function setUser(address addr,string memory role) public {
        users[addr] = role;
        userAddresses.push(addr);
    }

    address public owner;
    address public privateProvider;
    address public publicProvider;

    uint public totalSupply = 10000000; 
    string public name = "AmidToken";
    string public symbol = "amdt"; 
    uint public decimals = 12;
    uint public cost;

    mapping (address => Balance) public balanceOf;
    mapping (address => mapping (address => Balance)) public allowance;
    
    Phase currentPhase = Phase.seedPhase; 
    uint public phaseTokenLimit;
    address currentTokenOwner;
    mapping (address => string) public requests;
    address[] public requestAddresses;
    
    uint public startTime;
    uint currentTime;
    uint public timeDiff;

    mapping (address => bool) public whiteList;

    constructor(address owner_,address privateProvider_,address publicProvider_) {
        startTime = block.timestamp;
        currentTime = block.timestamp;
        owner = owner_;
        currentTokenOwner = owner;
        privateProvider = privateProvider_;
        publicProvider = publicProvider_;
        setUser(owner_,"owner");
        setUser(privateProvider_,"private");
        setUser(publicProvider_,"public");
        balanceOf[owner].privateTokens = totalSupply * 40 / 100;
        balanceOf[owner].publicTokens = totalSupply * 60 / 100;
        cost = 1 ether * 75 / 100000;
    }


    function getTime() public view returns(uint){
        return (block.timestamp - startTime)/60 + timeDiff;
    }

    function getPhase() public view returns(string memory){
        uint time = (block.timestamp - startTime)/60 + timeDiff;
        if (time >= 5 && time < 10) {
            return "private";
        }
        if (time >= 10) {
            return "public";
        }
        return "seed";
    }
    function changeCost(uint newValue) public {
        cost = newValue;
    }

    function transferPrivate(address to,uint amount) public checkPhase enoughtTokensPrivate(msg.sender,to,amount) checkToken(amount) {
        balanceOf[msg.sender].privateTokens -= amount;
        balanceOf[to].privateTokens += amount;
    }   

    function transferPublic(address to,uint amount) public checkPhase enoughtTokensPublic(msg.sender,to,amount) checkToken(amount) {
        balanceOf[msg.sender].publicTokens -= amount;
        balanceOf[to].publicTokens += amount;
    }   

    function transferFromPrivate(address from,address to,uint amount) public checkPhase enoughtTokensPrivate(from,to,amount)  {
        allowance[from][to].privateTokens -= amount;
        balanceOf[msg.sender].privateTokens -= amount;
        balanceOf[to].privateTokens += amount;
    }

    function transferFromPublic(address from,address to,uint amount) public checkPhase enoughtTokensPublic(from,to,amount)  {
        allowance[from][to].publicTokens -= amount;
        balanceOf[msg.sender].publicTokens -= amount;
        balanceOf[to].publicTokens += amount;
    }

    function approvePrivate(address from,address to,uint amount) public {
        allowance[from][to].privateTokens = amount;
    }

    function approvePublic(address from,address to,uint amount) public {
        allowance[from][to].publicTokens = amount;
    }

    modifier enoughtTokensPrivate(address from, address to, uint amount) {
        require(balanceOf[from].privateTokens > amount,"check balance");
        _;
    }

     modifier enoughtTokensPublic(address from, address to, uint amount) {
        require(balanceOf[from].publicTokens > amount,"check balance");
        _;
    }

    function timeTravel() public {
        timeDiff++;
    }

    function buyPrivate(uint amount) public payable checkPhase checkWhiteList{
        uint sum = amount * cost;
        payable(currentTokenOwner).transfer(sum);
        balanceOf[msg.sender].privateTokens += amount;
        if (sum < msg.value){
            payable(msg.sender).transfer(sum - msg.value);
        }
    }
    function buyPublic(uint amount) public payable checkPhase checkWhiteList{
         uint sum = amount * cost;
        payable(currentTokenOwner).transfer(sum);
        balanceOf[msg.sender].publicTokens += amount;
        if (sum < msg.value){
            payable(msg.sender).transfer(sum - msg.value);
        }
    }

    function sendRequest(string memory _name) checkPhase public {
        requests[msg.sender] = _name;
        requestAddresses.push(msg.sender);
    }

    function handleRequest(address sender,bool isAccept) public {
        whiteList[sender] = isAccept;
        requests[sender] = "";
        for (uint i = 0; i < requestAddresses.length; i++){
            if (sender == requestAddresses[i]){
                requestAddresses[i] = requestAddresses[requestAddresses.length - 1];
                requestAddresses.pop();
                return;
            }
        }
    }

    modifier checkToken(uint amount) {
        require(phaseTokenLimit >= amount);
        _;
    }

    modifier checkPhase() {
        currentTime = (block.timestamp - startTime)/60 + timeDiff;
        if (currentTime >= 5 && currentTime < 10) {
            balanceOf[privateProvider].privateTokens = balanceOf[owner].privateTokens;
            balanceOf[owner].privateTokens = 0;
            currentTokenOwner = privateProvider;
            currentPhase = Phase.privatePhase;
            phaseTokenLimit = 100000;
        }
        if (currentTime >= 10) {
            currentTokenOwner = publicProvider;
            balanceOf[publicProvider].publicTokens = balanceOf[owner].publicTokens;
            balanceOf[owner].privateTokens += balanceOf[privateProvider].privateTokens;
            balanceOf[privateProvider].privateTokens = 0;
            currentPhase = Phase.publicPhase;
            cost = 1 ether / 1000;
            phaseTokenLimit = 5000;
        }
        _;
    }

    modifier checkWhiteList() {
        if (currentPhase != Phase.publicPhase) {
            require(whiteList[msg.sender],"Free sale not started");
        }
        _;
    }

    
}