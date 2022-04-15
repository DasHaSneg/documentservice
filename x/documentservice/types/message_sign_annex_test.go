package types

import (
	"testing"

	"github.com/cosmonaut/documentservice/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgSignAnnex_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSignAnnex
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSignAnnex{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSignAnnex{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
