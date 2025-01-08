package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "bitmail/testutil/keeper"
	"bitmail/testutil/nullify"
	"bitmail/x/ehl/types"
)

func TestHashCidQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EhlKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNHashCid(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetHashCidRequest
		response *types.QueryGetHashCidResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetHashCidRequest{Id: msgs[0].Id},
			response: &types.QueryGetHashCidResponse{HashCid: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetHashCidRequest{Id: msgs[1].Id},
			response: &types.QueryGetHashCidResponse{HashCid: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetHashCidRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.HashCid(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestHashCidQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EhlKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNHashCid(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllHashCidRequest {
		return &types.QueryAllHashCidRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.HashCidAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.HashCid), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.HashCid),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.HashCidAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.HashCid), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.HashCid),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.HashCidAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.HashCid),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.HashCidAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
