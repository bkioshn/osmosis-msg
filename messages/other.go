package message

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

func CraftStoreCode(acc client.Account) sdk.Msg {
	wasmFile, _ := os.ReadFile("/Users/bubu/Work/Alles/script/ts/tx-generator/src/assets/wasm/test_sc.wasm")

	mySlice := make([]string, 1)
	mySlice[0] = "osmo1a95l6k4mkmz5ere064cmatnvnumj782tl2jlle"

	return &wasmtypes.MsgStoreCode{
		Sender:       "osmo1a95l6k4mkmz5ere064cmatnvnumj782tl2jlle",
		WASMByteCode: wasmFile,
		InstantiatePermission: &wasmtypes.AccessConfig{
			Permission: wasmtypes.AccessTypeAnyOfAddresses,
			Addresses:  mySlice,
		},
	}
}
