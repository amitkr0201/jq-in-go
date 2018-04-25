package cmd

import (
	"testing"
)

func Test_printOutput(t *testing.T) {
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
				input: "eh",
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
			if err := printOutput(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("printOutput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProcessCommand(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ZeroArg",
			args: args{
				args: []string{},
			},
			wantErr: false,
		},
		{
			name: "MoreThanOneArg",
			args: args{
				args: []string{`{"yes":"no"}`, `{"yes":"no"}`},
			},
			wantErr: true,
		},
		{
			name: "Passing",
			args: args{
				args: []string{`{"yes":"no"}`},
			},
			wantErr: false,
		},
		{
			name: "Failing",
			args: args{
				args: []string{`{"yes":"no",}`},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ProcessCommand(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("ProcessCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
