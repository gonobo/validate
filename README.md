# validator

[![Test](https://github.com/gonobo/validator/actions/workflows/test.yml/badge.svg)](https://github.com/gonobo/validator/actions/workflows/test.yml)
[![GoDoc](https://godoc.org/github.com/gonobo/validator?status.svg)](http://godoc.org/github.com/gonobo/validator)
[![Release](https://img.shields.io/github/release/gonobo/validate.svg)](https://github.com/gonobo/validator/releases)

A simple library for validating conditions. Conditions are defined with rules; If the rule fails, an error is returned.

## Installation

```bash
go get github.com/gonobo/validator
```

## Usage

```go
err := validator.Validate(
	validator.Rule(false, "must be false"),
)

if errors.Is(err, validator.ErrInvalid) {
	// handle validation error
}
```

The validator package also provides a function Any() that returns the first error encountered,
or `nil` if all rules pass.

```go
err := validator.Validate(
	validator.Any(
		validator.Rule(false, "must be false"),
		validator.Rule(true, "must be true"),
	),
) // returns "validation error: must be false"
```

The validator package also provides a function All() that evaluates all rules in the list:

```go
err := validator.Validate(
	validator.All(
		validator.Rule(true, "must be true"),
		validator.Rule(false, "must be false"),
	),
) // returns "validation error: must be false"

if errors.Is(err, validator.ErrInvalid) {
	// handle validation error
}
```
