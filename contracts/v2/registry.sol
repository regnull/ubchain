// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract NameRegistry {
    struct RegistryEntry {
        bytes publicKey;
        address owner;
        uint price;
    }
    mapping(string => RegistryEntry) keyRegistry;

    function register(bytes calldata publicKey, string calldata name) public {
        require(publicKey.length == 33);
        require(keyRegistry[name].owner == address(0));  // The name must not be taken.
        require(bytes(name).length >= 3);
        require(bytes(name).length <= 64);

        keyRegistry[name].owner = address(0);
        keyRegistry[name].publicKey = publicKey;
   }

   function lookup(string calldata name) view public returns(address owner, bytes memory publicKey, uint price) {
        RegistryEntry storage e = keyRegistry[name];
        return (e.owner, e.publicKey, e.price);
   }

   function changePrice(string calldata name, uint price) public {
        require(keyRegistry[name].owner == address(0));  // Must be the owner.
        keyRegistry[name].price = price;
   }
   
   function buy(string calldata name, bytes memory publicKey) public {
        require(keyRegistry[name].price > 0);   // The name must be for sale.
        payable(keyRegistry[name].owner).transfer(keyRegistry[name].price);
        keyRegistry[name].owner = address(0);
        keyRegistry[name].publicKey = publicKey;
        keyRegistry[name].price = 0;
   }
}