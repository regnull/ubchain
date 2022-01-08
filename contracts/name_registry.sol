// SPDX-License-Identifier: MIT
pragma solidity ^0.8.11;

import "./key_registry.sol";

contract NameRegistry {
    KeyRegistry keyRegistry;

    mapping(string => bytes) public registry;

    event NameRegistered(string, bytes);

    constructor(address keyRegistryAddress) {
        keyRegistry = KeyRegistry(keyRegistryAddress);
    }

    function register(string calldata name, bytes calldata key) public {
        require(key.length == 33);
        require(bytes(name).length >= 3);
        require(bytes(name).length <= 64);

        // The key must be known.
        require(keyRegistry.registered(key));

        if (registry[name].length > 0) {
            // We can re-assign the name to the new owner.
            require(keyRegistry.owner(registry[name]) == msg.sender);
        } else {
            require(keyRegistry.owner(key) == msg.sender);
        }
        registry[name] = key;
        emit NameRegistered(name, key);
    }

    function getKey(string calldata name) view public returns(bytes memory) {
        return registry[name];
    }
}