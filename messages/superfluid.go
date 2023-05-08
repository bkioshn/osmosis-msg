package message

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	superfluidtypes "github.com/osmosis-labs/osmosis/v15/x/superfluid/types"
)

// 1052504
// MsgLockTokens -> MsgSuperfluidDelegate
func CraftMsgSuperfluidDelegate(acc client.Account) sdk.Msg {
	return &superfluidtypes.MsgSuperfluidDelegate{
		Sender:  acc.GetAddress().String(),
		LockId:  1052456,
		ValAddr: "osmovaloper1c584m4lq25h83yp6ag8hh4htjr92d954kphp96",
	}
}

// MsgLockTokens -> MsgSuperfluidDelegate -> MsgSuperfluidUndelegate

func CraftMsgSuperfluidUndelegate(acc client.Account) sdk.Msg {
	return &superfluidtypes.MsgSuperfluidUndelegate{
		Sender: acc.GetAddress().String(),
		LockId: 1052456,
	}
}
func CraftMsgLockAndSuperfluidDelegate(acc client.Account) sdk.Msg {
	return &superfluidtypes.MsgLockAndSuperfluidDelegate{
		Sender:  acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin("gamm/pool/1", sdk.NewInt(1))),
		ValAddr: "osmovaloper1c584m4lq25h83yp6ag8hh4htjr92d954kphp96",
	}
}

// MsgLockAndSuperfluidDelegate -> MsgSuperfluidUnbondLock

func CraftMsgSuperfluidUnbondLock(acc client.Account) sdk.Msg {
	return &superfluidtypes.MsgSuperfluidUnbondLock{
		Sender: acc.GetAddress().String(),
		LockId: 1052457,
	}
}
func CraftMsgUnPoolWhitelistedPool(acc client.Account) sdk.Msg {
	return &superfluidtypes.MsgUnPoolWhitelistedPool{
		Sender: acc.GetAddress().String(),
		PoolId: 1,
	}
}

// MsgLockAndSuperfluidDelegate -> MsgSuperfluidUndelegateAndUnbondLock

func CraftMsgSuperfluidUndelegateAndUnbondLock(acc client.Account) sdk.Msg {
	return &superfluidtypes.MsgSuperfluidUndelegateAndUnbondLock{
		Sender: acc.GetAddress().String(),
		LockId: 1052457,
		Coin:   sdk.NewCoin("gamm/pool/1", sdk.NewInt(1)),
	}
}
