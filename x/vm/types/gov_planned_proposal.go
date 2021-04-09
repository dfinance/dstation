package types

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/gogo/protobuf/proto"
)

const (
	ProposalTypeVMPlanned = "VMPlannedProposal"
)

var (
	_ govTypes.Content = (*PlannedProposal)(nil)
)

// GetTitle implements govTypes.Content.
func (m PlannedProposal) GetTitle() string       {
	return "VM planned proposal: " + m.GetContent().GetTitle()
}

// GetDescription implements govTypes.Content.
func (m PlannedProposal) GetDescription() string {
	return m.GetContent().GetDescription()
}

// ProposalRoute implements govTypes.Content.
func (m PlannedProposal) ProposalRoute() string  {
	return RouterKey
}

// ProposalType implements govTypes.Content.
func (m PlannedProposal) ProposalType() string   {
	return ProposalTypeVMPlanned
}

// ValidateBasic implements govTypes.Content.
func (m PlannedProposal) ValidateBasic() error {
	if m.Height <= 0 {
		return fmt.Errorf( "height: should be GT 0")
	}

	if m.Content == nil {
		return fmt.Errorf( "content: nil")
	}
	content, ok := m.Content.GetCachedValue().(govTypes.Content)
	if !ok {
		return fmt.Errorf( "content: %T does not implement govTypes.Content", m.Content)
	}

	if err := content.ValidateBasic(); err != nil {
		return fmt.Errorf( "content: invalid: %w", err)
	}

	switch content.ProposalType() {
	case ProposalTypeStdlibUpdate:
	default:
		return fmt.Errorf( "content: unknown ProposalType: %s", content.ProposalType())
	}

	return nil
}

// String implements govTypes.Content.
func (m PlannedProposal) String() string {
	b := strings.Builder{}
	b.WriteString("PlannedProposal:\n")
	b.WriteString(fmt.Sprintf("  BlockHeight: %d\n", m.Height))

	if content := m.GetContent(); content != nil {
		b.WriteString(fmt.Sprintf("  Content: %s", content.String()))
	} else {
		b.WriteString("  Content: nil")
	}

	return b.String()
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (m PlannedProposal) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var content govTypes.Content
	return unpacker.UnpackAny(m.Content, &content)
}

// ShouldExecute checks if proposal should be executed now.
func (m PlannedProposal) ShouldExecute(ctx sdk.Context) bool {
	return ctx.BlockHeight() >= m.Height
}

// GetContent returns the proposal Content.
func (m PlannedProposal) GetContent() govTypes.Content {
	content, ok := m.Content.GetCachedValue().(govTypes.Content)
	if !ok {
		return nil
	}

	return content
}

// NewPlannedProposal creates a new Plan object.
func NewPlannedProposal(blockHeight int64, content govTypes.Content) (*PlannedProposal, error) {
	p := &PlannedProposal{
		Height: blockHeight,
	}

	protoMsg, ok := content.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("content: %T does not implement proto.Message", content)
	}

	protoAny, err := types.NewAnyWithValue(protoMsg)
	if err != nil {
		return nil, fmt.Errorf("content: converting to ProtoAny: %w", err)
	}
	p.Content = protoAny

	return p, nil
}
