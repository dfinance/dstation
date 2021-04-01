package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DumbGasMeter is a sdk.GasMeter implementation that doesn't count amount of gas during usage.
type DumbGasMeter struct{}

func (g DumbGasMeter) GasConsumed() sdk.Gas {
	return 0
}

func (g DumbGasMeter) Limit() sdk.Gas {
	return 0
}

func (g DumbGasMeter) GasConsumedToLimit() sdk.Gas {
	return 0
}

func (g *DumbGasMeter) ConsumeGas(_ sdk.Gas, _ string) {
}

func (g DumbGasMeter) IsPastLimit() bool {
	return false
}

func (g DumbGasMeter) IsOutOfGas() bool {
	return false
}

func (g DumbGasMeter) String() string {
	return "DumbGasMeter"
}

func NewDumbGasMeter() sdk.GasMeter {
	return &DumbGasMeter{}
}
