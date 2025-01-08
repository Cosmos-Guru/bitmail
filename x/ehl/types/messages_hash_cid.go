package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateHashCid = "create_hash_cid"
	TypeMsgUpdateHashCid = "update_hash_cid"
	TypeMsgDeleteHashCid = "delete_hash_cid"
)

var _ sdk.Msg = &MsgCreateHashCid{}

func NewMsgCreateHashCid(creator string, receiver string, hashlink string, vaultid string) *MsgCreateHashCid {
	return &MsgCreateHashCid{
		Creator:  creator,
		Receiver: receiver,
		Hashlink: hashlink,
		Vaultid:  vaultid,
	}
}

func (msg *MsgCreateHashCid) Route() string {
	return RouterKey
}

func (msg *MsgCreateHashCid) Type() string {
	return TypeMsgCreateHashCid
}

func (msg *MsgCreateHashCid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateHashCid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateHashCid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateHashCid{}

func NewMsgUpdateHashCid(creator string, id uint64, receiver string, hashlink string, vaultid string) *MsgUpdateHashCid {
	return &MsgUpdateHashCid{
		Id:       id,
		Creator:  creator,
		Receiver: receiver,
		Hashlink: hashlink,
		Vaultid:  vaultid,
	}
}

func (msg *MsgUpdateHashCid) Route() string {
	return RouterKey
}

func (msg *MsgUpdateHashCid) Type() string {
	return TypeMsgUpdateHashCid
}

func (msg *MsgUpdateHashCid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateHashCid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateHashCid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteHashCid{}

func NewMsgDeleteHashCid(creator string, id uint64) *MsgDeleteHashCid {
	return &MsgDeleteHashCid{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteHashCid) Route() string {
	return RouterKey
}

func (msg *MsgDeleteHashCid) Type() string {
	return TypeMsgDeleteHashCid
}

func (msg *MsgDeleteHashCid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteHashCid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteHashCid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
