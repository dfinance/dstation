package logger

import "github.com/rs/zerolog"

var _ zerolog.Hook = ZeroLogSentryHook{}

// ZeroLogSentryHook is a ZeroLog hook for Sentry integration (publish for certain log events).
type ZeroLogSentryHook struct {
	levels map[zerolog.Level]struct{}
}

// Run implements zerolog.Hook interface.
func (h ZeroLogSentryHook) Run(event *zerolog.Event, level zerolog.Level, message string) {
	if _, found := h.levels[level]; !found {
		return
	}

	sentryCaptureMessage(message)
}

// NewZeroLogSentryHook creates a new ZeroLogSentryHook object.
func NewZeroLogSentryHook() ZeroLogSentryHook {
	h := ZeroLogSentryHook{
		levels: make(map[zerolog.Level]struct{}),
	}
	h.levels[zerolog.ErrorLevel] = struct{}{}
	h.levels[zerolog.FatalLevel] = struct{}{}
	h.levels[zerolog.PanicLevel] = struct{}{}

	return h
}
