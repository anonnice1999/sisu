// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface BaseContract {
    /**
     * @dev Message is the input of onReceive function.
     * @param vault: address of the vault which sent this transaction.
     * @param callerChain: the chain id of caller.
     * @param caller: the caller address.
     * @param message: the message is sent from sender.
     */
    struct Message {
        address vault;
        uint256 callerChain;
        address caller;
        bytes message;
    }

    /**
     * @dev onReceive will be triggered if a message is sent to this contract.
     */
    function onReceive(Message calldata input) external returns (uint8 code);
}
