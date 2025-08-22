package common

import (
	"strings"
	"testing"
)

func Test_UnpackStr(t *testing.T) {
	tests := []struct {
		name    string
		argStr  string
		want    string
		wantErr bool
	}{
		{
			name:    "ok",
			argStr:  "a4bc2d5e",
			want:    "aaaabccddddde",
			wantErr: false,
		},
		{
			name:    "no digits",
			argStr:  "abcd",
			want:    "abcd",
			wantErr: false,
		},
		{
			name:    "empty string",
			argStr:  "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "only spaces",
			argStr:  "    ",
			want:    "",
			wantErr: true,
		},
		{
			name:    "only digits",
			argStr:  "45",
			want:    "",
			wantErr: true,
		},
		{
			name:    "big numbers1",
			argStr:  "f100b200",
			want:    strings.Repeat("f", 100) + strings.Repeat("b", 200),
			wantErr: false,
		},
		{
			name:    "big numbers2",
			argStr:  "qwe45",
			want:    "qw" + strings.Repeat("e", 45),
			wantErr: false,
		},
		{
			name:    "zeros before num",
			argStr:  "q0010f",
			want:    strings.Repeat("q", 10) + "f",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackStr(tt.argStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UnpackStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
