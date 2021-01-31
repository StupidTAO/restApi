package model

type  MyDocument struct {
	Number int32	`json: "number"`
	Name string		`json: "name""`
	Score float64	`json: "score"`
	Median float64	`json: "median"`
	Weight float64	`json: "weight"`
	Datetime string	`json: "datetime"`
}

type tmpType map[string]string

type TokenId int

const (
	AOA TokenId = iota
	BAT
	BNB
	CRO
	DAI
	ENJ
	GNT
	HOT
	HT
	INB
	KCS
	LAMB
	LEO
	LINK
	MCO
	MKR
	NEXO
	NPXS
	PAX
	QNT
	REP
	SNT
	SNX
	TUSD
	USDC
	USDT
	ZB
	ZRX
	INVALID
)