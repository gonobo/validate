package validator_test

import (
	"testing"

	"github.com/gonobo/validator"
)

type args struct {
	rule validator.Rule
}

var tests = []struct {
	name    string
	args    args
	wantErr bool
}{
	{
		name: "that rule is true",
		args: args{
			rule: validator.AssertThat(2+2 == 4, "two plus two should equal four"),
		},
		wantErr: false,
	},
	{
		name: "that rule is false",
		args: args{
			rule: validator.AssertThat(2+2 == 5, "two plus two should not equal five"),
		},
		wantErr: true,
	},
	{
		name: "that all rules are true",
		args: args{
			rule: validator.AssertAll(
				validator.AssertThat(2+2 == 4, "two plus two should equal four"),
				validator.AssertThat(3+3 == 6, "three plus three should equal six"),
			),
		},
		wantErr: false,
	},
	{
		name: "that at least one rule is true",
		args: args{
			rule: validator.AssertAny(
				validator.AssertThat(2+2 == 4, "two plus two should equal four"),
				validator.AssertThat(3+3 == 6, "three plus three should equal six"),
			),
		},
		wantErr: false,
	},
	{
		name: "that at least one rule is false",
		args: args{
			rule: validator.AssertAny(
				validator.AssertThat(2+2 == 5, "two plus two should not equal five"),
				validator.AssertThat(3+3 == 6, "three plus three should equal six"),
			),
		},
		wantErr: true,
	},
	{
		name: "that at least one rule is false",
		args: args{
			rule: validator.AssertAll(
				validator.AssertThat(2+2 == 5, "two plus two should not equal five"),
				validator.AssertThat(3+3 == 6, "three plus three should equal six"),
			),
		},
		wantErr: true,
	},
	{
		name: "that if a value is false, then the rule is skipped",
		args: args{
			rule: validator.AssertIf(false, validator.AssertThat(false, "this will never run")),
		},
		wantErr: false,
	},
	{
		name: "that if a value is true, then the rule is evaluated",
		args: args{
			rule: validator.AssertIf(true, validator.AssertThat(false, "this will run")),
		},
		wantErr: true,
	},
}

func TestValidate(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validator.Validate(tt.args.rule); (err != nil) != tt.wantErr {
				t.Errorf("Any() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
