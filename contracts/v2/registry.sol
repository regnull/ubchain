// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract NameRegistry {
    struct RegistryEntry {
        bytes publicKey;
        address owner;
        uint256 price;
    }
    mapping(string => RegistryEntry) keyRegistry;
    mapping(string => mapping(string => string)) connectorRegistry;

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

    // A connector was registered.
    event ConnectorRegistered(string name, string protocol, string location);

    // Register a name to an owner. The name must not be already registered.
    function registerName(bytes calldata publicKey, string calldata name)
        public
    {
        require(publicKey.length == 33);
        require(keyRegistry[name].owner == address(0)); // The name must not be taken.
        require(bytes(name).length >= 1);
        require(bytes(name).length <= 64);
        require(bytes(name)[0] != "-"); // Name cannot start with a dash.

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
        require(!onlySymbols);

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
        require(publicKey.length == 33);
        require(keyRegistry[name].owner == msg.sender); // The name must not be taken.
        keyRegistry[name].publicKey = publicKey;
        emit PublicKeyUpdated(name);
    }

    // Transfer name ownership to a different owner.
    function updateOwnership(string calldata name, address newOwner) public {
        require(keyRegistry[name].owner == msg.sender); // Must be the owner.
        keyRegistry[name].owner = newOwner;
        emit NameOwnershipUpdated(name, newOwner);
    }

    // Change name's price. If a price is non-zero, anyone can pay and assume ownership of a name.
    // A price of zero means the name is not for sale.
    function updatePrice(string calldata name, uint256 price) public {
        require(keyRegistry[name].owner == msg.sender); // Must be the owner.
        keyRegistry[name].price = price;
        emit PriceUpdated(name, price);
    }

    // Buy a name. The name must be listed for sale (price is greater than zero).
    function buy(string calldata name, bytes memory publicKey) public {
        require(keyRegistry[name].price > 0); // The name must be for sale.
        uint256 price = keyRegistry[name].price;
        payable(keyRegistry[name].owner).transfer(price);
        keyRegistry[name].owner = msg.sender;
        keyRegistry[name].publicKey = publicKey;
        keyRegistry[name].price = 0;
        emit Sale(name, price, msg.sender);
    }

    // Register a connector. Connectors control how to contact the owner of the name using various protocols.
    function registerConnector(
        string calldata name,
        string calldata protocol,
        string calldata location
    ) public {
        require(keyRegistry[name].owner == msg.sender); // Must be the owner.
        connectorRegistry[name][protocol] = location;
        emit ConnectorRegistered(name, protocol, location);
    }

    // Lookup connector's info.
    function lookupConnector(string calldata name, string calldata protocol)
        public
        view
        returns (string memory)
    {
        return (connectorRegistry[name][protocol]);
    }
}
