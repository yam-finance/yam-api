package routes

import "math/big"

type Assets struct {
	Ugas    []AssetInstance `json:"ugas"`
	Ustonks []AssetInstance `json:"ustonks"`
	Upunks  []AssetInstance `json:"upunks"`
}
type AssetsFile struct {
	Assets Assets `json:"mainnet"`
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
	Expired    bool
}
type Emp struct {
	Address string
	New     bool
}
type Pool struct {
	Address string
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
