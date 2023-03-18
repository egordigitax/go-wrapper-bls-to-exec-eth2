package main

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	types "github.com/wealdtech/go-eth2-types/v2"
	"os/exec"
	"strconv"
)

type Credential struct {
	WithdrawalKey        *types.BLSPrivateKey
	SigningKey           *types.BLSPrivateKey
	ChainSetting         *BaseChainSetting
	HexWithdrawalAddress common.Address
}

func (c *Credential) Compute(index int) *BlsExecutionChangeResponse {

	err := types.InitBLS()
	if err != nil {
		return nil
	}
	out, err := exec.Command(
		"./wrapper/.venv/bin/python3",
		"./wrapper/wrapper.py",
		common.Bytes2Hex(c.WithdrawalKey.Marshal()),
		common.Bytes2Hex(c.SigningKey.Marshal()),
		strconv.FormatInt(int64(index), 10),
		c.ChainSetting.NetworkName,
		c.HexWithdrawalAddress.Hex(),
	).Output()
	if err != nil {
		println(err.Error())
	}
	textOut := string(out[:])

	println(textOut)
	response := &BlsExecutionChangeResponse{}
	err = json.Unmarshal(out, response)
	if err != nil {
		println(err.Error())
	}
	return response
}

type BlsExecutionChangeResponse struct {
	Message struct {
		ValidatorIndex     string `json:"validator_index"`
		FromBlsPubkey      string `json:"from_bls_pubkey"`
		ToExecutionAddress string `json:"to_execution_address"`
	} `json:"message"`
	Signature string `json:"signature"`
	Metadata  struct {
		NetworkName           string `json:"network_name"`
		GenesisValidatorsRoot string `json:"genesis_validators_root"`
		DepositCliVersion     string `json:"deposit_cli_version"`
	} `json:"metadata"`
}
