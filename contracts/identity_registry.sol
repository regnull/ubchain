// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

contract IdentityRegistry {

    struct KeyEntry {
        bool initialized;
        bool disabled;
    }

    struct NameEntry {
        bool initialized;
        bytes publicKey;
        address owner;
    }

    struct AddressEntry {
        bool initialized;
        string destination;
    }

    mapping(bytes => KeyEntry) public keyRegistry;
    mapping(string => NameEntry) public nameRegistry;
    mapping(string => AddressEntry) public addressRegistry;

    function registerKey(bytes calldata publicKey) public {
        require(publicKey.length == 33);
        require(!keyRegistry[publicKey].initialized);
        keyRegistry[publicKey].initialized = true;    
    }

    function registerName(string calldata name, bytes calldata publicKey, string calldata messagingAddress) public {
        require(bytes(name).length <= 64);
        require(publicKey.length == 33);
        require(bytes(messagingAddress).length <= 128);

        require(keyRegistry[publicKey].initialized);
        require(!keyRegistry[publicKey].disabled);
        require(!nameRegistry[name].initialized || nameRegistry[name].owner == msg.sender);

        nameRegistry[name].initialized = true;
        nameRegistry[name].publicKey = publicKey;
        nameRegistry[name].owner = msg.sender;
    }

    function registerAddress(string calldata name, string calldata destination) public {
        require(nameRegistry[name].owner == msg.sender);
        addressRegistry[name].initialized = true;
        addressRegistry[name].destination = destination;
    }
}