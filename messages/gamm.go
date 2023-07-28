package message

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	stablepool "github.com/osmosis-labs/osmosis/v16/x/gamm/pool-models/stableswap"

	balancerpool "github.com/osmosis-labs/osmosis/v16/x/gamm/pool-models/balancer"
	gammtypes "github.com/osmosis-labs/osmosis/v16/x/gamm/types"
	poolmanagertypes "github.com/osmosis-labs/osmosis/v16/x/poolmanager/types"
)

func CraftMsgSwapExactAmountIn(acc client.Account) sdk.Msg {
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

func CraftMsgJoinPool(acc client.Account) sdk.Msg {
	return &gammtypes.MsgSwapExactAmountIn{
		Sender: acc.GetAddress().String(),
		Routes: []poolmanagertypes.SwapAmountInRoute{
			{
				PoolId:        2,
				TokenOutDenom: "uion",
			},
		},
		TokenIn:           sdk.NewCoin("uosmo", sdk.NewInt(2000000)),
		TokenOutMinAmount: sdk.NewInt(19),
	}
}
func CraftMsgExitPool(acc client.Account) sdk.Msg {
	return &gammtypes.MsgExitPool{
		Sender:        acc.GetAddress().String(),
		PoolId:        1,
		ShareInAmount: sdk.NewInt(357468035653720047),
		TokenOutMins: sdk.NewCoins(
			sdk.NewCoin("ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", sdk.NewInt(2000)),
			sdk.NewCoin("uosmo", sdk.NewInt(20000)),
		),
	}
}
func CraftMsgSwapExactAmountOut(acc client.Account) sdk.Msg {
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
func CraftMsgJoinSwapExternAmountIn(acc client.Account) sdk.Msg {
	return &gammtypes.MsgJoinSwapExternAmountIn{
		Sender:            acc.GetAddress().String(),
		PoolId:            788,
		TokenIn:           sdk.NewCoin("factory/osmo1acqpnvg2t4wmqfdv8hq47d3petfksjs5r9t45p/buhbubu", sdk.NewInt(2)),
		ShareOutMinAmount: sdk.NewInt(1),
	}
}
func CraftMsgJoinSwapShareAmountOut(acc client.Account) sdk.Msg {
	return &gammtypes.MsgJoinSwapShareAmountOut{
		Sender:           acc.GetAddress().String(),
		PoolId:           788,
		TokenInDenom:     "factory/osmo1acqpnvg2t4wmqfdv8hq47d3petfksjs5r9t45p/buhbubu",
		ShareOutAmount:   sdk.NewInt(100),
		TokenInMaxAmount: sdk.NewInt(200),
	}
}
func CraftMsgExitSwapShareAmountIn(acc client.Account) sdk.Msg {
	return &gammtypes.MsgExitSwapShareAmountIn{
		Sender:            acc.GetAddress().String(),
		PoolId:            788,
		TokenOutDenom:     "factory/osmo1acqpnvg2t4wmqfdv8hq47d3petfksjs5r9t45p/buhbubu",
		ShareInAmount:     sdk.NewInt(100000000000),
		TokenOutMinAmount: sdk.NewIntWithDecimal(10, 1),
	}
}
func CraftMsgExitSwapExternAmountOut(acc client.Account) sdk.Msg {
	return &gammtypes.MsgExitSwapExternAmountOut{
		Sender:           acc.GetAddress().String(),
		PoolId:           1,
		TokenOut:         sdk.NewCoin("uosmo", sdk.NewInt(2000000)),
		ShareInMaxAmount: sdk.NewInt(2),
	}
}
func CraftMsgCreateBalancerPool(acc client.Account) sdk.Msg {
	return &balancerpool.MsgCreateBalancerPool{
		Sender: acc.GetAddress().String(),
		PoolParams: &balancerpool.PoolParams{
			SwapFee: sdk.NewDec(0),
			ExitFee: sdk.NewDec(0),
		},
		PoolAssets: []balancerpool.PoolAsset{
			{
				Token:  sdk.NewCoin("uosmo", sdk.NewInt(1)),
				Weight: sdk.NewInt(1),
			},
			{
				Token:  sdk.NewCoin("factory/osmo1acqpnvg2t4wmqfdv8hq47d3petfksjs5r9t45p/buhbubu", sdk.NewInt(1)),
				Weight: sdk.NewInt(1),
			},
			// {
			// 	Token:  sdk.NewCoin("gamm/pool/782", sdk.NewInt(1)),
			// 	Weight: sdk.NewInt(1),
			// },
			// {
			// 	Token:  sdk.NewCoin("gamm/pool/1", sdk.NewInt(1)),
			// 	Weight: sdk.NewInt(1),
			// },
			// {
			// 	Token:  sdk.NewCoin("gamm/pool/784", sdk.NewInt(1)),
			// 	Weight: sdk.NewInt(1),
			// },
			// {
			// 	Token:  sdk.NewCoin("gamm/pool/788", sdk.NewInt(1)),
			// 	Weight: sdk.NewInt(1),
			// },
			// {
			// 	Token:  sdk.NewCoin("ibc/1DC495FCEFDA068A3820F903EDBD78B942FBD204D7E93D3BA2B432E9669D1A59", sdk.NewInt(1)),
			// 	Weight: sdk.NewInt(1),
			// },
			// {
			// 	Token:  sdk.NewCoin("ibc/307E5C96C8F60D1CBEE269A9A86C0834E1DB06F2B3788AE4F716EDB97A48B97D", sdk.NewInt(1)),
			// 	Weight: sdk.NewInt(1),
			// },
		},
		FuturePoolGovernor: "",
	}
}
func CrafMsgCreateStableswapPool(acc client.Account) sdk.Msg {
	return &stablepool.MsgCreateStableswapPool{
		Sender:     acc.GetAddress().String(),
		PoolParams: &stablepool.PoolParams{},
		InitialPoolLiquidity: sdk.NewCoins(sdk.NewCoin("uosmo", sdk.NewInt(1)),
			sdk.NewCoin("uion", sdk.NewInt(1)),
			sdk.NewCoin("ibc/307E5C96C8F60D1CBEE269A9A86C0834E1DB06F2B3788AE4F716EDB97A48B97D",
				sdk.NewInt(1)),
			sdk.NewCoin("ibc/1DC495FCEFDA068A3820F903EDBD78B942FBD204D7E93D3BA2B432E9669D1A59",
				sdk.NewInt(1)),
			sdk.NewCoin("gamm/pool/1", sdk.NewInt(1)),
			sdk.NewCoin("gamm/pool/784", sdk.NewInt(1)),
			sdk.NewCoin("gamm/pool/788", sdk.NewInt(1)),
			sdk.NewCoin("factory/osmo1acqpnvg2t4wmqfdv8hq47d3petfksjs5r9t45p/buhbubu", sdk.NewInt(1)),
		),
		ScalingFactors:          []uint64{1, 1, 1, 1, 1, 1, 1, 1},
		ScalingFactorController: "",
		FuturePoolGovernor:      acc.GetAddress().String(),
	}
}
func CraftMsgStableSwapAdjustScalingFactors(acc client.Account) sdk.Msg {
	return &stablepool.MsgStableSwapAdjustScalingFactors{
		Sender:         acc.GetAddress().String(),
		PoolID:         783,
		ScalingFactors: []uint64{1, 1},
	}
}
