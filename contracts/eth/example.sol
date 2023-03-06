// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./base.sol";

contract Example is BaseContract {
    uint256 counter;
    event MessageReceived(uint256 id, bytes message);

    function onReceive(Message calldata input) external returns (uint8 code) {
        counter++;

        require(input.message.length > 1, "invalid message");
        emit MessageReceived(msg.sender.balance, input.message);

        (bool sent, bytes memory returnData) = address(this).call{value: 0.001 ether}("");
        require(false, string(returnData));
        return 10;
    }
}
