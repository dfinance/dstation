package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

const (
	// SentryEnvDsn is a default Sentry client ENV.
	SentryEnvDsn = "SENTRY_DSN"
	// SentryEnvEnvironment is a default Sentry client ENV.
	SentryEnvEnvironment = "SENTRY_ENVIRONMENT"

	// DnSentryEnvDsn overrides the SentryEnvDsn env value.
	DnSentryEnvDsn = "DN_SENTRY_DSN"
	// DnSentryEnvEnvironment overrides the SentryEnvEnvironment env value.
	DnSentryEnvEnvironment = "DN_SENTRY_ENVIRONMENT"
)

// sentryCfg is a global Sentry sentryConfig.
var sentryCfg sentryConfig

// sentryConfig keeps Sentry client params.
type sentryConfig struct {
	enabled         bool
	logSends        bool
	dsnToken        string
	environmentCode string
	hostname        string
	appName         string
	appVersion      string
	appCommit       string
	sendTimeout     time.Duration
}

// getClientOptions builds sentry.ClientOptions.
func (c sentryConfig) getClientOptions() sentry.ClientOptions {
	sentryTransport := sentry.NewHTTPSyncTransport()
	sentryTransport.Timeout = c.sendTimeout

	return sentry.ClientOptions{
		AttachStacktrace: true,
		Transport:        sentryTransport,
		Dsn:              c.dsnToken,
		Environment:      c.environmentCode,
		ServerName:       c.hostname,
		Release:          fmt.Sprintf("%s@%s [%s]", c.appName, c.appVersion, c.appCommit),
	}
}

// sentryCaptureMessage publishes a message to Sentry.
func sentryCaptureMessage(format string, args ...interface{}) {
	if !sentryCfg.enabled {
		sentryLogSend(nil)
		return
	}

	msg := fmt.Sprintf(format, args...)
	sentryLogSend(sentry.CaptureMessage(msg))
}

// sentryCaptureObject publishes an exception to Sentry.
func sentryCaptureObject(obj interface{}) {
	if !sentryCfg.enabled {
		sentryLogSend(nil)
		return
	}

	err := fmt.Errorf("%T: %v", obj, obj)
	sentryLogSend(sentry.CaptureException(err))
}

// sentryLogSend checks if Sentry integration is enabled and prints publish results to the console.
func sentryLogSend(eventId *sentry.EventID) {
	if !sentryCfg.logSends {
		return
	}

	if !sentryCfg.enabled {
		fmt.Println("sentry send: skipped")
		return
	}

	if eventId == nil {
		fmt.Println("sentry send: failed")
		return
	}
	fmt.Println("sentry send: done")
}

// SetupSentry initializes global sentryCfg config from env variables and enabled Sentry integration.
func SetupSentry(appName, appVersion, appCommit string) error {
	// force overwrite standard Sentry envs
	if err := os.Setenv(SentryEnvDsn, ""); err != nil {
		return fmt.Errorf("can't overwrite %s env: %w", SentryEnvDsn, err)
	}
	if err := os.Setenv(SentryEnvEnvironment, ""); err != nil {
		return fmt.Errorf("can't overwrite %s env: %w", SentryEnvEnvironment, err)
	}

	sentryDsn := os.Getenv(DnSentryEnvDsn)
	sentryEnvironment := os.Getenv(DnSentryEnvEnvironment)
	hostname, _ := os.Hostname()

	if appName == "" {
		appName = "undefined"
	}
	if appVersion == "" {
		appVersion = "v0.0.0"
	}
	if sentryEnvironment == "" {
		sentryEnvironment = "undefined"
	}

	sentryCfg.enabled = false
	sentryCfg.logSends = true
	sentryCfg.dsnToken = sentryDsn
	sentryCfg.environmentCode = sentryEnvironment
	sentryCfg.appName = appName
	sentryCfg.appVersion = appVersion
	sentryCfg.appCommit = appCommit
	sentryCfg.hostname = hostname
	sentryCfg.sendTimeout = 2 * time.Second

	if sentryCfg.dsnToken != "" {
		if err := sentry.Init(sentryCfg.getClientOptions()); err != nil {
			return fmt.Errorf("sentry init: %w", err)
		}
		sentryCfg.enabled = true
	}

	return nil
}

// CrashDeferHandler is a panic interceptor which publish a panic event to Sentry and panics back.
func CrashDeferHandler() {
	r := recover()
	if r == nil {
		return
	}

	sentryCaptureObject(r)
	panic(r)
}
