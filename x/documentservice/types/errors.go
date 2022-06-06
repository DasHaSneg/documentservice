package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/documentservice module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrID     = sdkerrors.Register(ModuleName, 1200, "")
	ErrBuyer  = sdkerrors.Register(ModuleName, 1300, "")
	ErrState  = sdkerrors.Register(ModuleName, 1400, "")
	ErrAccess = sdkerrors.Register(ModuleName, 1500, "")
)
