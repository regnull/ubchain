// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

contract KeyRegistry {
    struct KeyEntry {
        bool disabled;
        address owner;
    }

    mapping(bytes => KeyEntry) public registry;

    event KeyRegistered(bytes, address);
    event KeyOwnerChanged(bytes, address);
    event KeyDisabled(bytes);

    function register(bytes calldata publicKey) public {
        require(publicKey.length == 33);
        require(registry[publicKey].owner == address(0));

        registry[publicKey].owner = msg.sender; 
        emit KeyRegistered(publicKey, msg.sender);
    }

    function registered(bytes calldata publicKey) view public returns(bool) {
        return registry[publicKey].owner != address(0);
    }

    function disabled(bytes calldata publicKey) view public returns(bool) {
        return registry[publicKey].disabled;
    }

    function owner(bytes calldata publicKey) view public returns(address) {
        return registry[publicKey].owner;
    }

    function changeOwner(bytes calldata publicKey, address newOwner) public {
        require(publicKey.length == 33);
        require(registry[publicKey].owner == msg.sender);
 
        registry[publicKey].owner = newOwner;
        emit KeyOwnerChanged(publicKey, newOwner);
    }

    function disable(bytes calldata publicKey) public {
        require(publicKey.length == 33);
        require(registry[publicKey].owner == msg.sender);

        registry[publicKey].disabled = true;
        emit KeyDisabled(publicKey);
    }
}