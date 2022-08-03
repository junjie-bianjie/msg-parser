package integration

import (
	"encoding/hex"
	"fmt"
	"github.com/kaifei-bianjie/msg-parser/codec"
	. "github.com/kaifei-bianjie/msg-parser/codec"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

func (s IntegrationTestSuite) TestMT() {
	cases := []SubTest{
		{
			"MTIssueDenom",
			MTIssueDenom,
		},
		{
			"MTTransferDenom",
			MTTransferDenom,
		},
		{
			"MintMT",
			MintMT,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func MTIssueDenom(s IntegrationTestSuite) {
	codec.SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)

	txBytes, err := hex.DecodeString("0A730A710A192F697269736D6F642E6D742E4D7367497373756544656E6F6D12540A12746573745F69737375735F64656E6F6D5F31121269737375655F64656E6F6D5F646174615F311A2A6961613134723663716B6663756571337332367634636D717966777967397A716D6B7A7532336176657A12680A520A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2102888C55F022E079EA72E9B6C18690CBDB133D2FCB08B20AD02158E1FFF2C1198712040A02080118C4A70112120A0C0A05756E79616E120334303010C09A0C1A40A8AB7BA9983E19F6829F6E31A7D4C7CB5E6231F7D006FA1891E8AFC4B613488B40F3E21D3393C546C59A7BDFE46E21046544831E968A288C03FFA20B62339A41")
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	for _, msg := range authTx.GetMsgs() {
		if mtDoc, ok := s.Mt.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(mtDoc))
		}
	}
}

func MTTransferDenom(s IntegrationTestSuite) {
	codec.SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)

	txBytes, err := hex.DecodeString("0ABE010ABB010A1C2F697269736D6F642E6D742E4D73675472616E7366657244656E6F6D129A010A4066623565633238313233633866653739393962313735383861616431616235366638343838616462383036616163356634303430393032306532653833646538122A6961613134723663716B6663756571337332367634636D717966777967397A716D6B7A7532336176657A1A2A6961613132763337346E706A3972366B663776747A61763470786E3333746A366D34347764783075786412680A520A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2102888C55F022E079EA72E9B6C18690CBDB133D2FCB08B20AD02158E1FFF2C1198712040A02080118C6A70112120A0C0A05756E79616E120334303010C09A0C1A4002DEC141E7168E1442E2799894E70B403941F783AAAA6175B9720CAD247097B271DD141DF64D1D30D0FCE611D3C37E111851DEA6BF32B1D62B255BAF68E668E9")
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	for _, msg := range authTx.GetMsgs() {
		if mtDoc, ok := s.Mt.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(mtDoc))
		}
	}
}

func MintMT(s IntegrationTestSuite) {
	codec.SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)

	txBytes, err := hex.DecodeString("0ABF010ABC010A152F697269736D6F642E6D742E4D73674D696E744D5412A20112403835383636376633393436316430626633313131646430383032336332646636356130653132386665363361613634646230633430366166636133383334306118052204636376762A2A6961613134723663716B6663756571337332367634636D717966777967397A716D6B7A7532336176657A322A6961613134723663716B6663756571337332367634636D717966777967397A716D6B7A7532336176657A12680A520A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2102888C55F022E079EA72E9B6C18690CBDB133D2FCB08B20AD02158E1FFF2C1198712040A02080118CAA70112120A0C0A05756E79616E120334303010C09A0C1A40A1EF1715725E3DA04182933F27E5CB30311EE61C09AC211B83E061D2B220457731E1DAA03C53E14C0BDAEB6EA2D48ED75943B0C5366EAB480FE04F467523C005")
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	for _, msg := range authTx.GetMsgs() {
		if mtDoc, ok := s.Mt.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(mtDoc))
		}
	}
}