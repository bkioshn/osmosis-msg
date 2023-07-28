package message

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	wasmpoolmodel "github.com/osmosis-labs/osmosis/v16/x/cosmwasmpool/model"
)

func CraftMsgCreateCosmWasmPool(acc client.Account) sdk.Msg {
	return &wasmpoolmodel.MsgCreateCosmWasmPool{
		Sender:         acc.GetAddress().String(),
		CodeId:         1,
		InstantiateMsg: []byte{},
	}
}
