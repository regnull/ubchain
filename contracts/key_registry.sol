// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

contract KeyRegistry {
    struct KeyEntry {
        bool initialized;
        bool disabled;
        address owner;
    }

    mapping(bytes => KeyEntry) public registry;

    event KeyRegistered(bytes);
    event KeyDisabled(bytes);

    function register(bytes calldata publicKey) public {
        require(publicKey.length == 33);
        require(!registry[publicKey].initialized);

        registry[publicKey].initialized = true; 
        registry[publicKey].owner = msg.sender;   
        emit KeyRegistered(publicKey);
    }

    function disable(bytes calldata publicKey) public {
        require(publicKey.length == 33);
        require(registry[publicKey].initialized);
        require(registry[publicKey].owner == msg.sender);

        registry[publicKey].disabled = true;
        emit KeyDisabled(publicKey);
    }
}