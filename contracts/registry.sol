// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract NameRegistry {
    struct RegistryEntry {
        bytes publicKey;
        address owner;
        uint price;
        mapping(string => string) connectors;
    }
    mapping(string => RegistryEntry) keyRegistry;

    function registerName(bytes calldata publicKey, string calldata name) public {
        require(publicKey.length == 33);
        require(keyRegistry[name].owner == address(0));  // The name must not be taken.
        require(bytes(name).length >= 3);
        require(bytes(name).length <= 64);

        keyRegistry[name].owner = address(0);
        keyRegistry[name].publicKey = publicKey;
   }

   function lookup(string calldata name) view public returns(address owner, bytes calldata publicKey, uint price) {
        RegistryEntry e = keyRegistry[name];
        return (e.owner, e.publicKey, e.price);
   }
}