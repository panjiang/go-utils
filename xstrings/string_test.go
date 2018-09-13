package xstrings

import "testing"

func TestLowerFirst(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"#1",
			args{"InvalidData"},
			"invalidData",
		},
		{
			"#2",
			args{"invalidData"},
			"invalidData",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerFirst(tt.args.s); got != tt.want {
				t.Errorf("LowerFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
