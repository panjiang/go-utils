package versions

import (
	"testing"
)

func TestCmp(t *testing.T) {
	type args struct {
		v1 string
		v2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1.0==1.0",
			args{"1.0", "1.0"},
			0,
		},
		{
			"1.1>1.0",
			args{"1.1", "1.0"},
			1,
		},
		{
			"1.0<2.0",
			args{"1.0", "2.0"},
			-1,
		},
		{
			"1.0<1.0.1",
			args{"1.0", "1.0.1"},
			-1,
		},
		{
			"1.0<1.0.0",
			args{"1.0", "1.0.0"},
			0,
		},
		{
			"1.0.0.1<1.0.1.1",
			args{"1.0.0.1", "1.0.1.1"},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cmp(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("Cmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValid(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"#1",
			args{"1.0"},
			true,
		},
		{
			"#1",
			args{"0"},
			true,
		},
		{
			"#1",
			args{"0.0"},
			true,
		},
		{
			"#1",
			args{"1.0.001"},
			true,
		},
		{
			"#1",
			args{"a"},
			false,
		},
		{
			"#1",
			args{".1"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Valid(tt.args.version); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
