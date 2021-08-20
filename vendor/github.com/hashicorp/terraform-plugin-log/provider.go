package tflog

import (
	"context"

	"github.com/hashicorp/go-hclog"
)

func getProviderSpaceRootLogger(ctx context.Context) hclog.Logger {
	logger := ctx.Value(providerSpaceRootLoggerKey)
	if logger == nil {
		return nil
	}
	return logger.(hclog.Logger)
}

func setProviderSpaceRootLogger(ctx context.Context, logger hclog.Logger) context.Context {
	return context.WithValue(ctx, providerSpaceRootLoggerKey, logger)
}

// New returns a new context.Context that contains a logger configured with the
// passed options.
//
// This function isn't usually needed by plugin developers. Typically, the SDK
// or framework the plugin is built on will call New and configure a logger
// before handing off to the plugin code. This function is here for SDK and
// framework developers to do that setup work. Plugin developers should be able
// to safely assume that the logger already exists and just start using it.
func New(ctx context.Context, options ...Option) context.Context {
	opts := applyLoggerOpts(options...)
	if opts.level == hclog.NoLevel {
		opts.level = hclog.Trace
	}
	loggerOptions := &hclog.LoggerOptions{
		Name:                     opts.name,
		Level:                    opts.level,
		JSONFormat:               true,
		IndependentLevels:        true,
		IncludeLocation:          opts.includeLocation,
		DisableTime:              !opts.includeTime,
		Output:                   opts.output,
		AdditionalLocationOffset: 1,
	}
	return setProviderSpaceRootLogger(ctx, hclog.New(loggerOptions))
}

// With returns a new context.Context that has a modified logger in it which
// will include key and value as arguments in all its log output.
func With(ctx context.Context, key string, value interface{}) context.Context {
	logger := getProviderSpaceRootLogger(ctx)
	if logger == nil {
		// this essentially should never happen in production
		// the root logger for provider code should be injected
		// by whatever SDK the provider developer is using, so
		// really this is only likely in unit tests, at most
		// so just making this a no-op is fine
		return ctx
	}
	return setProviderSpaceRootLogger(ctx, logger.With(key, value))
}

// Trace logs `msg` at the trace level to the logger in `ctx`, with `args` as
// structured arguments in the log output. `args` is expected to be pairs of
// key and value.
func Trace(ctx context.Context, msg string, args ...interface{}) {
	logger := getProviderSpaceRootLogger(ctx)
	if logger == nil {
		// this essentially should never happen in production
		// the root logger for provider code should be injected
		// by whatever SDK the provider developer is using, so
		// really this is only likely in unit tests, at most
		// so just making this a no-op is fine
		return
	}
	logger.Trace(msg, args...)
}

// Debug logs `msg` at the debug level to the logger in `ctx`, with `args` as
// structured arguments in the log output. `args` is expected to be pairs of
// key and value.
func Debug(ctx context.Context, msg string, args ...interface{}) {
	logger := getProviderSpaceRootLogger(ctx)
	if logger == nil {
		// this essentially should never happen in production
		// the root logger for provider code should be injected
		// by whatever SDK the provider developer is using, so
		// really this is only likely in unit tests, at most
		// so just making this a no-op is fine
		return
	}
	logger.Debug(msg, args...)
}

// Info logs `msg` at the info level to the logger in `ctx`, with `args` as
// structured arguments in the log output. `args` is expected to be pairs of
// key and value.
func Info(ctx context.Context, msg string, args ...interface{}) {
	logger := getProviderSpaceRootLogger(ctx)
	if logger == nil {
		// this essentially should never happen in production
		// the root logger for provider code should be injected
		// by whatever SDK the provider developer is using, so
		// really this is only likely in unit tests, at most
		// so just making this a no-op is fine
		return
	}
	logger.Info(msg, args...)
}

// Warn logs `msg` at the warn level to the logger in `ctx`, with `args` as
// structured arguments in the log output. `args` is expected to be pairs of
// key and value.
func Warn(ctx context.Context, msg string, args ...interface{}) {
	logger := getProviderSpaceRootLogger(ctx)
	if logger == nil {
		// this essentially should never happen in production
		// the root logger for provider code should be injected
		// by whatever SDK the provider developer is using, so
		// really this is only likely in unit tests, at most
		// so just making this a no-op is fine
		return
	}
	logger.Warn(msg, args...)
}

// Error logs `msg` at the error level to the logger in `ctx`, with `args` as
// structured arguments in the log output. `args` is expected to be pairs of
// key and value.
func Error(ctx context.Context, msg string, args ...interface{}) {
	logger := getProviderSpaceRootLogger(ctx)
	if logger == nil {
		// this essentially should never happen in production
		// the root logger for provider code should be injected
		// by whatever SDK the provider developer is using, so
		// really this is only likely in unit tests, at most
		// so just making this a no-op is fine
		return
	}
	logger.Error(msg, args...)
}
