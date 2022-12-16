// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract NameRegistry {
    struct RegistryEntry {
        bytes publicKey;
        address owner;
        uint256 price;
    }
    mapping(string => RegistryEntry) keyRegistry;
    mapping(string => mapping(string => string)) config;

    // A new name was just registered.
    event NameRegistered(string name, address owner);

    // Public key was updated for this name.
    event PublicKeyUpdated(string name);

    event NameOwnershipUpdated(string name, address newOwner);

    // Name's price was changed - maybe it was listed for sale, or was delisted.
    // Price of zero means the name is not for sale.
    event PriceUpdated(string name, uint256 price);

    // A name was sold.
    event Sale(string name, uint256 price, address newOwner);

    // Configuration was updated.
    event ConfigUpdated(string name, string configName, string configValue);

    // Register a name to an owner. The name must not be already registered.
    function registerName(bytes calldata publicKey, string calldata name)
        public
    {
        require(publicKey.length == 33, "invalid public key");
        require(keyRegistry[name].owner == address(0), "name is already registered"); // The name must not be taken.
        require(bytes(name).length >= 1, "name can't be empty");
        require(bytes(name).length <= 64, "name is too long");
        require(bytes(name)[0] != "-", "name can't start with a dash"); // Name cannot start with a dash.

        // Verify the name.
        bytes memory baseBytes = bytes(name);
        bool onlySymbols = true;
        for (uint256 i = 0; i < baseBytes.length; i++) {
            bytes1 b = baseBytes[i];
            // We only allow ASCII letters and numbers in the name, plus characters '-' and '_'.
            if (
                !(b == 0x2D ||
                    (b >= 0x30 && b <= 0x39) ||
                    b == 0x5F ||
                    (b >= 0x61 && b <= 0x7A))
            ) {
                revert("invalid name");
            }

            if (b != "-" && b != "_") {
                onlySymbols = false;
            }
        }
        require(!onlySymbols, "name must contain at least one letter or digit");

        keyRegistry[name].owner = msg.sender;
        keyRegistry[name].publicKey = publicKey;
        emit NameRegistered(name, msg.sender);
    }

    // Lookup information associated with a name.
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

    // Update public key for already registered name.
    function updatePublicKey(bytes calldata publicKey, string calldata name)
        public
    {
        require(publicKey.length == 33, "invalid public key");
        require(keyRegistry[name].owner == msg.sender, "not the owner"); // The name must not be taken.
        keyRegistry[name].publicKey = publicKey;
        emit PublicKeyUpdated(name);
    }

    // Transfer name ownership to a different owner.
    function updateOwnership(string calldata name, address newOwner) public {
        require(keyRegistry[name].owner == msg.sender, "not the owner"); // Must be the owner.
        keyRegistry[name].owner = newOwner;
        emit NameOwnershipUpdated(name, newOwner);
    }

    // Change name's price. If a price is non-zero, anyone can pay and assume ownership of a name.
    // A price of zero means the name is not for sale.
    function updatePrice(string calldata name, uint256 price) public {
        require(keyRegistry[name].owner == msg.sender, "not the owner"); // Must be the owner.
        keyRegistry[name].price = price;
        emit PriceUpdated(name, price);
    }

    // Buy a name. The name must be listed for sale (price is greater than zero).
    function buyName(string calldata name, bytes memory publicKey) public payable {
        require(keyRegistry[name].price > 0, "not for sale"); // The name must be for sale.
        require(keyRegistry[name].price <= msg.value, "insufficient value");
        uint256 price = keyRegistry[name].price;
        payable(keyRegistry[name].owner).transfer(msg.value);
        keyRegistry[name].owner = msg.sender;
        keyRegistry[name].publicKey = publicKey;
        keyRegistry[name].price = 0;
        emit Sale(name, price, msg.sender);
    }

    // Update configuration. The configuration is just a name/value pair, where both name and 
    // value are strings.
    function updateConfig(
        string calldata name,
        string calldata configName,
        string calldata configValue
    ) public {
        require(keyRegistry[name].owner == msg.sender, "not the owner"); // Must be the owner.
        config[name][configName] = configValue;
        emit ConfigUpdated(name, configName, configValue);
    }

    // Lookup configuration.
    function lookupConfig(string calldata name, string calldata configName)
        public
        view
        returns (string memory)
    {
        return (config[name][configName]);
    }
}
