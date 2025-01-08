package keeper

import (
	"context"

	"bitmail/x/ehl/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) HashCidAll(goCtx context.Context, req *types.QueryAllHashCidRequest) (*types.QueryAllHashCidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hashCids []types.HashCid
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	hashCidStore := prefix.NewStore(store, types.KeyPrefix(types.HashCidKey))

	pageRes, err := query.Paginate(hashCidStore, req.Pagination, func(key []byte, value []byte) error {
		var hashCid types.HashCid
		if err := k.cdc.Unmarshal(value, &hashCid); err != nil {
			return err
		}

		hashCids = append(hashCids, hashCid)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHashCidResponse{HashCid: hashCids, Pagination: pageRes}, nil
}

func (k Keeper) HashCid(goCtx context.Context, req *types.QueryGetHashCidRequest) (*types.QueryGetHashCidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	hashCid, found := k.GetHashCid(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetHashCidResponse{HashCid: hashCid}, nil
}
