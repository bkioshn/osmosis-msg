package message

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	clpoolmodel "github.com/osmosis-labs/osmosis/v16/x/concentrated-liquidity/model"
	cltypes "github.com/osmosis-labs/osmosis/v16/x/concentrated-liquidity/types"
)

func CraftMsgCreatePosition(acc client.Account) sdk.Msg {
	return &cltypes.MsgCreatePosition{
		Sender:          acc.GetAddress().String(),
		PoolId:          1,
		LowerTick:       1,
		UpperTick:       2,
		TokensProvided:  sdk.NewCoins(sdk.NewCoin("uosmo", sdk.NewInt(1))),
		TokenMinAmount0: sdk.NewInt(1),
		TokenMinAmount1: sdk.NewInt(1),
	}
}
func CraftMsgAddToPosition(acc client.Account) sdk.Msg {
	return &cltypes.MsgAddToPosition{
		Sender:          acc.GetAddress().String(),
		PositionId:      1,
		Amount0:         sdk.NewInt(1),
		Amount1:         sdk.NewInt(1),
		TokenMinAmount0: sdk.NewInt(1),
		TokenMinAmount1: sdk.NewInt(1),
	}
}
func CraftMsgWithdrawPosition(acc client.Account) sdk.Msg {
	return &cltypes.MsgWithdrawPosition{
		Sender:          acc.GetAddress().String(),
		PositionId:      1,
		LiquidityAmount: sdk.NewDec(1),
	}
}
func CraftMsgCollectSpreadRewards(acc client.Account) sdk.Msg {
	return &cltypes.MsgCollectSpreadRewards{
		Sender:      acc.GetAddress().String(),
		PositionIds: []uint64{1, 2, 3},
	}
}
func CraftMsgCollectIncentives(acc client.Account) sdk.Msg {
	return &cltypes.MsgCollectIncentives{
		Sender:      acc.GetAddress().String(),
		PositionIds: []uint64{1, 2, 3},
	}
}
func CraftMsgCreateConcentratedPool(acc client.Account) sdk.Msg {
	return &clpoolmodel.MsgCreateConcentratedPool{
		Sender:       acc.GetAddress().String(),
		Denom0:       "buhbubu",
		Denom1:       "uosmo",
		TickSpacing:  100,
		SpreadFactor: sdk.NewDec(1),
	}
}
