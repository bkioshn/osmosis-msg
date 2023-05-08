package message

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	lockuptypes "github.com/osmosis-labs/osmosis/v15/x/lockup/types"
)

// MsgLockTokens -> MsgBeginUnlocking
func CraftMsgLockTokens(acc client.Account) sdk.Msg {
	return &lockuptypes.MsgLockTokens{
		Duration: time.Duration(604800),
		Coins:    sdk.NewCoins(sdk.NewCoin("uosmo", sdk.NewInt(4))),
		Owner:    acc.GetAddress().String(),
	}
}

// MsgLockTokens -> MsgBeginUnlockingAll
func CraftMsgBeginUnlocking(acc client.Account) sdk.Msg {
	return &lockuptypes.MsgBeginUnlocking{
		Owner: acc.GetAddress().String(),
		ID:    1052450,
		Coins: sdk.NewCoins(sdk.NewCoin("uosmo", sdk.NewInt(1))),
	}
}

// MsgLockTokens -> MsgBeginUnlockingAll
func CraftMsgBeginUnlockingAll(acc client.Account) sdk.Msg {
	return &lockuptypes.MsgBeginUnlockingAll{
		Owner: acc.GetAddress().String(),
	}
}

// MsgBeginUnlockingAll -> MsgExtendLockup
func CraftMsgExtendLockup(acc client.Account) sdk.Msg {
	return &lockuptypes.MsgExtendLockup{
		Owner:    acc.GetAddress().String(),
		ID:       1052452,
		Duration: time.Duration(19000000),
	}
}
func CraftMsgForceUnlock(acc client.Account) sdk.Msg {
	return &lockuptypes.MsgForceUnlock{
		Owner: acc.GetAddress().String(),
		ID:    1052490,
		Coins: sdk.NewCoins(sdk.NewCoin("uosfmo", sdk.NewInt(1))),
	}
}
