package logger

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	sentryApi "github.com/atlassian/go-sentry-api"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

// Test ENVs.
const (
	sentryTestEnvToken        = "DN_SENTRY_TEST_TOKEN"
	sentryTestEnvUrl          = "DN_SENTRY_TEST_URL"
	sentryTestEnvOrganisation = "DN_SENTRY_TEST_ORG"
	sentryTestEnvProject      = "DN_SENTRY_TEST_PRJ"
)

// Test defaults.
const (
	sentryTestDefaultUrl          = "https://sentry.dfinance.co/api/0/"
	sentryTestDefaultOrganisation = "sentry"
	sentryTestDefaultProject      = "integ-testing"
	sentryTestDefaultEnvironment  = "testing"
	sentryTestConnectionTimeoutS  = 5
	sentryTestPollTimeoutS        = 10
)

// TestWriter is used to route console output within a test to test's context.
type TestWriter struct {
	t *testing.T
}

// Write implements io.Writer.
func (w TestWriter) Write(p []byte) (n int, err error) {
	w.t.Log(string(p))

	return len(p), nil
}

func Test_SentryIntegration(t *testing.T) {
	const errMsgFmt = "dstation Sentry integ test msg #%d"

	// check inputs
	require.NoError(t, os.Setenv(DnSentryEnvEnvironment, sentryTestDefaultEnvironment), "set env")
	if os.Getenv(DnSentryEnvDsn) == "" {
		t.Skipf("%s env: empty", DnSentryEnvDsn)
	}
	inputSentryToken := os.Getenv(sentryTestEnvToken)
	if inputSentryToken == "" {
		t.Skipf("%s env: empty", sentryTestEnvToken)
	}

	// default overwrites
	inputSentryUrl := sentryTestDefaultUrl
	if env := os.Getenv(sentryTestEnvUrl); env != "" {
		inputSentryUrl = env
	}
	inputSentryOrg := sentryTestDefaultOrganisation
	if env := os.Getenv(sentryTestEnvOrganisation); env != "" {
		inputSentryOrg = env
	}
	inputSentryPrj := sentryTestDefaultProject
	if env := os.Getenv(sentryTestEnvProject); env != "" {
		inputSentryPrj = env
	}

	// setup Sentry
	require.NoError(t, SetupSentry("dstation", "vx.x.x", "_"), "Sentry init")

	// setup Sentry client
	sentryConTimeout := sentryTestConnectionTimeoutS
	sentryClient, err := sentryApi.NewClient(inputSentryToken, &inputSentryUrl, &sentryConTimeout)
	require.NoError(t, err, "create Sentry client")

	sentryOrg, err := sentryClient.GetOrganization(inputSentryOrg)
	require.NoError(t, err, "get Sentry organization")

	sentryPrj, err := sentryClient.GetProject(sentryOrg, inputSentryPrj)
	require.NoError(t, err, "gGet Sentry project")

	// setup logger with Sentry hook
	logger := server.ZeroLogWrapper{
		Logger: zerolog.New(
			zerolog.ConsoleWriter{
				Out: TestWriter{t},
			},
		).Level(zerolog.ErrorLevel).With().Timestamp().Logger(),
	}
	logger.Logger = logger.Hook(NewZeroLogSentryHook())

	// prepare and send error message
	logger.Info("Should not be printed")
	//
	errMsg := fmt.Sprintf(errMsgFmt, rand.Int())
	logger.Error(errMsg)
	errMsgSendAt := time.Now().UTC()

	// poll project issues searching for errMsg
	timeoutDur := sentryTestPollTimeoutS * time.Second
	timeoutCh := time.After(timeoutDur)
	targetIssue := sentryApi.Issue{}
	for {
		time.Sleep(100 * time.Millisecond)

		issues, _, err := sentryClient.GetIssues(sentryOrg, sentryPrj, nil, nil, nil)
		require.NoError(t, err, "get issues")

		found := false
		for _, issue := range issues {
			require.NotNil(t, issue.Title, "issue.Title")
			if *issue.Title != errMsg {
				t.Logf("Unexpected issue.Title (%s) received (skip)", *issue.Title)
				continue
			}

			events, _, err := sentryClient.GetIssueEvents(issue)
			require.NoError(t, err, "get issue events")
			require.Len(t, events, 1, "issue events length")
			event := events[0]

			require.NotNil(t, event.DateCreated, "event.DateCreated")
			if (*event.DateCreated).After(errMsgSendAt) {
				t.Logf("Unexpected event.DateCreated (%s) received (expected %s) (skip)", event.DateCreated.String(), errMsgSendAt.String())
				continue
			}

			envTagChecked := false
			require.NotNil(t, event.Tags, "event.Tags")
			for i, tag := range *event.Tags {
				require.NotNil(t, tag.Key, "event.Tag.Key %d", i)
				require.NotNil(t, tag.Value, "event.Tag.Value %d", i)
				if *(tag.Key) == "environment" && *(tag.Value) == sentryTestDefaultEnvironment {
					envTagChecked = true
				}
			}
			require.True(t, envTagChecked, "tag %q not found in event.Tags", "environment")

			targetIssue, found = issue, true
			break
		}

		if found {
			break
		}

		select {
		case <-timeoutCh:
			t.Fatalf("issue with msg %q not found after %v", errMsg, timeoutDur)
		default:
		}
	}

	// remove issue
	require.NoError(t, sentryClient.DeleteIssue(targetIssue), "remove issue")
}
