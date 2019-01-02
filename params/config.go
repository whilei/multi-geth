// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Genesis hashes to enforce below configs on.
var (
	MainnetGenesisHash     = common.HexToHash("0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3")
	TestnetGenesisHash     = common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d")
	EllaismGenesisHash     = common.HexToHash("0x4d7df65052bb21264d6ad2d6fe2d5578a36be12f71bf8d0559b0c15c4dc539b5")
	SocialGenesisHash      = common.HexToHash("0xba8314d5c2ebddaf58eb882b364b27cbfa4d3402dacd32b60986754ac25cfe8d")
	MixGenesisHash         = common.HexToHash("0x4fa57903dad05875ddf78030c16b5da886f7d81714cf66946a4c02566dbb2af5")
	EthersocialGenesisHash = common.HexToHash("0x310dd3c4ae84dd89f1b46cfdd5e26c8f904dfddddc73f323b468127272e20e9f")
	RinkebyGenesisHash     = common.HexToHash("0x6341fd3daf94b748c72ced5a5b26028f2474f5f00d824504e4fa37a75767e177")
	KottiGenesisHash       = common.HexToHash("0xb75e5aceab7ec09ec4e94afb3e683bcc4e707a3c0958ad2d775d1ea936d7f232")
)

// TrustedCheckpoints associates each known checkpoint with the genesis hash of
// the chain it belongs to.
var TrustedCheckpoints = map[common.Hash]*TrustedCheckpoint{
	MainnetGenesisHash: MainnetTrustedCheckpoint,
	TestnetGenesisHash: TestnetTrustedCheckpoint,
	RinkebyGenesisHash: RinkebyTrustedCheckpoint,
	GoerliGenesisHash:  GoerliTrustedCheckpoint,
}

