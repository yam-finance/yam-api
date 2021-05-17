package routes

import "math/big"

type Assets struct {
	Ugas    []AssetInstance
	Ustonks []AssetInstance
}
type Asset struct {
	AssetName     string
	AssetInstance AssetInstance
	AssetPrice    *big.Float
}
type AssetInstance struct {
	Name       string
	Cycle      string
	Year       string
	Collateral string
	Token      Token
	Emp        Emp
	Pool       Pool
	Apr        AprData
}
type Emp struct {
	Address string
	New     bool
}
type Pool struct {
	Address string
}
type AprData struct {
	Force int
	Extra int
}
type Token struct {
	Address  string
	Decimals int
}
type EmpInfo struct {
	Address           string
	CollateralAddress string
	Size              big.Int
	Price             big.Float
	Decimals          int
}
type responseAprDegenerative struct {
	UGAS map[string]float64 `json:"UGAS"`
}
type responseAprYam struct {
	Value float64 `json:"farm"`
}
