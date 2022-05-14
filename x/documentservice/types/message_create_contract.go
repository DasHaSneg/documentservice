package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateContract = "create_contract"

var _ sdk.Msg = &MsgCreateContract{}

func NewMsgCreateContract(creator string, contractHash string, buyer string, sellerInn string, buyerInn string) *MsgCreateContract {
	return &MsgCreateContract{
		Creator:      creator,
		ContractHash: contractHash,
		Buyer:        buyer,
		SellerInn:    sellerInn,
		BuyerInn:     buyerInn,
	}
}

func (msg *MsgCreateContract) Route() string {
	return RouterKey
}

func (msg *MsgCreateContract) Type() string {
	return TypeMsgCreateContract
}

func (msg *MsgCreateContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
