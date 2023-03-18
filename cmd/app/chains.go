package main

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	MAINNET  Chain = "mainnet"
	GOERLI         = "goerli"
	PRATER         = "prater"
	SEPOLIA        = "sepolia"
	ZHEJIANG       = "zhejiang"
)

type BaseChainSetting struct {
	NetworkName           string
	GenesisForkVersion    []byte
	GenesisValidatorsRoot []byte
}

type Chain = string

func GetAllChains() map[string]*BaseChainSetting {
	allChains := map[string]*BaseChainSetting{
		MAINNET:  MainnetSettings,
		GOERLI:   GoerliSetting,
		PRATER:   GoerliSetting,
		SEPOLIA:  SepoliaSetting,
		ZHEJIANG: ZhejiangSetting,
	}
	return allChains
}

func GetChainSetting(chainName string) *BaseChainSetting {
	return GetAllChains()[chainName]
}

var MainnetSettings = &BaseChainSetting{
	NetworkName:           MAINNET,
	GenesisForkVersion:    common.Hex2Bytes("00000000"),
	GenesisValidatorsRoot: common.Hex2Bytes("4b363db94e286120d76eb905340fdd4e54bfe9f06bf33ff6cf5ad27f511bfe95"),
}

var GoerliSetting = &BaseChainSetting{
	NetworkName:           MAINNET,
	GenesisForkVersion:    common.Hex2Bytes("00001020"),
	GenesisValidatorsRoot: common.Hex2Bytes("043db0d9a83813551ee2f33450d23797757d430911a9320530ad8a0eabc43efb"),
}

var SepoliaSetting = &BaseChainSetting{
	NetworkName:           MAINNET,
	GenesisForkVersion:    common.Hex2Bytes("90000069"),
	GenesisValidatorsRoot: common.Hex2Bytes("d8ea171f3c94aea21ebc42a1ed61052acf3f9209c00e4efbaaddac09ed9b8078"),
}

var ZhejiangSetting = &BaseChainSetting{
	NetworkName:           MAINNET,
	GenesisForkVersion:    common.Hex2Bytes("00000069"),
	GenesisValidatorsRoot: common.Hex2Bytes("53a92d8f2bb1d85f62d16a156e6ebcd1bcaba652d0900b2c2f387826f3481f6f"),
}
