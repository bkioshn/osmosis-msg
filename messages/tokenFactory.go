package message

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	tokenfactorytypes "github.com/osmosis-labs/osmosis/v15/x/tokenfactory/types"
)

func CraftMsgCreateDenom(acc client.Account) sdk.Msg {
	return &tokenfactorytypes.MsgCreateDenom{
		Sender:   acc.GetAddress().String(),
		Subdenom: "buhbubu",
	}
}
func CraftMsgMint(acc client.Account) sdk.Msg {
	return &tokenfactorytypes.MsgMint{
		Sender: acc.GetAddress().String(),
		Amount: sdk.NewCoin("factory/osmo1acqpnvg2t4wmqfdv8hq47d3petfksjs5r9t45p/buhbubu", sdk.NewInt(100000000000000000)),
	}
}
func CraftMsgBurn(acc client.Account) sdk.Msg {
	return &tokenfactorytypes.MsgBurn{
		Sender: acc.GetAddress().String(),
		Amount: sdk.NewCoin("factory/osmo1acqpnvg2t4wmqfdv8hq47d3petfksjs5r9t45p/buhbubu", sdk.NewInt(1)),
	}
}
func CraftMsgChangeAdmin(acc client.Account) sdk.Msg {
	return &tokenfactorytypes.MsgChangeAdmin{
		Sender:   acc.GetAddress().String(),
		NewAdmin: acc.GetAddress().String(),
		Denom:    "factory/osmo1acqpnvg2t4wmqfdv8hq47d3petfksjs5r9t45p/buhbubu",
	}
}
func CraftMsgSetDenomMetadata(acc client.Account) sdk.Msg {
	return &tokenfactorytypes.MsgSetDenomMetadata{
		Sender: acc.GetAddress().String(),
		Metadata: banktypes.Metadata{
			Description: "This token is for testing",
			Name:        "Test buhbubu",
			Symbol:      "BUBU",
			Base:        "factory/osmo1acqpnvg2t4wmqfdv8hq47d3petfksjs5r9t45p/buhbubu",
			Display:     "factory/osmo1acqpnvg2t4wmqfdv8hq47d3petfksjs5r9t45p/buhbubu",
			DenomUnits: []*banktypes.DenomUnit{{
				Denom:    "factory/osmo1acqpnvg2t4wmqfdv8hq47d3petfksjs5r9t45p/buhbubu",
				Exponent: 0,
				Aliases:  []string{"BUHBUBU"}}},
		},
	}
}
