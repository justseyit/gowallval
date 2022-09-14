package models



var Aeternity = Currency{
	Symbol:  "AE",
	Name:    "Aeternity",
	Network: AE,
}

var Algorand = Currency{
	Symbol:  "ALGO",
	Name:    "Algorand",
	Network: ALGO,
}

var Ardor = Currency{
	Symbol:  "ARDR",
	Name:    "Ardor",
	Network: ARDR,
}

var Cosmos = Currency{
	Symbol:  "ATOM",
	Name:    "Cosmos",
	Network: ATOM,
}

var Bitcoin = Currency{
	Symbol:  "BTC",
	Name:    "Bitcoin",
	Network: BTC,
}

var BitcoinSegWit = Currency{
	Symbol:  "BTC",
	Name:    "Bitcoin",
	Network: SEGWITBTC,
}

var BitcoinBSC = Currency{
	Symbol:  "BTC",
	Name:    "Bitcoin",
	Network: BSC,
}

var BitcoinBNB = Currency{
	Symbol:  "BTC",
	Name:    "Bitcoin",
	Network: BNB,
}

var BitcoinETH = Currency{
	Symbol:  "BTC",
	Name:    "Bitcoin",
	Network: ETH,
}

var BitShares = Currency{
	Symbol:  "BTS",
	Name:    "BitShares",
	Network: BTS,
}

var Cardano = Currency{
	Symbol:  "ADA",
	Name:    "Cardano",
	Network: ADA,
}

var CardanoBSC = Currency{
	Symbol:  "ADA",
	Name:    "Cardano",
	Network: BSC,
}

var CardanoBNB = Currency{
	Symbol:  "ADA",
	Name:    "Cardano",
	Network: BNB,
}

var Eos = Currency{
	Symbol:  "EOS",
	Name:    "EOS",
	Network: EOS,
}

var EosBSC = Currency{
	Symbol:  "EOS",
	Name:    "EOS",
	Network: BSC,
}

var EosBNB = Currency{
	Symbol:  "EOS",
	Name:    "EOS",
	Network: BNB,
}

var Ethereum = Currency{
	Symbol:  "ETH",
	Name:    "Ethereum",
	Network: ETH,
}

var EthereumBSC = Currency{
	Symbol:  "ETH",
	Name:    "Ethereum",
	Network: BSC,
}

var EthereumBNB = Currency{
	Symbol:  "ETH",
	Name:    "Ethereum",
	Network: BNB,
}

var HederaHashgraph = Currency{
	Symbol:  "HBAR",
	Name:    "Hedera Hashgraph",
	Network: HBAR,
}

var Icon = Currency{
	Symbol:  "ICX",
	Name:    "ICON",
	Network: ICX,
}

var Iost = Currency{
	Symbol:  "IOST",
	Name:    "IOST",
	Network: IOST,
}

var Lisk = Currency{
	Symbol:  "LSK",
	Name:    "Lisk",
	Network: LSK,
}

var StellarLumens = Currency{
	Symbol:  "XLM",
	Name:    "Stellar Lumens",
	Network: XLM,
}

var StellarLumensBSC = Currency{
	Symbol:  "XLM",
	Name:    "Stellar Lumens",
	Network: BSC,
}

var Nano = Currency{
	Symbol:  "XNO",
	Name:    "Nano",
	Network: NANO,
}

var Nem = Currency{
	Symbol:  "XEM",
	Name:    "NEM",
	Network: XEM,
}

var Ripple = Currency{
	Symbol:  "XRP",
	Name:    "Ripple",
	Network: XRP,
}

var RippleBSC = Currency{
	Symbol:  "XRP",
	Name:    "Ripple",
	Network: BSC,
}

var RippleBNB = Currency{
	Symbol:  "XRP",
	Name:    "Ripple",
	Network: BNB,
}

var RippleETH = Currency{
	Symbol:  "XRP",
	Name:    "Ripple",
	Network: ETH,
}

var Syscoin = Currency{
	Symbol:  "SYS",
	Name:    "Syscoin",
	Network: SYS,
}

var Steem = Currency{
	Symbol:  "STEEM",
	Name:    "Steem",
	Network: STEEM,
}

var Zilliqa = Currency{
	Symbol:  "ZIL",
	Name:    "Zilliqa",
	Network: ZIL,
}

var ZilliqaBSC = Currency{
	Symbol:  "ZIL",
	Name:    "Zilliqa",
	Network: BSC,
}

var UnsupportedCurrencies = []string{"XTZ", "STX", "IOTA", "XMR", "NXS", "AUD"}

var Currencies = []Currency{
	Aeternity,
	Algorand,
	Ardor,
	Cosmos,
	Bitcoin,
	BitcoinSegWit,
	BitcoinBSC,
	BitcoinBNB,
	BitcoinETH,
	BitShares,
	Cardano,
	CardanoBSC,
	CardanoBNB,
	Eos,
	EosBSC,
	EosBNB,
	Ethereum,
	EthereumBSC,
	EthereumBNB,
	HederaHashgraph,
	Icon,
	Iost,
	Lisk,
	StellarLumens,
	StellarLumensBSC,
	Nano,
	Nem,
	Ripple,
	RippleBSC,
	RippleBNB,
	RippleETH,
	Syscoin,
	Steem,
	Zilliqa,
	ZilliqaBSC,
}