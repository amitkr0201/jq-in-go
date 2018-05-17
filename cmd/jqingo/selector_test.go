package jqingo

import (
	"reflect"
	"testing"
)

func Test_getSelectedJSON(t *testing.T) {
	type args struct {
		selector string
		input    string
	}
	tests := []struct {
		name             string
		args             args
		wantSelectedJSON string
		wantErr          bool
	}{
		{
			name: "Test1",
			args: args{
				selector: ".amit",
				input:    `{"amit":"kumar"}`,
			},
			wantSelectedJSON: `"kumar"`,
			wantErr:          false,
		},
		{
			name: "Empty selector",
			args: args{
				selector: ".",
				input:    `{"amit":"kumar"}`,
			},
			wantSelectedJSON: `{"amit":"kumar"}`,
			wantErr:          false,
		},
		{
			name: "Incorrect JSON",
			args: args{
				selector: ".",
				input:    `{"amit":"kumar}`,
			},
			wantSelectedJSON: ``,
			wantErr:          true,
		},
		{
			name: "Bigger JSON",
			args: args{
				selector: ".",
				input:    `{"amit":"kumar","incorrect":"JSON","third":{"pink":"yellow","dirty":true}}`,
			},
			wantSelectedJSON: `{"amit":"kumar","incorrect":"JSON","third":{"dirty":true,"pink":"yellow"}}`,
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSelectedJSON, err := getSelectedJSON(tt.args.selector, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSelectedJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSelectedJSON != tt.wantSelectedJSON {
				t.Errorf("getSelectedJSON() = %v, want %v", gotSelectedJSON, tt.wantSelectedJSON)
			}
		})
	}
}

func Test_applySelector(t *testing.T) {
	type args struct {
		input interface{}
		s     string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := applySelector(tt.args.input, tt.args.s); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("applySelector() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_interface2String(t *testing.T) {
	type args struct {
		f interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := interface2String(tt.args.f); gotOutput != tt.wantOutput {
				t.Errorf("interface2String() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
