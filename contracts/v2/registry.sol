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

    event NameOwnershipChanged(string name, address newOwner);

    // Name's price was changed - maybe it was listed for sale, or was delisted.
    // Price of zero means the name is not for sale.
    event PriceChanged(string name, uint256 price);

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

        // Verify the name.
        bytes memory _baseBytes = bytes(name);
        for (uint256 i = 0; i < _baseBytes.length; i++) {
            bytes1 b = _baseBytes[i];
            // We only allow ASCII letters and numbers in the name, plus characters '-' and '_'.
            if (
                (b < 0x30 && b != 0x2D) || // Symbols, except '-'.
                b > 0x7A || // Symbols.
                (b >= 0x3A && b <= 0x40) || // More symbols.
                (b >= 0x5B && b <= 0x5E) || // Symbols, except '_'.
                b == 0x60 // Backtick.
            ) {
                revert("invalid name");
            }

            _baseBytes[i] = _lower(b);
        }
        string memory lowercaseName = string(_baseBytes);

        keyRegistry[lowercaseName].owner = msg.sender;
        keyRegistry[lowercaseName].publicKey = publicKey;
        emit NameRegistered(lowercaseName, msg.sender);
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
        string memory lowercaseName = lower(name);
        RegistryEntry storage e = keyRegistry[lowercaseName];
        return (e.owner, e.publicKey, e.price);
    }

    // Update public key for already registered name.
    function updatePublicKey(bytes calldata publicKey, string calldata name)
        public
    {
        string memory lowercaseName = lower(name);
        require(publicKey.length == 33);
        require(keyRegistry[lowercaseName].owner == msg.sender); // The name must not be taken.
        keyRegistry[lowercaseName].publicKey = publicKey;
        emit PublicKeyUpdated(lowercaseName);
    }

    // Transfer name ownership to a different owner.
    function transferOwnership(string calldata name, address newOwner) public {
        string memory lowercaseName = lower(name);
        require(keyRegistry[lowercaseName].owner == msg.sender); // Must be the owner.
        keyRegistry[lowercaseName].owner = newOwner;
        emit NameOwnershipChanged(lowercaseName, newOwner);
    }

    // Change name's price. If a price is non-zero, anyone can pay and assume ownership of a name.
    // A price of zero means the name is not for sale.
    function changePrice(string calldata name, uint256 price) public {
        string memory lowercaseName = lower(name);
        require(keyRegistry[lowercaseName].owner == msg.sender); // Must be the owner.
        keyRegistry[lowercaseName].price = price;
        emit PriceChanged(lowercaseName, price);
    }

    // Buy a name. The name must be listed for sale (price is greater than zero).
    function buy(string calldata name, bytes memory publicKey) public {
        string memory lowercaseName = lower(name);
        require(keyRegistry[lowercaseName].price > 0); // The name must be for sale.
        uint256 price = keyRegistry[lowercaseName].price;
        payable(keyRegistry[lowercaseName].owner).transfer(price);
        keyRegistry[lowercaseName].owner = msg.sender;
        keyRegistry[lowercaseName].publicKey = publicKey;
        keyRegistry[lowercaseName].price = 0;
        emit Sale(lowercaseName, price, msg.sender);
    }

    // Register a connector. Connectors control how to contact the owner of the name using various protocols.
    function registerConnector(
        string calldata name,
        string calldata protocol,
        string calldata location
    ) public {
        string memory lowercaseName = lower(name);
        require(keyRegistry[lowercaseName].owner == msg.sender); // Must be the owner.
        connectorRegistry[lowercaseName][protocol] = location;
        emit ConnectorRegistered(lowercaseName, protocol, location);
    }

    // Lookup connector's info.
    function lookupConnector(string calldata name, string calldata protocol)
        public
        view
        returns (string memory)
    {
        string memory lowercaseName = lower(name);
        return (connectorRegistry[lowercaseName][protocol]);
    }

    function lower(string memory _base) internal pure returns (string memory) {
        bytes memory _baseBytes = bytes(_base);
        for (uint256 i = 0; i < _baseBytes.length; i++) {
            _baseBytes[i] = _lower(_baseBytes[i]);
        }
        return string(_baseBytes);
    }

    function _lower(bytes1 _b1) private pure returns (bytes1) {
        if (_b1 >= 0x41 && _b1 <= 0x5A) {
            return bytes1(uint8(_b1) + 32);
        }

        return _b1;
    }
}
