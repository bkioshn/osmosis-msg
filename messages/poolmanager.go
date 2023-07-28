package message

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	poolmanagertypes "github.com/osmosis-labs/osmosis/v16/x/poolmanager/types"
)

func CraftMsgSwapExactAmountInManager(acc client.Account) sdk.Msg {
	return &poolmanagertypes.MsgSwapExactAmountIn{
		Sender: acc.GetAddress().String(),
		Routes: []poolmanagertypes.SwapAmountInRoute{
			{
				PoolId:        1,
				TokenOutDenom: "uion",
			},
		},
		TokenIn:           sdk.NewCoin("uosmo", sdk.NewInt(2000000)),
		TokenOutMinAmount: sdk.NewInt(19),
	}
}
func CraftMsgSplitRouteSwapExactAmountIn(acc client.Account) sdk.Msg {
	return &poolmanagertypes.MsgSplitRouteSwapExactAmountIn{
		Sender:            acc.GetAddress().String(),
		Routes:            []poolmanagertypes.SwapAmountInSplitRoute{},
		TokenInDenom:      "osmo",
		TokenOutMinAmount: sdk.NewInt(19),
	}

}
func CraftMsgSwapExactAmountOutManager(acc client.Account) sdk.Msg {
	return &poolmanagertypes.MsgSwapExactAmountOut{
		Sender: acc.GetAddress().String(),
		Routes: []poolmanagertypes.SwapAmountOutRoute{
			{
				PoolId:       1,
				TokenInDenom: "uosmo",
			},
		},
		TokenInMaxAmount: sdk.NewInt(100000),
		TokenOut:         sdk.NewCoin("uion", sdk.NewInt(1)),
	}

}
func CraftMsgSplitRouteSwapExactAmountOut(acc client.Account) sdk.Msg {
	return &poolmanagertypes.MsgSplitRouteSwapExactAmountOut{
		Sender:           acc.GetAddress().String(),
		Routes:           []poolmanagertypes.SwapAmountOutSplitRoute{},
		TokenOutDenom:    "osmo",
		TokenInMaxAmount: sdk.NewInt(19),
	}

}
