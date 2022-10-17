package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	ctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	tmjson "github.com/tendermint/tendermint/libs/json"
	pc "github.com/tendermint/tendermint/proto/tendermint/crypto"
	"github.com/tendermint/tendermint/types"
	recordtypes "gitlab.bianjie.ai/cschain/cschain/modules/ibc/applications/record/types"
	bcos "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/bcos/types"
	brochain "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/brochain/types"
	fabric "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/fabric/types"
	tendermint "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/tendermint/types"
	wutong "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/wutong/types"
)

var (
	appModules []module.AppModuleBasic
	encodecfg  params.EncodingConfig
)

// 初始化账户地址前缀
func MakeEncodingConfig() {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := ctypes.NewInterfaceRegistry()
	tmjson.RegisterType((*pc.PublicKey_Sm2)(nil), "tendermint.crypto.PublicKey_Sm2")
	brochain.RegisterInterfaces(interfaceRegistry)
	fabric.RegisterInterfaces(interfaceRegistry)
	tendermint.RegisterInterfaces(interfaceRegistry)
	wutong.RegisterInterfaces(interfaceRegistry)
	bcos.RegisterInterfaces(interfaceRegistry)
	recordtypes.RegisterInterfaces(interfaceRegistry)
	moduleBasics := module.NewBasicManager(appModules...)
	moduleBasics.RegisterInterfaces(interfaceRegistry)
	moduleBasics.RegisterLegacyAminoCodec(amino)
	std.RegisterInterfaces(interfaceRegistry)
	std.RegisterLegacyAminoCodec(amino)
	sdk.RegisterInterfaces(interfaceRegistry)
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)
	encodecfg = params.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}

func GetTxDecoder() sdk.TxDecoder {
	return encodecfg.TxConfig.TxDecoder()
}

func GetMarshaler() codec.Marshaler {
	return encodecfg.Marshaler
}

func GetSigningTx(txBytes types.Tx) (signing.Tx, error) {
	Tx, err := GetTxDecoder()(txBytes)
	if err != nil {
		return nil, err
	}
	signingTx := Tx.(signing.Tx)
	return signingTx, nil
}

func RegisterAppModules(module ...module.AppModuleBasic) {
	appModules = append(appModules, module...)
}
