package message

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	incentivetypes "github.com/osmosis-labs/osmosis/v15/x/incentives/types"

	lockuptypes "github.com/osmosis-labs/osmosis/v15/x/lockup/types"
)

func CraftMsgCreateGauge(acc client.Account) sdk.Msg {
	return &incentivetypes.MsgCreateGauge{
		IsPerpetual: false,
		Owner:       acc.GetAddress().String(),
		DistributeTo: lockuptypes.QueryCondition{
			LockQueryType: 0,
			Denom:         "factory/osmo1acqpnvg2t4wmqfdv8hq47d3petfksjs5r9t45p/buhbubu",
			Duration:      time.Duration(8.64e+13),
			Timestamp:     time.Now(),
		},
		Coins:             sdk.NewCoins(sdk.NewCoin("factory/osmo1acqpnvg2t4wmqfdv8hq47d3petfksjs5r9t45p/buhbubu", sdk.NewInt(1))),
		StartTime:         time.Now(),
		NumEpochsPaidOver: 1,
	}
}
func CraftMsgAddToGauge(acc client.Account) sdk.Msg {
	return &incentivetypes.MsgAddToGauge{
		Owner:   acc.GetAddress().String(),
		GaugeId: 2608,
		Rewards: sdk.NewCoins(sdk.NewCoin("factory/osmo1acqpnvg2t4wmqfdv8hq47d3petfksjs5r9t45p/buhbubu", sdk.NewInt(1))),
	}
}
