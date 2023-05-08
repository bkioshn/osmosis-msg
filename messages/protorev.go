package message

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	protorevtypes "github.com/osmosis-labs/osmosis/v15/x/protorev/types"
)

func CraftMsgSetHotRoutes(acc client.Account) sdk.Msg {
	return &protorevtypes.MsgSetHotRoutes{
		Admin: acc.GetAddress().String(),
		HotRoutes: []protorevtypes.TokenPairArbRoutes{{
			ArbRoutes: []protorevtypes.Route{{
				Trades: []protorevtypes.Trade{{
					Pool:     0,
					TokenIn:  "uosmo",
					TokenOut: "uion",
				}, {
					Pool:     0,
					TokenIn:  "uion",
					TokenOut: "uosmo",
				},
				},
				StepSize: sdk.NewInt(2),
			}},
			TokenIn:  "uosmo",
			TokenOut: "uion",
		},
			{
				ArbRoutes: []protorevtypes.Route{{

					Trades: []protorevtypes.Trade{{
						Pool:     0,
						TokenIn:  "uosmo",
						TokenOut: "gamm/pool/1",
					}, {
						Pool:     0,
						TokenIn:  "gamm/pool/1",
						TokenOut: "uosmo",
					},
					},
					StepSize: sdk.NewInt(2),
				}},
				TokenIn:  "uosmo",
				TokenOut: "gamm/pool/1",
			},
		},
	}
}
func CraftMsgSetDeveloperAccount(acc client.Account) sdk.Msg {
	return &protorevtypes.MsgSetDeveloperAccount{
		Admin:            acc.GetAddress().String(),
		DeveloperAccount: acc.GetAddress().String(),
	}
}
func CraftMsgSetPoolWeights(acc client.Account) sdk.Msg {
	return &protorevtypes.MsgSetPoolWeights{
		Admin: acc.GetAddress().String(),
		PoolWeights: protorevtypes.PoolWeights{
			StableWeight:       1,
			BalancerWeight:     1,
			ConcentratedWeight: 1,
		},
	}
}
func CraftMsgSetMaxPoolPointsPerTx(acc client.Account) sdk.Msg {
	return &protorevtypes.MsgSetMaxPoolPointsPerTx{
		Admin:              acc.GetAddress().String(),
		MaxPoolPointsPerTx: 1,
	}
}
func CraftMsgSetMaxPoolPointsPerBlock(acc client.Account) sdk.Msg {
	return &protorevtypes.MsgSetMaxPoolPointsPerBlock{
		Admin:                 acc.GetAddress().String(),
		MaxPoolPointsPerBlock: 1,
	}
}
func CraftMsgSetBaseDenoms(acc client.Account) sdk.Msg {
	return &protorevtypes.MsgSetBaseDenoms{
		Admin: acc.GetAddress().String(),
		BaseDenoms: []protorevtypes.BaseDenom{{
			Denom:    "uosmo",
			StepSize: sdk.NewInt(2),
		},
		},
	}
}