var (
	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(1),
		HomesteadBlock:      big.NewInt(1150000),
		DAOForkBlock:        big.NewInt(1920000),
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(2463000),
		EIP150Hash:          common.HexToHash("0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0"),
		EIP155Block:         big.NewInt(2675000),
		EIP158Block:         big.NewInt(2675000),
		ByzantiumBlock:      big.NewInt(4370000),
		DisposalBlock:       nil,
		SocialBlock:         nil,
		EthersocialBlock:    nil,
		ConstantinopleBlock: nil,
		Ethash:              new(EthashConfig),
	}

	// EllaismChainConfig is the chain parameters to run a node on the Ellaism main network.
	EllaismChainConfig = &ChainConfig{
		ChainID:             big.NewInt(64),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("0x4d7df65052bb21264d6ad2d6fe2d5578a36be12f71bf8d0559b0c15c4dc539b5"),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         nil,
		ByzantiumBlock:      nil,
		DisposalBlock:       big.NewInt(0),
		SocialBlock:         nil,
		EthersocialBlock:    nil,
		ConstantinopleBlock: nil,
		ECIP1017EraRounds:   big.NewInt(10000000),
		EIP160Block:         big.NewInt(0),
		Ethash:              new(EthashConfig),
	}

	// ClassicChainConfig is the chain parameters to run a node on the Ellaism main network.
	ClassicChainConfig = &ChainConfig{
		ChainID:             big.NewInt(61),
		HomesteadBlock:      big.NewInt(1150000),
		DAOForkBlock:        big.NewInt(1920000),
		DAOForkSupport:      false,
		EIP150Block:         big.NewInt(2500000),
		EIP150Hash:          common.HexToHash("0xca12c63534f565899681965528d536c52cb05b7c48e269c2a6cb77ad864d878a"),
		EIP155Block:         big.NewInt(3000000),
		EIP158Block:         nil,
		ByzantiumBlock:      nil,
		DisposalBlock:       big.NewInt(5900000),
		SocialBlock:         nil,
		EthersocialBlock:    nil,
		ConstantinopleBlock: nil,
		ECIP1017EraRounds:   big.NewInt(5000000),
		EIP160Block:         big.NewInt(3000000),
		ECIP1010PauseBlock:  big.NewInt(3000000),
		ECIP1010Length:      big.NewInt(2000000),
		Ethash:              new(EthashConfig),
	}

	// SocialChainConfig is the chain parameters to run a node on the Ethereum Social main network.
	SocialChainConfig = &ChainConfig{
		ChainID:             big.NewInt(28),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("0xba8314d5c2ebddaf58eb882b364b27cbfa4d3402dacd32b60986754ac25cfe8d"),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         nil,
		ByzantiumBlock:      nil,
		DisposalBlock:       big.NewInt(0),
		SocialBlock:         big.NewInt(0),
		EthersocialBlock:    nil,
		ConstantinopleBlock: nil,
		ECIP1017EraRounds:   big.NewInt(5000000),
		EIP160Block:         big.NewInt(0),
		Ethash:              new(EthashConfig),
	}

	// MainnetTrustedCheckpoint contains the light client trusted checkpoint for the main network.
	MainnetTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "mainnet",
		SectionIndex: 195,
		SectionHead:  common.HexToHash("0x1cdd2a84cf6c1261ffccc88f6bcefb513abd7934a96c1e909fbf74767560f16b"),
		CHTRoot:      common.HexToHash("0xe453333c20391d16b91b6fe11c104704f62c8dba15f69db73b4cdf7e100105eb"),
		BloomRoot:    common.HexToHash("0x47f30069473072e00d2cdca146dce40f0aad243dfc8221bf810822c091674efe"),
	}

	// MixChainConfig is the chain parameters to run a node on the MIX main network.
	MixChainConfig = &ChainConfig{
		ChainID:             big.NewInt(76),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      false,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("0x4fa57903dad05875ddf78030c16b5da886f7d81714cf66946a4c02566dbb2af5"),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      nil,
		DisposalBlock:       nil,
		SocialBlock:         nil,
		EthersocialBlock:    nil,
		ConstantinopleBlock: nil,
		EIP160Block:         big.NewInt(0),
	}

	// EthersocialChainConfig is the chain parameters to run a node on the Ethersocial main network.
	EthersocialChainConfig = &ChainConfig{
		ChainID:             big.NewInt(31102),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        big.NewInt(0),
		DAOForkSupport:      false,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("0x310dd3c4ae84dd89f1b46cfdd5e26c8f904dfddddc73f323b468127272e20e9f"),
		EIP155Block:         big.NewInt(845000),
		EIP158Block:         big.NewInt(845000),
		ByzantiumBlock:      big.NewInt(600000),
		DisposalBlock:       nil,
		SocialBlock:         nil,
		EthersocialBlock:    big.NewInt(0),
		ConstantinopleBlock: nil,
		Ethash:              new(EthashConfig),
	}

	// TestnetChainConfig contains the chain parameters to run a node on the Ropsten test network.
	TestnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(3),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d"),
		EIP155Block:         big.NewInt(10),
		EIP158Block:         big.NewInt(10),
		ByzantiumBlock:      big.NewInt(1700000),
		DisposalBlock:       nil,
		SocialBlock:         nil,
		EthersocialBlock:    nil,
		ConstantinopleBlock: big.NewInt(4230000),
		Ethash:              new(EthashConfig),
	}

	// TestnetTrustedCheckpoint contains the light client trusted checkpoint for the Ropsten test network.
	TestnetTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "testnet",
		SectionIndex: 161,
		SectionHead:  common.HexToHash("0x5378afa734e1feafb34bcca1534c4d96952b754579b96a4afb23d5301ecececc"),
		CHTRoot:      common.HexToHash("0x1cf2b071e7443a62914362486b613ff30f60cea0d9c268ed8c545f876a3ee60c"),
		BloomRoot:    common.HexToHash("0x5ac25c84bd18a9cbe878d4609a80220f57f85037a112644532412ba0d498a31b"),
	}

	// RinkebyChainConfig contains the chain parameters to run a node on the Rinkeby test network.
	RinkebyChainConfig = &ChainConfig{
		ChainID:             big.NewInt(4),
		HomesteadBlock:      big.NewInt(1),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x9b095b36c15eaf13044373aef8ee0bd3a382a5abb92e402afa44b8249c3a90e9"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(1035301),
		DisposalBlock:       nil,
		SocialBlock:         nil,
		EthersocialBlock:    nil,
		ConstantinopleBlock: nil,
		Clique: &CliqueConfig{
			Period: 15,
			Epoch:  30000,
		},
	}

	// RinkebyTrustedCheckpoint contains the light client trusted checkpoint for the Rinkeby test network.
	RinkebyTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "rinkeby",
		SectionIndex: 125,
		SectionHead:  common.HexToHash("0x8a738386f6bb34add15846f8f49c4c519a2f32519096e792b9f43bcb407c831c"),
		CHTRoot:      common.HexToHash("0xa1e5720a9bad4dce794f129e4ac6744398197b652868011486a6f89c8ec84a75"),
		BloomRoot:    common.HexToHash("0xa3048fe8b7e30f77f11bc755a88478363d7d3e71c2bdfe4e8ab9e269cd804ba2"),
	}

	// GoerliChainConfig contains the chain parameters to run a node on the Görli test network.
	GoerliChainConfig = &ChainConfig{
		ChainID:             big.NewInt(5),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		Clique: &CliqueConfig{
			Period: 15,
			Epoch:  30000,
		},
	}

	// GoerliTrustedCheckpoint contains the light client trusted checkpoint for the Görli test network.
	GoerliTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "goerli",
		SectionIndex: 9,
		SectionHead:  common.HexToHash("0x8e223d827391eee53b07cb8ee057dbfa11c93e0b45352188c783affd7840a921"),
		CHTRoot:      common.HexToHash("0xe0a817ac69b36c1e437c5b0cff9e764853f5115702b5f66d451b665d6afb7e78"),
		BloomRoot:    common.HexToHash("0x50d672aeb655b723284969c7c1201fb6ca003c23ed144bcb9f2d1b30e2971c1b"),
	}

	// KottiChainConfig is the chain parameters to run a node on the Ellaism main network.
	KottiChainConfig = &ChainConfig{
		ChainID:             big.NewInt(6),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      false,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("0xb75e5aceab7ec09ec4e94afb3e683bcc4e707a3c0958ad2d775d1ea936d7f232"),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         nil,
		ByzantiumBlock:      nil,
		DisposalBlock:       big.NewInt(0),
		SocialBlock:         nil,
		EthersocialBlock:    nil,
		ConstantinopleBlock: nil,
		ECIP1017EraRounds:   big.NewInt(5000000),
		EIP160Block:         big.NewInt(0),
		ECIP1010PauseBlock:  big.NewInt(0),
		ECIP1010Length:      big.NewInt(2000000),
		Clique: &CliqueConfig{
			Period: 15,
			Epoch:  30000,
		},
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Ethash consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllEthashProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, nil, nil, nil, nil, nil, nil, nil, new(EthashConfig), nil}

	// AllCliqueProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Clique consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllCliqueProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), nil, nil, nil, big.NewInt(0), big.NewInt(0), nil, nil, nil, nil, nil, nil, &CliqueConfig{Period: 0, Epoch: 30000}}

	TestChainConfig = &ChainConfig{big.NewInt(1), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), nil, nil, nil, big.NewInt(0), big.NewInt(0), nil, nil, nil, nil, nil, new(EthashConfig), nil}
	TestRules       = TestChainConfig.Rules(new(big.Int))
)

