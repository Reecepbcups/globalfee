package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkmath "cosmossdk.io/math"
)

func TestDefaultParams(t *testing.T) {
	p := DefaultParams()
	require.EqualValues(t, p.MinimumGasPrices, sdk.DecCoins(nil))
}

func Test_validateParams(t *testing.T) {
	tests := map[string]struct {
		coins     interface{} // not sdk.DeCoins, but Decoins defined in glboalfee
		expectErr bool
	}{
		"DefaultParams, pass": {
			DefaultParams().MinimumGasPrices,
			false,
		},
		"DecCoins conversion fails, fail": {
			sdk.Coins{sdk.NewCoin("photon", sdkmath.OneInt())},
			true,
		},
		"coins amounts are zero, pass": {
			sdk.DecCoins{
				sdk.NewDecCoin("atom", sdkmath.ZeroInt()),
				sdk.NewDecCoin("photon", sdkmath.ZeroInt()),
			},
			false,
		},
		"duplicate coins denoms, fail": {
			sdk.DecCoins{
				sdk.NewDecCoin("photon", sdkmath.OneInt()),
				sdk.NewDecCoin("photon", sdkmath.OneInt()),
			},
			true,
		},
		"coins are not sorted by denom alphabetically, fail": {
			sdk.DecCoins{
				sdk.NewDecCoin("photon", sdkmath.OneInt()),
				sdk.NewDecCoin("atom", sdkmath.OneInt()),
			},
			true,
		},
		"negative amount, fail": {
			sdk.DecCoins{
				sdk.DecCoin{Denom: "photon", Amount: sdkmath.LegacyOneDec().Neg()},
			},
			true,
		},
		"invalid denom, fail": {
			sdk.DecCoins{
				sdk.DecCoin{Denom: "photon!", Amount: sdkmath.LegacyOneDec().Neg()},
			},
			true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateMinimumGasPrices(test.coins)
			if test.expectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
