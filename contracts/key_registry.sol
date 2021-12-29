// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

contract KeyRegistry {
    struct KeyEntry {
        bool initialized;
        bool disabled;
        bytes parent;
    }

    mapping(bytes => KeyEntry) public registry;

    event KeyRegistered(bytes);
    event KeyParentRegistered(bytes, bytes);
    event KeyDisabled(bytes);

    function register(bytes calldata publicKey) public {
        require(publicKey.length == 33);
        require(!registry[publicKey].initialized);

        registry[publicKey].initialized = true; 
        emit KeyRegistered(publicKey);
    }

    function registerParent(bytes calldata publicKey, bytes calldata parentKey) public {
        require(publicKey.length == 33);
        require(parentKey.length == 33);
        require(registry[publicKey].initialized);
        require(registry[parentKey].initialized);

        address adr = address(bytes20(keccak256(publicKey)));
        require(adr == msg.sender);
        registry[publicKey].parent = parentKey;
        emit KeyParentRegistered(publicKey, parentKey);
    }

    function disable(bytes calldata publicKey) public {
        require(publicKey.length == 33);
        require(registry[publicKey].initialized);

        // If a parent is registered, only the parent can disable this key.
        if (registry[publicKey].parent.length == 0) {
            address adr = address(bytes20(keccak256(publicKey)));
            require(adr == msg.sender);
        } else {
            address adr = address(bytes20(keccak256(registry[publicKey].parent)));
            require(adr == msg.sender);
        }

        registry[publicKey].disabled = true;
        emit KeyDisabled(publicKey);
    }
}