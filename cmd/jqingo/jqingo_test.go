package jqingo

import (
	"reflect"
	"testing"
)

func TestProcessCommand(t *testing.T) {
	type args struct {
		compact bool
		args    []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Passing non compact JSON",
			args: args{
				args:    []string{`.`, `{"Test1":true,"Test2":"Yes"}`},
				compact: false,
			},
			wantErr: false,
		},
		{
			name: "Passing print compact JSON",
			args: args{
				args:    []string{`.`, `{"Test1": true,"Test2": "Yes"}`},
				compact: true,
			},
			wantErr: false,
		},
		{
			name: "Failing JSON",
			args: args{
				args:    []string{`.`, `{"Test1": true,"Test2": "Yes",}`},
				compact: false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ProcessCommand(tt.args.compact, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("ProcessCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_formatOutput(t *testing.T) {
	type args struct {
		compact bool
		input   string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Print non compact JSON",
			args: args{
				input:   `{"Test1":true,"Test2":"Yes"}`,
				compact: false,
			},
			want: []byte(`{
  "Test1": true,
  "Test2": "Yes"
}`),
			wantErr: false,
		},
		{
			name: "Print compact JSON",
			args: args{
				input: `{
  "Test1": true,
  "Test2": "Yes"
				  }`,
				compact: true,
			},
			want:    []byte(`{"Test1":true,"Test2":"Yes"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := formatOutput(tt.args.compact, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printOutput(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Fake test",
			args: args{
				input: []byte(`{"Test1":true,"Test2":"Yes"}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printOutput(tt.args.input)
		})
	}
}

func Test_compactOutput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []byte
		wantErr    bool
	}{
		{
			name: "Print compact JSON",
			args: args{
				input: `{
  "Test1": true,
  "Test2": "Yes"
				  }`,
			},
			wantOutput: []byte(`{"Test1":true,"Test2":"Yes"}`),
			wantErr:    false,
		},
		{
			name: "Fail with invalid JSON",
			args: args{
				input: `{
  "Test1": true,
  "Test2": "Yes
				  }`,
			},
			wantOutput: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := compactOutput(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("compactOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("compactOutput() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
