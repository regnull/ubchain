// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "./key_registry.sol";

contract NameRegistry {
    KeyRegistry keyRegistry;

    mapping(string => bytes) public registry;

    constructor(address keyRegistryAddress) {
        keyRegistry = KeyRegistry(keyRegistryAddress);
    }

    function register(string calldata name, bytes calldata key) public {
        require(key.length == 33);
        require(bytes(name).length >= 3);
        require(bytes(name).length <= 64);

        // The key must be known.
        require(keyRegistry.exists(key));

        if (registry[name].length > 0) {
            bytes memory ownerKey = key;
            bytes memory parentKey = keyRegistry.parent(key);
            if (parentKey.length > 0) {
                ownerKey = parentKey;
            }
            address adr = address(bytes20(keccak256(ownerKey)));
            require(msg.sender == adr);
        }
        registry[name] = key;
    }

    function getKey(string calldata name) view public returns(bytes memory) {
        return registry[name];
    }
}