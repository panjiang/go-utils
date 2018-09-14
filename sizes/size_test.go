package sizes

import (
	"fmt"
	"testing"
)

func TestToByte(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			"B#1",
			args{"10"},
			10,
			false,
		},
		{
			"B#2",
			args{"1024B"},
			1024,
			false,
		},
		{
			"K",
			args{"1K"},
			1024,
			false,
		},
		{
			"M",
			args{"1M"},
			1024 * 1024,
			false,
		},
		{
			"G",
			args{"1G"},
			1024 * 1024 * 1024,
			false,
		},
		{
			"T",
			args{"1T"},
			1024 * 1024 * 1024 * 1024,
			false,
		},
		{
			"E#1",
			args{"1TT"},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToByte(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	t.Log(toFixed(1.001, 1))
	t.Log(fmt.Sprintf("%s", toFixed(1.001, 1)))
	t.Log(ToString(1023))
	t.Log(ToString(1024))
	t.Log(ToString(1024 * 1024))
	t.Log(ToString(1023 * 500))
	t.Log(ToString(1024 * 1024 * 1024))
	t.Log(ToString(1024 * 1024 * 1024 * 1024))
	t.Log(ToString(1024 * 1024 * 1024 * 1024 * 1024))
}
