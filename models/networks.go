package models

import (
	AeternityValidator "github.com/seyitahmetgkc/gowallval/validators/ae_validator"
	AlgorandValidator "github.com/seyitahmetgkc/gowallval/validators/algo_validator"
	ArdorValidator "github.com/seyitahmetgkc/gowallval/validators/ardr_validator"
	AtomValidator "github.com/seyitahmetgkc/gowallval/validators/atom_validator"
	BitcoinValidator "github.com/seyitahmetgkc/gowallval/validators/bitcoin_validator"
	BEP2Validator "github.com/seyitahmetgkc/gowallval/validators/bnb_validator"
	//BTSValidator "github.com/seyitahmetgkc/gowallval/validators/bts_validator"
	CardanoValidator "github.com/seyitahmetgkc/gowallval/validators/cardano_validator"
	EOSValidator "github.com/seyitahmetgkc/gowallval/validators/eos_validator"
	EthereumValidator "github.com/seyitahmetgkc/gowallval/validators/ethereum_validator"
	HBARValidator "github.com/seyitahmetgkc/gowallval/validators/hbar_validator"
	ICONValidator "github.com/seyitahmetgkc/gowallval/validators/icx_validator"
	IOSTValidator "github.com/seyitahmetgkc/gowallval/validators/iost_validator"
	LiskValidator "github.com/seyitahmetgkc/gowallval/validators/lisk_validator"
	StellarValidator "github.com/seyitahmetgkc/gowallval/validators/lumen_validator"
	NanoValidator "github.com/seyitahmetgkc/gowallval/validators/nano_validator"
	NEMValidator "github.com/seyitahmetgkc/gowallval/validators/nem_validator"
	RippleValidator "github.com/seyitahmetgkc/gowallval/validators/ripple_validator"
	SiacoinValidator "github.com/seyitahmetgkc/gowallval/validators/sc_validator"
	SteemValidator "github.com/seyitahmetgkc/gowallval/validators/steem_validator"
	SyscoinValidator "github.com/seyitahmetgkc/gowallval/validators/sys_validator"
	ZilliqaValidator "github.com/seyitahmetgkc/gowallval/validators/zil_validator"
)

var AE = Network{
	Name:      "Aeternity",
	Symbol:    "æ",
	Validator: AeternityValidator.IsValidAddress,
}

var ALGO = Network{
	Name:      "Algorand",
	Symbol:    "ALGO",
	Validator: AlgorandValidator.IsValidAddress,
}

var ARDR = Network{
	Name:      "Ardor",
	Symbol:    "ARDR",
	Validator: ArdorValidator.IsValidAddress,
}

var ATOM = Network{
	Name:      "Cosmos",
	Symbol:    "ATOM",
	Validator: AtomValidator.IsValidAddress,
}

var BSC = Network{
	Name:      "Binance Smart Chain (BEP20)",
	Symbol:    "BSC",
	Validator: EthereumValidator.IsValidAddress,
}

var BTC = Network{
	Name:      "Bitcoin",
	Symbol:    "BTC",
	Validator: BitcoinValidator.IsValidAddress,
}

var BNB = Network{
	Name:      "Binance Beacon Chain (BEP2)",
	Symbol:    "BNB",
	Validator: BEP2Validator.IsValidAddress,
}

var ETH = Network{
	Name:      "Ethereum (ERC20)",
	Symbol:    "ETH",
	Validator: EthereumValidator.IsValidAddress,
}

var SEGWITBTC = Network{
	Name:      "Bitcoin (SegWit)",
	Symbol:    "BTC",
	Validator: BitcoinValidator.IsValidAddress,
}

/*var BTS = Network{
	Name:      "BitShares",
	Symbol:    "BTS",
	Validator: BTSValidator.IsValidAddress,
}*/

var ADA = Network{
	Name:      "Cardano",
	Symbol:    "ADA",
	Validator: CardanoValidator.IsValidAddress,
}

var EOS = Network{
	Name:      "EOS",
	Symbol:    "EOS",
	Validator: EOSValidator.IsValidAddress,
}

var HBAR = Network{
	Name:      "Hedera Hashgraph",
	Symbol:    "HBAR",
	Validator: HBARValidator.IsValidAddress,
}

var ICX = Network{
	Name:      "ICON",
	Symbol:    "ICX",
	Validator: ICONValidator.IsValidAddress,
}

var IOST = Network{
	Name:      "IOST",
	Symbol:    "IOST",
	Validator: IOSTValidator.IsValidAddress,
}

var LSK = Network{
	Name:      "Lisk",
	Symbol:    "LSK",
	Validator: LiskValidator.IsValidAddress,
}

var XLM = Network{
	Name:      "Stellar Lumens",
	Symbol:    "XLM",
	Validator: StellarValidator.IsValidAddress,
}

var NANO = Network{
	Name:      "Nano",
	Symbol:    "NANO",
	Validator: NanoValidator.IsValidAddress,
}

var XEM = Network{
	Name:      "NEM",
	Symbol:    "XEM",
	Validator: NEMValidator.IsValidAddress,
}

var XRP = Network{
	Name:      "Ripple",
	Symbol:    "XRP",
	Validator: RippleValidator.IsValidAddress,
}

var SC = Network{
	Name:      "Siacoin",
	Symbol:    "SC",
	Validator: SiacoinValidator.IsValidAddress,
}

var STEEM = Network{
	Name:      "Steem",
	Symbol:    "STEEM",
	Validator: SteemValidator.IsValidAddress,
}

var SYS = Network{
	Name:      "Syscoin",
	Symbol:    "SYS",
	Validator: SyscoinValidator.IsValidAddress,
}

var ZIL = Network{
	Name:      "Zilliqa",
	Symbol:    "ZIL",
	Validator: ZilliqaValidator.IsValidAddress,
}

var Networks = make(map[string]Network)
var UnsupportedNetworks = []string{"XTZ", "STX", "IOTA", "XMR", "NXS", "BTS"}

func InitializeNetworks() {
	Networks["AE"] = AE
	Networks["ALGO"] = ALGO
	Networks["ARDR"] = ARDR
	Networks["ATOM"] = ATOM
	Networks["BSC"] = BSC
	Networks["BTC"] = BTC
	Networks["BNB"] = BNB
	Networks["ETH"] = ETH
	Networks["SEGWITBTC"] = SEGWITBTC
	//Networks["BTS"] = BTS
	Networks["ADA"] = ADA
	Networks["EOS"] = EOS
	Networks["HBAR"] = HBAR
	Networks["ICX"] = ICX
	Networks["IOST"] = IOST
	Networks["LSK"] = LSK
	Networks["XLM"] = XLM
	Networks["NANO"] = NANO
	Networks["XEM"] = XEM
	Networks["XRP"] = XRP
	Networks["SC"] = SC
	Networks["STEEM"] = STEEM
	Networks["SYS"] = SYS
	Networks["ZIL"] = ZIL
}
