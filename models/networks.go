package models

import (
	//. "gowallval/models"
	AeternityValidator "gowallval/validators/ae_validator"
	AlgorandValidator "gowallval/validators/algo_validator"
	ArdorValidator "gowallval/validators/ardr_validator"
	AtomValidator "gowallval/validators/atom_validator"
	EthereumValidator "gowallval/validators/ethereum_validator"
	BitcoinValidator "gowallval/validators/bitcoin_validator"
	BEP2Validator "gowallval/validators/bnb_validator"
	BTSValidator "gowallval/validators/bts_validator"
	CardanoValidator "gowallval/validators/cardano_validator"
	EOSValidator "gowallval/validators/eos_validator"
	HBARValidator "gowallval/validators/hbar_validator"
	ICONValidator "gowallval/validators/icx_validator"
	IOSTValidator "gowallval/validators/iost_validator"
	LiskValidator "gowallval/validators/lisk_validator"
	StellarValidator "gowallval/validators/lumen_validator"
	NanoValidator "gowallval/validators/nano_validator"
	NEMValidator "gowallval/validators/nem_validator"
	RippleValidator "gowallval/validators/ripple_validator"
	SiacoinValidator "gowallval/validators/sc_validator"
	SteemValidator "gowallval/validators/steem_validator"
	SyscoinValidator "gowallval/validators/sys_validator"
	ZilliqaValidator "gowallval/validators/zil_validator"
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

var BTS = Network{
	Name:      "BitShares",
	Symbol:    "BTS",
	Validator: BTSValidator.IsValidAddress,
}

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
var UnsupportedNetworks = []string{"XTZ", "STX", "IOTA", "XMR", "NXS"}

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
	Networks["BTS"] = BTS
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