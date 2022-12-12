// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract NameRegistry {
    struct RegistryEntry {
        bytes publicKey;
        address owner;
        uint256 price;
    }
    mapping(string => RegistryEntry) keyRegistry;
    mapping(string => mapping(string => string)) connectorRegistry;

    event NameRegistered(string, address);
    event PriceChanged(string, uint256);
    event Sale(string, uint256);
    event ConnectorRegistered(string, string, string);

    function registerName(bytes calldata publicKey, string calldata name)
        public
    {
        require(publicKey.length == 33);
        require(keyRegistry[name].owner == address(0)); // The name must not be taken.
        require(bytes(name).length >= 1);
        require(bytes(name).length <= 64);

        keyRegistry[name].owner = msg.sender;
        keyRegistry[name].publicKey = publicKey;
        emit NameRegistered(name, msg.sender);
    }

    function lookupName(string calldata name)
        public
        view
        returns (
            address owner,
            bytes memory publicKey,
            uint256 price
        )
    {
        RegistryEntry storage e = keyRegistry[name];
        return (e.owner, e.publicKey, e.price);
    }

    function changePrice(string calldata name, uint256 price) public {
        require(keyRegistry[name].owner == msg.sender); // Must be the owner.
        keyRegistry[name].price = price;
        emit PriceChanged(name, price);
    }

    function buy(string calldata name, bytes memory publicKey) public {
        require(keyRegistry[name].price > 0); // The name must be for sale.
        uint256 price = keyRegistry[name].price;
        payable(keyRegistry[name].owner).transfer(price);
        keyRegistry[name].owner = address(0);
        keyRegistry[name].publicKey = publicKey;
        keyRegistry[name].price = 0;
        emit Sale(name, price);
    }

    function registerConnector(
        string calldata name,
        string calldata protocol,
        string calldata location
    ) public {
        require(keyRegistry[name].owner == msg.sender); // Must be the owner.
        connectorRegistry[name][protocol] = location;
        emit ConnectorRegistered(name, protocol, location);
    }

    function lookupConnector(string calldata name, string calldata protocol)
        public
        view
        returns (string memory)
    {
        return (connectorRegistry[name][protocol]);
    }
}
