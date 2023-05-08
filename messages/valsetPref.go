package message

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	valsetpreftypes "github.com/osmosis-labs/osmosis/v15/x/valset-pref/types"
)

func CraftMsgSetValidatorSetPreference(acc client.Account) sdk.Msg {
	return &valsetpreftypes.MsgSetValidatorSetPreference{
		Delegator: acc.GetAddress().String(),
		Preferences: []valsetpreftypes.ValidatorPreference{
			{
				ValOperAddress: "osmovaloper1c584m4lq25h83yp6ag8hh4htjr92d954kphp96",
				Weight:         sdk.NewDec(1),
			},
		},
	}
}
func CraftMsgDelegateToValidatorSet(acc client.Account) sdk.Msg {
	return &valsetpreftypes.MsgDelegateToValidatorSet{
		Delegator: acc.GetAddress().String(),
		Coin:      sdk.NewCoin("uosmo", sdk.NewInt(20)),
	}
}
func CraftMsgRedelegateValidatorSet(acc client.Account) sdk.Msg {
	return &valsetpreftypes.MsgRedelegateValidatorSet{
		Delegator: acc.GetAddress().String(),
		Preferences: []valsetpreftypes.ValidatorPreference{
			{
				ValOperAddress: "osmovaloper1acqpnvg2t4wmqfdv8hq47d3petfksjs5ejrkrx",
				Weight:         sdk.NewDec(1),
			},
		},
	}
}
func CraftMsgUndelegateFromValidatorSet(acc client.Account) sdk.Msg {
	return &valsetpreftypes.MsgUndelegateFromValidatorSet{
		Delegator: acc.GetAddress().String(),
		Coin:      sdk.NewCoin("uosmo", sdk.NewInt(1)),
	}
}
func CraftMsgWithdrawDelegationRewards(acc client.Account) sdk.Msg {
	return &valsetpreftypes.MsgWithdrawDelegationRewards{
		Delegator: acc.GetAddress().String(),
	}
}
func CraftMsgDelegateBondedTokens(acc client.Account) sdk.Msg {
	return &valsetpreftypes.MsgDelegateBondedTokens{
		Delegator: acc.GetAddress().String(),
		LockID:    1052461,
	}
}
