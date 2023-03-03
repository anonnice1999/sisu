// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./base.sol";

contract Example is BaseContract {
    uint256 counter;
    event MessageReceived(uint256 id, bytes message);

    function onReceive(Message calldata input) external returns (uint8 code) {
        require(input.vault != address(0));
        counter++;

        require(input.message.length > 1, "invalid message");
        emit MessageReceived(counter, input.message);
        return 0;
    }
}
