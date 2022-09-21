package data


import (
	AeternityValidator "github.com/seyitahmetgkc/gowallval/validators/ae_validator"
	AlgorandValidator "github.com/seyitahmetgkc/gowallval/validators/algo_validator"
	ArdorValidator "github.com/seyitahmetgkc/gowallval/validators/ardr_validator"
	AtomValidator "github.com/seyitahmetgkc/gowallval/validators/atom_validator"
	BitcoinValidator "github.com/seyitahmetgkc/gowallval/validators/bitcoin_validator"
	BEP2Validator "github.com/seyitahmetgkc/gowallval/validators/bnb_validator"
	BTSValidator "github.com/seyitahmetgkc/gowallval/validators/bts_validator"
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
	RoninValidator "github.com/seyitahmetgkc/gowallval/validators/ronin_validator"
	LitecoinValidator "github.com/seyitahmetgkc/gowallval/validators/ltc_validator"
	ZcashValidator "github.com/seyitahmetgkc/gowallval/validators/zcash_validator"
	AvaxValidator "github.com/seyitahmetgkc/gowallval/validators/avax_validator"
	DashValidator "github.com/seyitahmetgkc/gowallval/validators/dash_validator"
	MoneroValidator "github.com/seyitahmetgkc/gowallval/validators/monero_validator"
	SolanaValidator "github.com/seyitahmetgkc/gowallval/validators/sol_validator"
	TronValidator "github.com/seyitahmetgkc/gowallval/validators/tron_validator"
	TezosValidator "github.com/seyitahmetgkc/gowallval/validators/xtz_validator"
	BandValidator "github.com/seyitahmetgkc/gowallval/validators/band_validator"
)

var Validators = make(map[string]func(string) bool)

func initializeValidators() {
	Validators["AE"] = AeternityValidator.IsValidAddress
	Validators["ALGO"] = AlgorandValidator.IsValidAddress
	Validators["ARDR"] = ArdorValidator.IsValidAddress
	Validators["ATOM"] = AtomValidator.IsValidAddress
	Validators["BSC"] = EthereumValidator.IsValidAddress
	Validators["BTC"] = BitcoinValidator.IsValidAddress
	Validators["BNB"] = BEP2Validator.IsValidAddress
	Validators["ETH"] = EthereumValidator.IsValidAddress
	Validators["SEGWITBTC"] = BitcoinValidator.IsValidAddress
	Validators["BTS"] = BTSValidator.IsValidAddress
	Validators["ADA"] = CardanoValidator.IsValidAddress
	Validators["EOS"] = EOSValidator.IsValidAddress
	Validators["HBAR"] = HBARValidator.IsValidAddress
	Validators["ICX"] = ICONValidator.IsValidAddress
	Validators["IOST"] = IOSTValidator.IsValidAddress
	Validators["LSK"] = LiskValidator.IsValidAddress
	Validators["XLM"] = StellarValidator.IsValidAddress
	Validators["NANO"] = NanoValidator.IsValidAddress
	Validators["XEM"] = NEMValidator.IsValidAddress
	Validators["XRP"] = RippleValidator.IsValidAddress
	Validators["SC"] = SiacoinValidator.IsValidAddress
	Validators["STEEM"] = SteemValidator.IsValidAddress
	Validators["SYS"] = SyscoinValidator.IsValidAddress
	Validators["ZIL"] = ZilliqaValidator.IsValidAddress
	Validators["RON"] = RoninValidator.IsValidAddress
	Validators["BCH"] = BitcoinValidator.IsValidAddress
	Validators["MATIC"] = EthereumValidator.IsValidAddress
	Validators["LTC"] = LitecoinValidator.IsValidAddress
	Validators["ZEC"] = ZcashValidator.IsValidAddress
	Validators["AVAX"] = AvaxValidator.IsValidAddress
	Validators["DASH"] = DashValidator.IsValidAddress
	Validators["XMR"] = MoneroValidator.IsValidAddress
	Validators["SOL"] = SolanaValidator.IsValidAddress
	Validators["TRX"] = TronValidator.IsValidAddress
	Validators["XTZ"] = TezosValidator.IsValidAddress
	Validators["BAND"] = BandValidator.IsValidAddress
}