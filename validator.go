package validator

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("validation error")
)

// Rule functions return an error if the validation fails.
type Rule func() error

func Validate(rule Rule) error {
	if err := rule(); err != nil {
		return fmt.Errorf("%w: %w", ErrValidation, err)
	}
	return nil
}

// Any returns the first error encountered, or nil if all rules pass.
func Any(rules ...Rule) Rule {
	return func() error {
		for _, rule := range rules {
			if err := rule(); err != nil {
				return err
			}
		}
		return nil
	}
}

// All returns an error if all rules fail.
func All(rules ...Rule) Rule {
	return func() error {
		errs := make([]error, 0)
		for _, rule := range rules {
			if err := rule(); err != nil {
				errs = append(errs, err)
			}
		}
		if len(errs) > 0 {
			errs = append([]error{errors.New("multiple errors:\n%s")}, errs...)
			return errors.Join(errs...)
		}
		return nil
	}
}

// That returns a rule that checks the given condition evaluates to true.
func That(test bool, format string, args ...any) Rule {
	return func() error {
		if test {
			return nil
		}
		return fmt.Errorf(format, args...)
	}
}

// If returns a rule that is only executed if the given condition is true.
func If(test bool, rule Rule) Rule {
	return func() error {
		if test {
			return rule()
		}
		return nil
	}
}
