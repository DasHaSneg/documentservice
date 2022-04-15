package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSignAnnex = "sign_annex"

var _ sdk.Msg = &MsgSignAnnex{}

func NewMsgSignAnnex(creator string, annexId uint64) *MsgSignAnnex {
	return &MsgSignAnnex{
		Creator: creator,
		AnnexId: annexId,
	}
}

func (msg *MsgSignAnnex) Route() string {
	return RouterKey
}

func (msg *MsgSignAnnex) Type() string {
	return TypeMsgSignAnnex
}

func (msg *MsgSignAnnex) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSignAnnex) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSignAnnex) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
