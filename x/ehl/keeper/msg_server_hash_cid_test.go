package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"bitmail/x/ehl/types"
)

func TestHashCidMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateHashCid(ctx, &types.MsgCreateHashCid{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestHashCidMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateHashCid
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateHashCid{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateHashCid{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateHashCid{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateHashCid(ctx, &types.MsgCreateHashCid{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateHashCid(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestHashCidMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteHashCid
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteHashCid{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteHashCid{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteHashCid{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateHashCid(ctx, &types.MsgCreateHashCid{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteHashCid(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
