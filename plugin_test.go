package extfieldgen_test

import (
	"testing"

	"github.com/Warashi/extfieldgen"
)

func TestIsBuiltin(t *testing.T) {
	type test struct {
		name  string
		input string
		want  bool
	}

	tests := []test{
		{name: "string", input: "string", want: true},
		{name: "[]string", input: "[]string", want: true},
		{name: "*string", input: "*string", want: true},
		{name: "time.Time", input: "time.Time", want: false},
		{name: "SomeType", input: "SomeType", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extfieldgen.IsBuiltin(tt.input); got != tt.want {
				t.Errorf("extfieldgen.IsBuiltin(%s) = %t, but want = %t", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsFullName(t *testing.T) {
	type test struct {
		name  string
		input string
		want  bool
	}

	tests := []test{
		{name: "string", input: "string", want: false},
		{name: "[]string", input: "[]string", want: false},
		{name: "*string", input: "*string", want: false},
		{name: "time.Time", input: "time.Time", want: true},
		{name: "SomeType", input: "SomeType", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extfieldgen.IsFullName(tt.input); got != tt.want {
				t.Errorf("extfieldgen.IsFullName(%s) = %t, but want = %t", tt.input, got, tt.want)
			}
		})
	}
}

func TestMakeType(t *testing.T) {
	const importPath = "github.com/Warashi/extfieldgen/test-model"
	
	type test struct {
		name string
		input string
		want string
	}

	tests := []test{
		{name: "string", input: "string", want: "string"},
		{name: "[]string", input: "[]string", want: "[]string"},
		{name: "*string", input: "*string", want: "*string"},
		{name: "time.Time", input: "time.Time", want: "time.Time"},
		{name: "SomeType", input: "SomeType", want: importPath + ".SomeType"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extfieldgen.MakeType(importPath, tt.input); got != tt.want {
				t.Errorf("extfieldgen.MakeType(%s) = %s, but want = %s", tt.input, got, tt.want)
			}
		})
	}

}

