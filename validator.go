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

// AssertAny returns the first error encountered, or nil if all rules pass.
func AssertAny(rules ...Rule) Rule {
	return func() error {
		for _, rule := range rules {
			if err := rule(); err != nil {
				return err
			}
		}
		return nil
	}
}

// AssertAll returns an error if all rules fail.
func AssertAll(rules ...Rule) Rule {
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

// AssertThat returns a rule that checks the given condition evaluates to true.
func AssertThat(test bool, format string, args ...any) Rule {
	return func() error {
		if test {
			return nil
		}
		return fmt.Errorf(format, args...)
	}
}

// AssertIf returns a rule that is only executed if the given condition is true.
func AssertIf(test bool, rule Rule) Rule {
	return func() error {
		if test {
			return rule()
		}
		return nil
	}
}
