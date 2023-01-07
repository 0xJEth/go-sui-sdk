package client

import (
	"context"

	"github.com/coming-chat/go-sui/types"
)

// MintNFT
// Create an unsigned transaction to mint a nft at devnet
func (c *Client) MintNFT(ctx context.Context, signer types.Address, nftName, nftDescription, nftUri string, gas *types.ObjectId, gasBudget uint64) (*types.TransactionBytes, error) {
	packageId, _ := types.NewAddressFromHex("0x2")
	args := []any{
		nftName, nftDescription, nftUri,
	}
	return c.MoveCall(ctx, signer, *packageId, "devnet_nft", "mint", []string{}, args, gas, gasBudget)
}

func (c *Client) GetNFTsOwnedByAddress(ctx context.Context, address types.Address) ([]types.ObjectRead, error) {
	return c.BatchGetObjectsOwnedByAddress(ctx, address, "0x2::devnet_nft::DevNetNFT")
}
