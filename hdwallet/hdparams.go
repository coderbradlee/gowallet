package hdwallet

import (
	"github.com/btcsuite/btcd/chaincfg"
	"time"
	"github.com/btcsuite/btcd/wire"
)

const (
	// MainNet represents the main litecoin network.
	LTCMainNet wire.BitcoinNet = 0xdbb6c0fb

	// MainNet represents the main litecoin network.
	DOGEMainNet wire.BitcoinNet = 0x4081889


	// MainNet represents the main litecoin network.
	QTUMMainNet wire.BitcoinNet = 0xb081889

	// MainNet represents the main litecoin network.
	HCASHMainNet wire.BitcoinNet = 0xb091989
)

var BtcNetParams = chaincfg.MainNetParams

var LtcNetParams = chaincfg.Params{
	Name:        "mainnet",
	Net:         LTCMainNet,
	DefaultPort: "9333",
	DNSSeeds: []chaincfg.DNSSeed{
		{"seed-a.litecoin.loshan.co.uk", true},
		{"dnsseed.thrasher.io", true},
		{"dnsseed.litecointools.com", false},
		{"dnsseed.litecoinpool.org", false},
		{"dnsseed.koin-project.com", false},
	},

	PowLimitBits:             504365055,
	BIP0034Height:            710000,
	BIP0065Height:            918684,
	BIP0066Height:            811879,
	CoinbaseMaturity:         100,
	SubsidyReductionInterval: 840000,
	TargetTimespan:           (time.Hour * 24 * 3) + (time.Hour * 12), // 3.5 days
	TargetTimePerBlock:       (time.Minute * 2) + (time.Second * 30),  // 2.5 minutes
	RetargetAdjustmentFactor: 4,                                       // 25% less, 400% more
	ReduceMinDifficulty:      false,
	MinDiffReductionTime:     0,
	GenerateSupported:        false,


	// Consensus rule change deployments.
	//
	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationThreshold: 6048, // 75% of MinerConfirmationWindow
	MinerConfirmationWindow:       8064, //

	// Mempool parameters
	RelayNonStdTxs: false,

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "ltc", // always ltc for main net

	// Address encoding magics
	PubKeyHashAddrID:        0x30, // starts with L
	ScriptHashAddrID:        0x50, // starts with M
	PrivateKeyID:            0xB0, // starts with 6 (uncompressed) or T (compressed)
	WitnessPubKeyHashAddrID: 0x06, // starts with p2
	WitnessScriptHashAddrID: 0x0A, // starts with 7Xh

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x88, 0xad, 0xe4}, // starts with xprv
	HDPublicKeyID:  [4]byte{0x04, 0x88, 0xb2, 0x1e}, // starts with xpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 2,
}

var DogeNetParams = chaincfg.Params{
	Name:        "mainnet",
	Net:         DOGEMainNet,
	DefaultPort: "9333",
	DNSSeeds: []chaincfg.DNSSeed{
		{"seed-a.litecoin.loshan.co.uk", true},
		{"dnsseed.thrasher.io", true},
		{"dnsseed.litecointools.com", false},
		{"dnsseed.litecoinpool.org", false},
		{"dnsseed.koin-project.com", false},
	},

	// Mempool parameters
	RelayNonStdTxs: false,

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "doge", // always ltc for main net

	// Address encoding magics
	PubKeyHashAddrID:        0x1e, // starts with L
	ScriptHashAddrID:        0x16, // starts with M
	PrivateKeyID:            0x9e, // starts with 6 (uncompressed) or T (compressed)
	WitnessPubKeyHashAddrID: 0x06, // starts with p2
	WitnessScriptHashAddrID: 0x0A, // starts with 7Xh

	// BIP32 hierarchical deterministic extended key magics
	HDPublicKeyID:  [4]byte{0x02, 0x88, 0xc3, 0x98}, // starts with xpub
	HDPrivateKeyID: [4]byte{0x02, 0xfa, 0xca, 0xfd}, // starts with xprv

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 2,
}

var QtumNetParams = chaincfg.Params{
	Name:        "main",
	Net:         QTUMMainNet,
	DefaultPort: "3888",
	DNSSeeds: []chaincfg.DNSSeed{
		{"seed-a.litecoin.loshan.co.uk", true},
		{"dnsseed.thrasher.io", true},
		{"dnsseed.litecointools.com", false},
		{"dnsseed.litecoinpool.org", false},
		{"dnsseed.koin-project.com", false},
	},

	// Mempool parameters
	RelayNonStdTxs: false,

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "qtum", // always ltc for main net

	// Address encoding magics
	PubKeyHashAddrID:        0x3a, // starts with L
	ScriptHashAddrID:        0x32, // starts with M
	PrivateKeyID:            0x80, // starts with 6 (uncompressed) or T (compressed)
	WitnessPubKeyHashAddrID: 0x06, // starts with p2
	WitnessScriptHashAddrID: 0x0A, // starts with 7Xh

	// BIP32 hierarchical deterministic extended key magics
	HDPublicKeyID:  [4]byte{0x04, 0x88, 0xB2, 0x1E}, // starts with xpub
	HDPrivateKeyID: [4]byte{0x04, 0x88, 0xAD, 0xE4}, // starts with xprv

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 2,
}


func mustRegister(params *chaincfg.Params) {
	if err := chaincfg.Register(params); err != nil {
		panic("failed to register network: " + err.Error())
	}
}

func init() {
	// Register all default networks when the package is initialized.
	mustRegister(&LtcNetParams)
	mustRegister(&DogeNetParams)
	mustRegister(&QtumNetParams)

}
