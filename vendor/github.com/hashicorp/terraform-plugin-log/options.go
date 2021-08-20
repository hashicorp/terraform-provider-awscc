package tflog

import (
	"io"
	"os"
	"strings"

	"github.com/hashicorp/go-hclog"
)

// Option defines a modification to the configuration for a logger.
type Option func(loggerOpts) loggerOpts

type loggerOpts struct {
	name            string
	level           hclog.Level
	includeLocation bool

	// some private options to be used only by tflog for testing purposes
	// we should never export an Option that sets these
	output      io.Writer
	includeTime bool
}

func applyLoggerOpts(opts ...Option) loggerOpts {
	// set some defaults
	l := loggerOpts{
		includeLocation: true,
		includeTime:     true,
		output:          os.Stderr,
	}
	for _, opt := range opts {
		l = opt(l)
	}
	return l
}

func withOutput(output io.Writer) Option {
	return func(l loggerOpts) loggerOpts {
		l.output = output
		return l
	}
}

func withoutTimestamp() Option {
	return func(l loggerOpts) loggerOpts {
		l.includeTime = false
		return l
	}
}

// WithLogName returns an option that will set the logger name explicitly to
// `name`. This has no effect when used with NewSubsystem.
func WithLogName(name string) Option {
	return func(l loggerOpts) loggerOpts {
		l.name = name
		return l
	}
}

// WithLevelFromEnv returns an option that will set the level of the logger
// based on the string in an environment variable. The environment variable
// checked will be `name` and `subsystems`, joined by _ and in all caps.
func WithLevelFromEnv(name string, subsystems ...string) Option {
	return func(l loggerOpts) loggerOpts {
		envVar := strings.Join(subsystems, "_")
		if envVar != "" {
			envVar = "_" + envVar
		}
		envVar = strings.ToUpper(name + envVar)
		l.level = hclog.LevelFromString(os.Getenv(envVar))
		return l
	}
}

// WithLevel returns an option that will set the level of the logger.
func WithLevel(level hclog.Level) Option {
	return func(l loggerOpts) loggerOpts {
		l.level = level
		return l
	}
}

// WithoutLocation returns an option that disables including the location of
// the log line in the log output, which is on by default. This has no effect
// when used with NewSubsystem.
func WithoutLocation() Option {
	return func(l loggerOpts) loggerOpts {
		l.includeLocation = false
		return l
	}
}

// WithStderrFromInit returns an option that tells the logger to write to the
// os.Stderr that was present when the program started, not the one that is
// available at runtime. Some versions of Terraform overwrite os.Stderr with an
// io.Writer that is never read, so any log lines written to it will be lost.
func WithStderrFromInit() Option {
	return withOutput(stderr)
}
