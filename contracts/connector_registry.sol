// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "./key_registry.sol";
import "./name_registry.sol";

contract ConnectorRegistry {
    KeyRegistry keyRegistry;
    NameRegistry nameRegistry;

    mapping(string => mapping(string => string)) public registry;

    event ConnectorRegistered(string, string, string);

    constructor(address keyRegistryAddress, address nameRegistryAddress) {
        keyRegistry = KeyRegistry(keyRegistryAddress);
        nameRegistry = NameRegistry(nameRegistryAddress);
    }
   
   function register(string calldata name, string calldata protocol, string calldata location) public {
       // Make sure the name is registered.
       bytes memory key = nameRegistry.getKey(name);
       require(key.length == 33);
       require(keyRegistry.owner(key) == msg.sender);

       registry[name][protocol] = location;
       emit ConnectorRegistered(name, protocol, location);
   }
}