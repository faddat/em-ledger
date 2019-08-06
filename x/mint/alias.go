// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/mint/internal/keeper
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/mint/internal/types
package mint

import (
	"emoney/x/mint/internal/keeper"
	"emoney/x/mint/internal/types"
)

const (
	ModuleName            = types.ModuleName
	DefaultParamspace     = types.DefaultParamspace
	StoreKey              = types.StoreKey
	QuerierRoute          = types.QuerierRoute
	QueryParameters       = types.QueryParameters
	QueryInflation        = types.QueryInflation
	QueryAnnualProvisions = types.QueryAnnualProvisions
)

var (
	// functions aliases
	NewKeeper            = keeper.NewKeeper
	NewQuerier           = keeper.NewQuerier
	NewMinter            = types.NewMinter
	InitialMinter        = types.InitialMinter
	DefaultInitialMinter = types.DefaultInitialMinter
	ValidateMinter       = types.ValidateMinter
	ParamKeyTable        = types.ParamKeyTable
	NewParams            = types.NewParams
	DefaultParams        = types.DefaultParams
	ValidateParams       = types.ValidateParams

	// variable aliases
	ModuleCdc    = types.ModuleCdc
	MinterKey    = types.MinterKey
	KeyMintDenom = types.KeyMintDenom
)

type (
	Keeper = keeper.Keeper
	Minter = types.Minter
	Params = types.Params
)
