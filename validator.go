// Package validator is a simple library for validating conditions. Conditions are defined
// with rules; If the rule fails, an error is returned.
//
//	err := validator.Validate(
//		validator.Rule(false, "must be false"),
//	)
//
//	if errors.Is(err, validator.ErrInvalid) {
//		// handle validation error
//	}
//
// The validator package also provides a function Any() that returns the first error encountered,
// or nil if all rules pass.
//
//	err := validator.Validate(
//		validator.Any(
//			validator.Rule(false, "must be false"),
//			validator.Rule(true, "must be true"),
//		),
//	) // returns "validation error: must be false"
//
// The validator package also provides a function All() that evaluates all rules in the list:
//
//	err := validator.Validate(
//		validator.All(
//			validator.Rule(true, "must be true"),
//			validator.Rule(false, "must be false"),
//		),
//	) // returns "validation error: must be false"
//
//	if errors.Is(err, validator.ErrInvalid) {
//		// handle validation error
//	}
package validator

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalid is the sentinel error that is wrapped by any error returned
	// by Validate().
	//
	//	err := validator.Validate(
	//		validator.Rule(true, "must be true"),
	//		validator.Rule(false, "must be false"),
	//	)
	//
	//	if errors.Is(err, validator.ErrInvalid) {
	//		// handle validation error
	//	}
	ErrInvalid = errors.New("validation error")
)

// ValidationRule functions return an error if the validation fails.
type ValidationRule func() error

// Validate evaluates the given validation rule. If the rule fails
// an validation error is returned.
func Validate(rule ValidationRule) error {
	if err := rule(); err != nil {
		return fmt.Errorf("%w: %w", ErrInvalid, err)
	}
	return nil
}

// Any evaluates all rules in the list and returns the first error encountered, or nil if all rules pass.
func Any(rules ...ValidationRule) ValidationRule {
	return func() error {
		for _, rule := range rules {
			if err := rule(); err != nil {
				return err
			}
		}
		return nil
	}
}

// All evaluates all rules in the list, and returns the first error encountered, or nil if all rules pass.
func All(rules ...ValidationRule) ValidationRule {
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

// Rule returns a validation rule that checks if the given condition evaluates to true.
func Rule(test bool, format string, args ...any) ValidationRule {
	return func() error {
		if test {
			return nil
		}
		return fmt.Errorf(format, args...)
	}
}

// If returns a rule that is only executed if the given condition is true.
func If(test bool, rule ValidationRule) ValidationRule {
	return func() error {
		if test {
			return rule()
		}
		return nil
	}
}
