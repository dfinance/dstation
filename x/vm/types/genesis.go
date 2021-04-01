package types

import (
	"encoding/hex"
	"fmt"

	"github.com/dfinance/dvm-proto/go/vm_grpc"
)

func (m GenesisState_WriteOp) String() string {
	return fmt.Sprintf("%s::%s", m.Address, m.Path)
}

// ToBytes converts WriteOp to vm_grpc.VMAccessPath and []byte representation for value.
func (m GenesisState_WriteOp) ToBytes() (*vm_grpc.VMAccessPath, []byte, error) {
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

	return &vm_grpc.VMAccessPath{
		Address: bzAddr,
		Path:    bzPath,
	}, bzValue, nil
}

// Validate checks that genesis state is valid.
func (m GenesisState) Validate() error {
	writeOpsSet := make(map[string]bool, len(m.WriteSet))
	for woIdx, writeOp := range m.WriteSet {
		if _, _, err := writeOp.ToBytes(); err != nil {
			return fmt.Errorf("writeSet [%d]: %w", woIdx, err)
		}

		writeOpId := writeOp.String()
		if writeOpsSet[writeOpId] {
			return fmt.Errorf("writeSet [%d]: duplicated (%s)", woIdx, writeOpId)
		}
		writeOpsSet[writeOpId] = true
	}

	return nil
}

// DefaultGenesisState returns default genesis state (validation is done on module init).
func DefaultGenesisState() GenesisState {
	return defaultGenesisState
}
