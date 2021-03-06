package types

import (
	"encoding/hex"
	"fmt"

	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
)

func (m GenesisState_WriteOp) String() string {
	return fmt.Sprintf("%s::%s", m.Address, m.Path)
}

// ToBytes converts WriteOp to dvmTypes.VMAccessPath and []byte representation for value.
func (m GenesisState_WriteOp) ToBytes() (*dvmTypes.VMAccessPath, []byte, error) {
	bzAddr, err := hex.DecodeString(m.Address)
	if err != nil {
		return nil, nil, fmt.Errorf("address: hex decode: %w", err)
	}
	if len(bzAddr) != VMAddressLength {
		return nil, nil, fmt.Errorf("address: incorrect length (should be %d bytes)", VMAddressLength)
	}

	bzPath, err := hex.DecodeString(m.Path)
	if err != nil {
		return nil, nil, fmt.Errorf("path: hex decode: %w", err)
	}

	bzValue, err := hex.DecodeString(m.Value)
	if err != nil {
		return nil, nil, fmt.Errorf("value: hex decode: %w", err)
	}

	return &dvmTypes.VMAccessPath{
		Address: bzAddr,
		Path:    bzPath,
	}, bzValue, nil
}

// Validate checks that genesis state is valid.
func (m GenesisState) Validate() error {
	// VM writeSets
	writeOpsSet := make(map[string]struct{}, len(m.WriteSet))
	for woIdx, writeOp := range m.WriteSet {
		if _, _, err := writeOp.ToBytes(); err != nil {
			return fmt.Errorf("writeSet [%d]: %w", woIdx, err)
		}

		writeOpId := writeOp.String()
		if _, ok := writeOpsSet[writeOpId]; ok {
			return fmt.Errorf("writeSet [%d]: duplicated (%s)", woIdx, writeOpId)
		}
		writeOpsSet[writeOpId] = struct{}{}
	}

	return nil
}

// DefaultGenesisState returns default genesis state (validation is done on module init).
func DefaultGenesisState() GenesisState {
	return defaultGenesisState
}
