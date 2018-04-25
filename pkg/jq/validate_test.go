package jq

import (
	"testing"
)

func TestIsJSONValid(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Fail",
			args: args{
				input: "",
			},
			wantErr: true,
		},
		{
			name: "Pass",
			args: args{
				input: `{"Test1":true,"Test2":"Yes"}`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsJSONValid(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("IsJSONValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
