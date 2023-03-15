package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/joho/godotenv"
	"github.com/osmosis-labs/osmosis/v15/app"
	poolmanagertypes "github.com/osmosis-labs/osmosis/v15/x/poolmanager/types"

	gammtypes "github.com/osmosis-labs/osmosis/v15/x/gamm/types"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func main() {
	// Set prefix
	sdk.GetConfig().SetBech32PrefixForAccount(os.Getenv("PREFIX"), "")
	sendTx()
}

func sendTx() {
	MNEMONIC := os.Getenv("MNEMONIC")
	RPC_PATH := os.Getenv("RPC_PATH")
	CHAIN_ID := os.Getenv("CHAIN_ID")
	GAS := os.Getenv("GAS")
	FEE := os.Getenv("FEE")

	gasInt, err := strconv.ParseInt(GAS, 10, 64)
	if err != nil {
		panic(err)
	}

	// Setup keyring
	kb := keyring.NewInMemory()
	path := sdk.GetConfig().GetFullBIP44Path()
	key, err := kb.NewAccount("alice", MNEMONIC, "", path, hd.Secp256k1)
	if err != nil {
		panic(err)
	}

	println("Account:", key.GetAddress().String())

	// Create client
	clientNode, err := client.NewClientFromNode(RPC_PATH)
	if err != nil {
		panic(err)
	}
	clientCtx := client.Context{
		Client:            clientNode,
		ChainID:           CHAIN_ID,
		NodeURI:           RPC_PATH,
		InterfaceRegistry: app.MakeEncodingConfig().InterfaceRegistry,
		TxConfig:          app.MakeEncodingConfig().TxConfig,
		Keyring:           kb,
	}

	// Retrieve account info
	accountRetriever := authtypes.AccountRetriever{}
	acc, err := accountRetriever.GetAccount(clientCtx, key.GetAddress())
	if err != nil {
		panic(err)
	}

	// Crafting message
	// Example: using MsgSwapExactAmountIn from osmosis gamm
	msgSwap := gammtypes.MsgSwapExactAmountIn{
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

	// Create transaction factory
	txf := tx.Factory{}.
		WithKeybase(kb).
		WithTxConfig(app.MakeEncodingConfig().TxConfig).
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithAccountNumber(acc.GetAccountNumber()).
		WithSequence(acc.GetSequence()).
		WithGas(uint64(gasInt)).WithGasAdjustment(2).
		WithChainID(CHAIN_ID).
		WithMemo("").
		WithFees(FEE).
		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT)

	txb, err := tx.BuildUnsignedTx(txf, &msgSwap)
	if err != nil {
		panic(err)
	}

	err = tx.Sign(txf, key.GetName(), txb, true)
	if err != nil {
		panic(err)
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(txb.GetTx())
	if err != nil {
		panic(err)
	}

	res, err := clientCtx.BroadcastTxCommit(txBytes)
	if err != nil {
		panic(err)
	}

	fmt.Println("response", res)
}
