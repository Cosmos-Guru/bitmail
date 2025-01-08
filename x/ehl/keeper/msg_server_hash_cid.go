package keeper

import (
	"context"
	"fmt"

	"bitmail/x/ehl/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateHashCid(goCtx context.Context, msg *types.MsgCreateHashCid) (*types.MsgCreateHashCidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var hashCid = types.HashCid{
		Creator:  msg.Creator,
		Receiver: msg.Receiver,
		Hashlink: msg.Hashlink,
		Vaultid:  msg.Vaultid,
	}

	id := k.AppendHashCid(
		ctx,
		hashCid,
	)

	return &types.MsgCreateHashCidResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateHashCid(goCtx context.Context, msg *types.MsgUpdateHashCid) (*types.MsgUpdateHashCidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var hashCid = types.HashCid{
		Creator:  msg.Creator,
		Id:       msg.Id,
		Receiver: msg.Receiver,
		Hashlink: msg.Hashlink,
		Vaultid:  msg.Vaultid,
	}

	// Checks that the element exists
	val, found := k.GetHashCid(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetHashCid(ctx, hashCid)

	return &types.MsgUpdateHashCidResponse{}, nil
}

func (k msgServer) DeleteHashCid(goCtx context.Context, msg *types.MsgDeleteHashCid) (*types.MsgDeleteHashCidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetHashCid(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHashCid(ctx, msg.Id)

	return &types.MsgDeleteHashCidResponse{}, nil
}
