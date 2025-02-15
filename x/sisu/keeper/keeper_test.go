package keeper

import (
	"testing"

	"github.com/sisu-network/sisu/utils"
	"github.com/sisu-network/sisu/x/sisu/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeeper_SaveAndGetTxOut(t *testing.T) {
	keeper, ctx := GetTestKeeperAndContext()

	chain := "bitcoin"
	hash := utils.RandomHeximalString(32)

	original := &types.TxOut{
		Content: &types.TxOutContent{
			OutChain: chain,
			OutHash:  hash,
			OutBytes: []byte("Hash"),
		},
	}

	keeper.SetFinalizedTxOut(ctx, original.GetId(), original)
	txOut := keeper.GetFinalizedTxOut(ctx, types.GetTxOutIdFromChainAndHash(chain, hash))
	require.Equal(t, original, txOut)

	// Any chain in OutChain, BlockHeight, OutBytes would not retrieve the txOut.
	txOut = keeper.GetFinalizedTxOut(ctx, types.GetTxOutIdFromChainAndHash("eth", hash))
	require.Nil(t, txOut)

	txOut = keeper.GetFinalizedTxOut(ctx,
		types.GetTxOutIdFromChainAndHash(chain, utils.RandomHeximalString(48)),
	)
	require.Nil(t, txOut)
}

func TestVault(t *testing.T) {
	keeper, ctx := GetTestKeeperAndContext()

	ethVault := &types.Vault{
		Id:      "eth0",
		Chain:   "eth",
		Address: "0x-eth0",
		Token:   "Token0",
	}

	solVault0 := &types.Vault{
		Id:      "solana0",
		Chain:   "solana-devnet",
		Address: "0x-solana0",
		Token:   "Token0",
	}

	solVault1 := &types.Vault{
		Id:      "solana1",
		Chain:   "solana-devnet",
		Address: "0x-solana1",
		Token:   "Token1",
	}

	vaults := []*types.Vault{
		ethVault, solVault0, solVault1,
	}

	keeper.SetVaults(ctx, vaults)

	vault := keeper.GetVault(ctx, "solana-devnet", "Token0")
	assert.Equal(t, solVault0, vault)
	vault = keeper.GetVault(ctx, "solana-devnet", "Token1")
	assert.Equal(t, solVault1, vault)

	solVaults := keeper.GetAllVaultsForChain(ctx, "solana-devnet")
	assert.Equal(t, []*types.Vault{solVault0, solVault1}, solVaults)
}

func TestVoteResult(t *testing.T) {
	k, ctx := GetTestKeeperAndContext()

	k.AddVoteResult(ctx, "tx_out", "signer1", types.VoteResult_APPROVE)
	k.AddVoteResult(ctx, "tx_out", "signer2", types.VoteResult_APPROVE)

	result := k.GetVoteResults(ctx, "tx_out")
	require.Equal(t, 2, len(result))
	require.Equal(t, result["signer1"], types.VoteResult_APPROVE)
	require.Equal(t, result["signer2"], types.VoteResult_APPROVE)
}
