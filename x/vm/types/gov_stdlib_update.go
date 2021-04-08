package types

import (
	"fmt"
	"net/url"
	"strings"

	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeStdlibUpdate = "StdlibUpdate"
)

var (
	_ govTypes.Content = (*StdLibUpdateProposal)(nil)
)

// GetTitle implements govTypes.Content (wrapped by PlannedProposal).
func (m StdLibUpdateProposal) GetTitle() string       { return "DVM stdlib update" }

// GetDescription implements govTypes.Content (wrapped by PlannedProposal).
func (m StdLibUpdateProposal) GetDescription() string { return "Updates DVM stdlib code" }

// ProposalRoute implements govTypes.Content (wrapped by PlannedProposal).
func (m StdLibUpdateProposal) ProposalRoute() string  { return RouterKey }

// ProposalType implements govTypes.Content (wrapped by PlannedProposal).
func (m StdLibUpdateProposal) ProposalType() string   { return ProposalTypeStdlibUpdate }

// ValidateBasic implements govTypes.Content (wrapped by PlannedProposal).
func (m StdLibUpdateProposal) ValidateBasic() error {
	if m.Url == "" {
		return sdkErrors.Wrapf(ErrGovInvalidProposal, "url: empty")
	}
	if _, err := url.Parse(m.Url); err != nil {
		return sdkErrors.Wrapf(ErrGovInvalidProposal, "url: %v", err)
	}
	if m.UpdateDescription == "" {
		return sdkErrors.Wrapf(ErrGovInvalidProposal, "update_description: empty")
	}

	if len(m.Code) == 0 {
		return sdkErrors.Wrapf(ErrGovInvalidProposal, "code: empty")
	}
	for i, code := range m.Code {
		if len(code) == 0 {
			return sdkErrors.Wrapf(ErrGovInvalidProposal, "code [%d]: empty", i)
		}
	}

	return nil
}

// String implements govTypes.Content (wrapped by PlannedProposal).
func (m StdLibUpdateProposal) String() string {
	b := strings.Builder{}
	b.WriteString("Proposal:\n")
	b.WriteString(fmt.Sprintf("  Title: %s\n", m.GetTitle()))
	b.WriteString(fmt.Sprintf("  Description: %s\n", m.GetDescription()))
	b.WriteString(fmt.Sprintf("  Source URL: %s\n", m.Url))
	b.WriteString(fmt.Sprintf("  Update description: %s", m.UpdateDescription))

	return b.String()
}

// NewStdLibUpdateProposal creates a StdLibUpdateProposal object.
func NewStdLibUpdateProposal(url, updateDescription string, byteCode ...[]byte) govTypes.Content {
	return &StdLibUpdateProposal{
		Url:               url,
		UpdateDescription: updateDescription,
		Code:              byteCode,
	}
}
