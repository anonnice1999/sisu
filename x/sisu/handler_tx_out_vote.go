package sisu

import (
	"fmt"

	"github.com/sisu-network/sisu/x/sisu/components"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sisu-network/lib/log"
	"github.com/sisu-network/sisu/x/sisu/keeper"
	"github.com/sisu-network/sisu/x/sisu/types"
)

const (
	VoteKey = "TxOutVote"
)

type HandlerTxOutVote struct {
	pmm       components.PostedMessageManager
	keeper    keeper.Keeper
	privateDb keeper.PrivateDb
}

func NewHandlerTxOutConsensed(
	pmm components.PostedMessageManager,
	keeper keeper.Keeper,
	privateDb keeper.PrivateDb,
) *HandlerTxOutVote {
	return &HandlerTxOutVote{
		pmm:       pmm,
		keeper:    keeper,
		privateDb: privateDb,
	}
}

func (h *HandlerTxOutVote) DeliverMsg(
	ctx sdk.Context,
	msg *types.TxOutVoteMsg,
) (*sdk.Result, error) {
	txOut := h.keeper.GetProposedTxOut(ctx, msg.Data.TxOutId, msg.Data.AssignedValidator)
	if txOut == nil {
		log.Error("Cannot get proposed txout, txOutId = ", msg.Data.TxOutId)
		return &sdk.Result{}, nil
	}

	done := h.keeper.IsTxRecordProcessed(ctx,
		[]byte(fmt.Sprintf("%s__%s", txOut.GetId(), msg.Data.AssignedValidator)))
	if done {
		log.Info("Ignore the completed proposed txout, txOutId = ", msg.Data.TxOutId)
		return &sdk.Result{}, nil
	}

	counter := h.keeper.GetTransferCounter(ctx, txOut.Input.TransferIds[0])
	prefix := fmt.Sprintf("%s__%s__%d", VoteKey, msg.Data.TxOutId, counter)
	h.keeper.AddVoteResult(ctx, prefix, msg.Signer, msg.Data.Vote)

	h.checkVoteResult(ctx, txOut, counter, msg.Data.AssignedValidator)

	return &sdk.Result{}, nil
}

func (h *HandlerTxOutVote) checkVoteResult(
	ctx sdk.Context,
	txOut *types.TxOut,
	counter int,
	assignedValidator string,
) {
	txOutId := txOut.GetId()

	prefix := fmt.Sprintf("%s__%s__%d", VoteKey, txOutId, counter)
	results := h.keeper.GetVoteResults(ctx, prefix)
	params := h.keeper.GetParams(ctx)
	if params == nil {
		log.Warn("tssParams is nil")
		return
	}

	approveCount := 0
	rejectCount := 0
	for _, result := range results {
		if result == types.VoteResult_APPROVE {
			approveCount++
		} else {
			rejectCount++
		}
	}

	threshold := int(params.MajorityThreshold)

	if approveCount < threshold && rejectCount < threshold {
		// TODO: handler the case or do timeout in the module.go
		return
	}

	h.keeper.ProcessTxRecord(ctx, []byte(fmt.Sprintf("%s__%s", txOut.GetId(), assignedValidator)))

	if approveCount >= threshold {
		finalizedTxOut := h.keeper.GetFinalizedTxOut(ctx, txOutId)
		if finalizedTxOut == nil {
			h.doTxOut(ctx, h.keeper, h.privateDb, txOut)
		} else {
			log.Verbosef("Finalized TxOut has been processed for txOut with id %s", txOutId)
		}
	} else {
		log.Verbose("TxOut is rejected, txOutId = ", txOutId)
		h.keeper.IncTransferCounter(ctx, txOut.Input.TransferIds[0])
		h.privateDb.SetHoldProcessing(types.TransferHoldKey, txOut.Content.OutChain, false)
	}
}

// doTxOut saves a TxOut in the keeper and add it the TxOut Queue.
func (h *HandlerTxOutVote) doTxOut(ctx sdk.Context, k keeper.Keeper, privateDb keeper.PrivateDb,
	txOut *types.TxOut) ([]byte, error) {
	log.Info("Finalizing TxOut, id = ", txOut.GetId())

	// Save this to KVStore
	k.SetFinalizedTxOut(ctx, txOut.GetId(), txOut)

	// If this is a txOut deployment, mark the contract as being deployed.
	switch txOut.TxType {
	case types.TxOutType_TRANSFER_OUT:
		h.handlerTransfer(ctx, k, privateDb, txOut)
	}

	return nil, nil
}

func (h *HandlerTxOutVote) handlerTransfer(ctx sdk.Context, k keeper.Keeper, privateDb keeper.PrivateDb,
	txOut *types.TxOut) {
	// 1. Update TxOut txOutQ.
	txOutQ := k.GetTxOutQueue(ctx, txOut.Content.OutChain)
	txOutQ = append(txOutQ, txOut)
	k.SetTxOutQueue(ctx, txOut.Content.OutChain, txOutQ)

	// 2. Remove the transfers in txOut from the queue.
	transferQ := k.GetTransferQueue(ctx, txOut.Content.OutChain)
	ids := make(map[string]bool, 0)
	for _, inHash := range txOut.Input.TransferIds {
		ids[inHash] = true
	}

	newQueue := make([]*types.TransferDetails, 0)
	for _, transfer := range transferQ {
		if !ids[transfer.Id] {
			newQueue = append(newQueue, transfer)
		}
	}

	k.SetTransferQueue(ctx, txOut.Content.OutChain, newQueue)

	// 3. Remove failed transfers (if any).
	for _, id := range txOut.Input.TransferIds {
		k.RemoveFailedTransfer(ctx, id)
	}

	// 4. Update the HoldProcessing for transfer queue so that we do not process any more transfer.
	privateDb.SetHoldProcessing(types.TransferHoldKey, txOut.Content.OutChain, true)

	// 5. Set Expiration Block.
	params := k.GetParams(ctx)
	k.SetExpirationBlock(ctx, types.ExpirationBlock_TxOut, txOut.GetId(),
		ctx.BlockHeight()+int64(params.ExpirationBlock))
}
