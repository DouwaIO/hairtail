package linter

import (
	"fmt"

	"github.com/DouwaIO/hairtail/src/yaml/pipeline"
)

// A Linter lints a pipeline configuration.
type Linter struct {
	trusted bool
}

// New creates a new Linter with options.
func New(opts ...Option) *Linter {
	linter := new(Linter)
	for _, opt := range opts {
		opt(linter)
	}
	return linter
}

// Lint lints the configuration.
func (l *Linter) Lint(c *Pipeline) error {
	if err := l.lint(c); err != nil {
		return err
	}
	return nil
}

func (l *Linter) lint(s *Pipeline) error {
	for _, column := range s.Columns {
		if err := l.lintColumnName(column); err != nil {
			return err
		}
	}
	return nil
}

func (l *Linter) lintColumnName(c *Column) error {
	if len(c.Name) == 0 {
		return fmt.Errorf("Invalid or missing name")
	}
	return nil
}
