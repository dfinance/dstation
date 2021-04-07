package tests

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptoTypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	FmtInfColorPrefix = "\033[1;34m"
	FmtWrnColorPrefix = "\033[1;33m"
	FmtColorEndLine   = "\033[0m\n"
)

// GenAccAddress generates an sdk.AccAddress with its public / private keys.
func GenAccAddress() (sdk.AccAddress, cryptoTypes.PubKey, cryptoTypes.PrivKey) {
	pk := secp256k1.GenPrivKey()
	//pk := ed25519.GenPrivKey()

	return sdk.AccAddress(pk.PubKey().Address()), pk.PubKey(), pk
}

// WaitForFileExists waits for a file to be created.
func WaitForFileExists(filePath string, timeoutDur time.Duration) error {
	timeoutCh := time.After(timeoutDur)

	for {
		select {
		case <-timeoutCh:
			return fmt.Errorf("file %q did not appear after %v", filePath, timeoutDur)
		default:
			if _, err := os.Stat(filePath); err == nil {
				return nil
			}
		}
	}
}

// PingTcpAddress check the TCP connection (Dial with retry).
func PingTcpAddress(address string, timeout time.Duration) error {
	const dialTimeout = 500 * time.Millisecond

	// remove scheme prefix
	if i := strings.Index(address, "://"); i != -1 {
		address = address[i+3:]
	}

	retryCount := int(timeout / dialTimeout)
	connected := false
	for i := 0; i < retryCount; i++ {
		conn, err := net.DialTimeout("tcp", address, dialTimeout)
		if err == nil {
			connected = true
		}
		if conn != nil {
			conn.Close()
		}

		if connected {
			break
		}
	}

	if !connected {
		return fmt.Errorf("TCP ping to %s failed after %d retry attempts with %v timeout", address, retryCount, dialTimeout)
	}

	return nil
}