// TrustedCheckpoint represents a set of post-processed trie roots (CHT and
// BloomTrie) associated with the appropriate section index and head hash. It is
// used to start light syncing from this checkpoint and avoid downloading the
// entire header chain while still being able to securely access old headers/logs.
type TrustedCheckpoint struct {
	Name         string      `json:"-"`
	SectionIndex uint64      `json:"sectionIndex"`
	SectionHead  common.Hash `json:"sectionHead"`
	CHTRoot      common.Hash `json:"chtRoot"`
	BloomRoot    common.Hash `json:"bloomRoot"`
}

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection

	HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	DAOForkBlock   *big.Int `json:"daoForkBlock,omitempty"`   // TheDAO hard-fork switch block (nil = no fork)
	DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int    `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	EIP155Block      *big.Int `json:"eip155Block,omitempty"`      // EIP155 HF block
	EIP158Block      *big.Int `json:"eip158Block,omitempty"`      // EIP158 HF block
	DisposalBlock    *big.Int `json:"disposalBlock,omitempty"`    // Bomb disposal HF block
	SocialBlock      *big.Int `json:"socialBlock,omitempty"`      // Ethereum Social Reward block
	EthersocialBlock *big.Int `json:"ethersocialBlock,omitempty"` // Ethersocial Reward block

	ByzantiumBlock      *big.Int `json:"byzantiumBlock,omitempty"`      // Byzantium switch block (nil = no fork, 0 = already on byzantium)
	ConstantinopleBlock *big.Int `json:"constantinopleBlock,omitempty"` // Constantinople switch block (nil = no fork, 0 = already activated)
	PetersburgBlock     *big.Int `json:"petersburgBlock,omitempty"`     // Petersburg switch block (nil = same as Constantinople)
	EWASMBlock          *big.Int `json:"ewasmBlock,omitempty"`          // EWASM switch block (nil = no fork, 0 = already activated)

	ECIP1017EraRounds *big.Int `json:"ecip1017EraRounds,omitempty"` // ECIP1017 era rounds
	EIP160Block       *big.Int `json:"eip160Block,omitempty"`       // EIP160 HF block

	ECIP1010PauseBlock *big.Int `json:"ecip1010PauseBlock,omitempty"` // ECIP1010 pause HF block
	ECIP1010Length     *big.Int `json:"ecip1010Length,omitempty"`     // ECIP1010 length

	// Various consensus engines
	Ethash *EthashConfig `json:"ethash,omitempty"`
	Clique *CliqueConfig `json:"clique,omitempty"`
}

// EthashConfig is the consensus engine configs for proof-of-work based sealing.
type EthashConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *EthashConfig) String() string {
	return "ethash"
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Ethash != nil:
		engine = c.Ethash
	case c.Clique != nil:
		engine = c.Clique
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v Homestead: %v DAO: %v DAOSupport: %v EIP150: %v EIP155: %v EIP158: %v Byzantium: %v Disposal: %v Social: %v Ethersocial: %v ECIP1017: %v EIP160: %v ECIP1010PauseBlock: %v ECIP1010Length: %v Constantinople: %v Engine: %v}",
		c.ChainID,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		c.DisposalBlock,
		c.SocialBlock,
		c.EthersocialBlock,
		c.ECIP1017EraRounds,
		c.EIP160Block,
		c.ECIP1010PauseBlock,
		c.ECIP1010Length,
		c.ConstantinopleBlock,
		c.PetersburgBlock,
		engine,
	)
}

// HasECIP1017 returns whether the chain is configured with ECIP1017.
func (c *ChainConfig) HasECIP1017() bool {
	if c.ECIP1017EraRounds == nil {
		return false
	} else {
		return true
	}
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	return isForked(c.HomesteadBlock, num)
}

// IsEIP160 returns whether num is either equal to the EIP160 block or greater.
func (c *ChainConfig) IsEIP160(num *big.Int) bool {
	return isForked(c.EIP160Block, num)
}

// IsDAO returns whether num is either equal to the DAO fork block or greater.
func (c *ChainConfig) IsDAOFork(num *big.Int) bool {
	return isForked(c.DAOForkBlock, num)
}

// IsEIP150 returns whether num is either equal to the EIP150 fork block or greater.
func (c *ChainConfig) IsEIP150(num *big.Int) bool {
	return isForked(c.EIP150Block, num)
}

// IsEIP155 returns whether num is either equal to the EIP155 fork block or greater.
func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	return isForked(c.EIP155Block, num)
}

// IsEIP158 returns whether num is either equal to the EIP158 fork block or greater.
func (c *ChainConfig) IsEIP158(num *big.Int) bool {
	return isForked(c.EIP158Block, num)
}

// IsByzantium returns whether num is either equal to the Byzantium fork block or greater.
func (c *ChainConfig) IsByzantium(num *big.Int) bool {
	return isForked(c.ByzantiumBlock, num)
}

// IsConstantinople returns whether num is either equal to the Constantinople fork block or greater.
func (c *ChainConfig) IsConstantinople(num *big.Int) bool {
	return isForked(c.ConstantinopleBlock, num)
}

func (c *ChainConfig) IsBombDisposal(num *big.Int) bool {
	return isForked(c.DisposalBlock, num)
}

func (c *ChainConfig) IsSocial(num *big.Int) bool {
	return isForked(c.SocialBlock, num)
}

func (c *ChainConfig) IsEthersocial(num *big.Int) bool {
	return isForked(c.EthersocialBlock, num)
}

func (c *ChainConfig) IsECIP1010(num *big.Int) bool {
	return isForked(c.ECIP1010PauseBlock, num)
}

// IsEWASM returns whether num represents a block number after the EWASM fork
func (c *ChainConfig) IsEWASM(num *big.Int) bool {
	return isForked(c.EWASMBlock, num)
}

// GasTable returns the gas table corresponding to the current phase (homestead or homestead reprice).
//
// The returned GasTable's fields shouldn't, under any circumstances, be changed.
func (c *ChainConfig) GasTable(num *big.Int) GasTable {
	if num == nil {
		return GasTableHomestead
	}
	switch {
	case c.IsEIP160(num):
		return GasTableEIP160
	case c.IsEIP158(num):
		return GasTableEIP158
	case c.IsEIP150(num):
		return GasTableEIP150
	default:
		return GasTableHomestead
	}
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	if isForkIncompatible(c.HomesteadBlock, newcfg.HomesteadBlock, head) {
		return newCompatError("Homestead fork block", c.HomesteadBlock, newcfg.HomesteadBlock)
	}
	if isForkIncompatible(c.DAOForkBlock, newcfg.DAOForkBlock, head) {
		return newCompatError("DAO fork block", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if c.IsDAOFork(head) && c.DAOForkSupport != newcfg.DAOForkSupport {
		return newCompatError("DAO fork support flag", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if isForkIncompatible(c.EIP150Block, newcfg.EIP150Block, head) {
		return newCompatError("EIP150 fork block", c.EIP150Block, newcfg.EIP150Block)
	}
	if isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
		return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isForkIncompatible(c.EIP158Block, newcfg.EIP158Block, head) {
		return newCompatError("EIP158 fork block", c.EIP158Block, newcfg.EIP158Block)
	}
	if c.IsEIP158(head) && !configNumEqual(c.ChainID, newcfg.ChainID) {
		return newCompatError("EIP158 chain ID", c.EIP158Block, newcfg.EIP158Block)
	}
	if isForkIncompatible(c.ByzantiumBlock, newcfg.ByzantiumBlock, head) {
		return newCompatError("Byzantium fork block", c.ByzantiumBlock, newcfg.ByzantiumBlock)
	}
	if isForkIncompatible(c.ConstantinopleBlock, newcfg.ConstantinopleBlock, head) {
		return newCompatError("Constantinople fork block", c.ConstantinopleBlock, newcfg.ConstantinopleBlock)
	}
	if isForkIncompatible(c.PetersburgBlock, newcfg.PetersburgBlock, head) {
		return newCompatError("ConstantinopleFix fork block", c.PetersburgBlock, newcfg.PetersburgBlock)
	}
	if isForkIncompatible(c.EWASMBlock, newcfg.EWASMBlock, head) {
		return newCompatError("ewasm fork block", c.EWASMBlock, newcfg.EWASMBlock)
	}
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntactic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainID                                     *big.Int
	IsHomestead, IsEIP150, IsEIP155, IsEIP158   bool
	IsByzantium, IsConstantinople, IsPetersburg bool
}

// Rules ensures c's ChainID is not nil.
func (c *ChainConfig) Rules(num *big.Int) Rules {
	chainID := c.ChainID
	if chainID == nil {
		chainID = new(big.Int)
	}
	return Rules{
		ChainID:          new(big.Int).Set(chainID),
		IsHomestead:      c.IsHomestead(num),
		IsEIP150:         c.IsEIP150(num),
		IsEIP155:         c.IsEIP155(num),
		IsEIP158:         c.IsEIP158(num),
		IsByzantium:      c.IsByzantium(num),
		IsConstantinople: c.IsConstantinople(num),
		IsPetersburg:     c.IsPetersburg(num),
	}
}
