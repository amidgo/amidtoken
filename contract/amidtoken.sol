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
    uint public cost = 1 ether * 75 / 100000;

    mapping (address => uint) public balanceOf;
    mapping (address => mapping (address => uint)) public allowance;
    
    Phase currentPhase = Phase.seedPhase; 
    uint public phaseTokenLimit;
    address currentTokenOwner;
    mapping (address => string) public requests;
    address[] public requestAddresses;
    
    uint public startTime;
    uint currentTime;
    uint public timeDiff;

    mapping (address => bool) public whiteList;

    mapping (address => uint) public seedTokenTx;
    mapping (address => uint) public privateTokenTx;
    mapping (address => uint) public publicTokenTx;
    

    constructor(address owner_, address privateProvider_,address publicProvider_,address investor1,address investor2,address friend) {
        startTime = block.timestamp;
        currentTime = block.timestamp;
        owner = owner_;
        currentTokenOwner = owner;
        privateProvider = privateProvider_;
        publicProvider = publicProvider_;
        setUser(owner_,"owner");
        setUser(privateProvider_,"private");
        setUser(publicProvider_,"public");
        setUser(investor1,"user");
        setUser(investor2,"user");
        setUser(friend,"user");

        balanceOf[owner] += 10000000;
        seedTokenTx[owner] += 10000000;
    

        transferFrom_(owner,investor1,300000);

        transferFrom_(owner,investor2,400000);

        transferFrom_(owner,friend,200000);

    }


    function getTime() public view returns(uint){
        // return (block.timestamp - startTime)/60 + timeDiff;
        return timeDiff;
    }

    function getPhase() public view returns(string memory){
        // uint time = (block.timestamp - startTime)/60 + timeDiff;
        uint time = timeDiff;
        if (time >= 5 && time < 15) {
            return "private";
        }
        if (time >= 15) {
            return "public";
        }
        return "seed";
    }
    function changeCost(uint newValue) public {
        require(msg.sender == publicProvider,"only public provider can do this operation");
        cost = newValue;
    }

    function transferFrom_(address from,address to,uint amount) private enoughtTokens(from,amount)  {
        balanceOf[from] -= amount;
        balanceOf[to] += amount;
        setPhaseHistory(to,amount);
    }

    function transfer(address to,uint amount) public enoughtTokens(msg.sender,amount) {
        transferFrom_(msg.sender,to,amount);
    }   

    function transferFrom(address from,address to,uint amount) public enoughtTokens(from,amount) {
        allowance[from][msg.sender] -= amount;
        transferFrom_(from,to,amount);
    }

    function approve(address to,uint amount) public {
        allowance[msg.sender][to] = amount;
    }

    modifier enoughtTokens(address from, uint amount) {
        require(balanceOf[from] >= amount,"check balance");
        _;
    }

    function timeTravel() public {
        timeDiff++;
    }

    function buy(uint amount) public payable checkWhiteList enoughtTokens(currentTokenOwner,amount) checkPhaseLimit(amount) {
        uint sum = amount * cost;
        payable(currentTokenOwner).transfer(sum);
        if (sum < msg.value){
            payable(msg.sender).transfer(sum - msg.value);
        }
        transferFrom_(currentTokenOwner,msg.sender,amount);
    }

    function sendRequest(string memory _name) public {
        require(currentPhase == Phase.seedPhase,"you can send requests only in seed phase");
        requests[msg.sender] = _name;
        requestAddresses.push(msg.sender);
    }

    function handleRequest(address sender,bool isAccept) public {
        require(msg.sender == privateProvider,"only private provider can do this operation");
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

    function setPrivatePhase() public {
        if (currentPhase == Phase.seedPhase) {
            uint privatePhaseAmount = 3000000;
            currentTokenOwner = privateProvider;
            currentPhase = Phase.privatePhase;
            transferFrom_(owner,privateProvider,privatePhaseAmount);
            phaseTokenLimit = 100000;
        }
    }

    function setPublicPhase() public {
        if (currentPhase == Phase.privatePhase){
            currentTokenOwner = publicProvider;
            uint publicTokenAmount = 6000000;
            transferFrom_(privateProvider,owner,balanceOf[privateProvider]);
            currentTokenOwner = publicProvider;
            currentPhase = Phase.publicPhase;
            transferFrom_(owner,publicProvider,publicTokenAmount);
            cost = 1 ether / 1000;
            phaseTokenLimit = 5000;
        }
    }

    modifier checkWhiteList() {
        if (currentPhase == Phase.seedPhase){
            revert("Private sale not started");
        }
        if (currentPhase == Phase.privatePhase){
            require(whiteList[msg.sender],"Free sale not started");
        }
        _;
    }
    modifier checkPhaseLimit(uint value) {
        require(value<=phaseTokenLimit,"value bigger than currentPhaseLimit");
        _;
    }

    function setPhaseHistory(address to,uint value) private {
        if (currentPhase == Phase.publicPhase) {
            publicTokenTx[to] += value;
        }
        if (currentPhase == Phase.privatePhase) {
            privateTokenTx[to] += value;
        }
        if (currentPhase == Phase.seedPhase) {
            seedTokenTx[to] += value;
        }
        
    }
    
}