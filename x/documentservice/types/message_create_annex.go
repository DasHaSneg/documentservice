package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateAnnex = "create_annex"

var _ sdk.Msg = &MsgCreateAnnex{}

func NewMsgCreateAnnex(creator string, annexHash string, contractId uint64, buyer string) *MsgCreateAnnex {
	return &MsgCreateAnnex{
		Creator:    creator,
		AnnexHash:  annexHash,
		ContractId: contractId,
		Buyer:      buyer,
	}
}

func (msg *MsgCreateAnnex) Route() string {
	return RouterKey
}

func (msg *MsgCreateAnnex) Type() string {
	return TypeMsgCreateAnnex
}

func (msg *MsgCreateAnnex) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAnnex) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAnnex) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
