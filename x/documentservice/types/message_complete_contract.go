package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCompleteContract = "complete_contract"

var _ sdk.Msg = &MsgCompleteContract{}

func NewMsgCompleteContract(creator string, contractId uint64) *MsgCompleteContract {
	return &MsgCompleteContract{
		Creator:    creator,
		ContractId: contractId,
	}
}

func (msg *MsgCompleteContract) Route() string {
	return RouterKey
}

func (msg *MsgCompleteContract) Type() string {
	return TypeMsgCompleteContract
}

func (msg *MsgCompleteContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCompleteContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCompleteContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
